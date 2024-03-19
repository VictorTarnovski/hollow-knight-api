package storage

import (
	"database/sql"
	"errors"

	"github.com/VictorTarnovski/hollow-knight-api/types"
	_ "github.com/lib/pq"
)

type GetKingdomsFilters struct {
	ID *string
}

type GetCollectablesFilters struct {
	TypeName string
	ID       *string
}

type Storage interface {
	GetKingdoms(filters GetKingdomsFilters) ([]*types.Kingdom, error)
	GetKingdomAreas(kingdomId string) ([]*types.Area, error)
	GetCollectables(filters GetCollectablesFilters) ([]*types.Collectable, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=higher_being password=pure_vessel dbname=hollow_knight_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) GetKingdoms(filters GetKingdomsFilters) ([]*types.Kingdom, error) {
	params := []any{}
	query := `
	SELECT 
		kingdoms.id,
		kingdoms.name
	FROM kingdoms
	WHERE 1 = 1
	`

	if filters.ID != nil {
		query = query + " AND kingdoms.id = $1"
		params = append(params, *filters.ID)
	}

	rows, err := s.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	kingdoms := []*types.Kingdom{}
	for rows.Next() {
		kingdom := &types.Kingdom{}

		err := rows.Scan(&kingdom.ID, &kingdom.Name)
		if err != nil {
			return nil, err
		}

		kingdoms = append(kingdoms, kingdom)
	}

	return kingdoms, nil
}

func (s *PostgresStore) GetKingdomAreas(kingdomId string) ([]*types.Area, error) {
	query := `
	SELECT 
		areas.id, 
		areas.name, 
		areas.quote
	FROM areas
	WHERE areas.kingdom_id = $1
	`

	rows, err := s.db.Query(query, kingdomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	areas := []*types.Area{}
	for rows.Next() {
		area := &types.Area{}

		err := rows.Scan(&area.ID, &area.Name, &area.Quote)
		if err != nil {
			return nil, err
		}

		areas = append(areas, area)
	}

	return areas, nil
}

var collectableListQuery = `
SELECT DISTINCT 
	collectables.id,
	collectables.name,
	collectables.description,
  collectable_upgrades.collectable_id IS NOT NULL AS has_upgrade
FROM collectables
INNER JOIN collectable_types ON collectable_types.id = collectables.collectable_type_id
LEFT JOIN collectable_upgrades ON collectable_upgrades.collectable_id = collectables.id
WHERE 1 = 1
  AND collectable_types.name = $1
  AND collectables.id not in
    (
     SELECT DISTINCT 
		 	collectable_upgrades.upgraded_collectable_id
     FROM collectable_upgrades
     INNER JOIN collectables ON collectables.id = collectable_upgrades.upgraded_collectable_id
     INNER JOIN collectable_types ON collectable_types.id = collectables.collectable_type_id
     WHERE 1 = 1
       AND collectable_types.name = $1
       AND collectable_upgrades.collectable_id != collectable_upgrades.upgraded_collectable_id 
    )
`
var collectableSubTreeQuery = `
SELECT 
	upgraded_collectables.id,
	upgraded_collectables.name,
	upgraded_collectables.description
FROM collectable_upgrades 
INNER JOIN collectables upgraded_collectables ON upgraded_collectables.id = collectable_upgrades.upgraded_collectable_id
WHERE 1 = 1
	and collectable_upgrades.collectable_id != collectable_upgrades.upgraded_collectable_id	
	and collectable_upgrades.collectable_id = $1
`
var collectableWithTreeQuery = `
SELECT 
    CASE
        WHEN upgraded_collectables.id IS NOT NULL THEN upgraded_collectables.id
        ELSE collectables.id
    END AS id,
    CASE
        WHEN upgraded_collectables.id IS NOT NULL THEN upgraded_collectables.name
        ELSE collectables.name
    END AS name,
    CASE
        WHEN upgraded_collectables.id IS NOT NULL THEN upgraded_collectables.description
        ELSE collectables.description
    END AS description,
   CASE
	WHEN upgraded_collectables.id IS NOT NULL THEN true
        ELSE false
    END AS has_upgrade
FROM collectables
INNER JOIN collectable_types ON collectable_types.id = collectables.collectable_type_id
LEFT JOIN collectable_upgrades ON collectable_upgrades.collectable_id = collectables.id
LEFT JOIN collectables AS upgraded_collectables ON upgraded_collectables.id = collectable_upgrades.upgraded_collectable_id
WHERE 1 = 1
  AND collectable_types.name = $1
  AND collectables.id = $2
ORDER BY collectable_upgrades.length
`

func (s *PostgresStore) GetCollectables(filters GetCollectablesFilters) ([]*types.Collectable, error) {
	if filters.TypeName == "" {
		return nil, errors.New("collectable typeName not informed")
	}

	if (filters.TypeName != "Item") && (filters.TypeName != "Spell") && (filters.TypeName != "Nail Art") {
		return nil, errors.New("invalid collectable typeName")
	}

	query := ""
	params := []any{filters.TypeName}

	if filters.ID == nil {
		query = collectableListQuery
	} else {
		query = collectableWithTreeQuery
		params = append(params, *filters.ID)
	}

	rows, err := s.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	collectables := []*types.Collectable{}
	for rows.Next() {
		collectable := &types.Collectable{}

		err := rows.Scan(&collectable.ID, &collectable.Name, &collectable.Description, &collectable.HasUpgrade)
		if err != nil {
			return nil, err
		}

		if collectable.HasUpgrade && filters.ID == nil {
			subRows, err := s.db.Query(collectableSubTreeQuery, collectable.ID)
			if err != nil {
				return nil, err
			}
			defer subRows.Close()
			upgrades := []*types.Collectable{}
			for subRows.Next() {
				upgrade := &types.Collectable{}

				err := subRows.Scan(&upgrade.ID, &upgrade.Name, &upgrade.Description)
				if err != nil {
					return nil, err
				}

				upgrades = append(upgrades, upgrade)
			}
			collectable.UpgradesTo = SliceToCollectableTree(upgrades)

		}
		collectables = append(collectables, collectable)
	}

	if len(collectables) != 0 && len(collectables) != 1 {
		if filters.ID != nil {
			collectables[0] = SliceToCollectableTree(collectables)
			return collectables[:1], nil
		}
	}

	return collectables, nil
}

func SliceToCollectableTree(slice []*types.Collectable) *types.Collectable {
	if len(slice) == 0 || len(slice) == 1 {
		return nil
	}

	for i := len(slice) - 1; i > 0; i-- {
		slice[i-1].UpgradesTo = slice[i]
	}
	return slice[0]
}
