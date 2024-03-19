package api

import (
	"net/http"

	"github.com/VictorTarnovski/hollow-knight-api/storage"
	"github.com/VictorTarnovski/hollow-knight-api/types"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleCollectables(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGETCollectables(w, r)
	}

	err := MethodNotAllowedHandler(w, r.Method)
	if err != nil {
		return err
	}
	return nil
}

func (s *APIServer) handleGETCollectables(w http.ResponseWriter, r *http.Request) error {
	var collectables []*types.Collectable
	var usesId bool
	filters := storage.GetCollectablesFilters{}

	templ, err := mux.CurrentRoute(r).GetPathTemplate()
	if err != nil {
		return err
	}

	if templ == "/items" || templ == "/items/{id}" {
		filters.TypeName = "Item"
		usesId = (templ == "/items/{id}")
	} else if templ == "/spells" || templ == "/spells/{id}" {
		filters.TypeName = "Spell"
		usesId = (templ == "/spells/{id}")
	} else if templ == "/nail-arts" || templ == "/nail-arts/{id}" {
		filters.TypeName = "Nail Art"
		usesId = (templ == "/nail-arts/{id}")
	}

	if usesId {
		collectableId := mux.Vars(r)["id"]
		filters.ID = &collectableId
	}

	collectables, err = s.store.GetCollectables(filters)
	if err != nil {
		return err
	}

	if usesId {
		if len(collectables) == 0 {
			err := NotFoundHandler(w, filters.TypeName)
			if err != nil {
				return err
			}
			return nil
		}

		err := OKHandler(w, collectables[0])
		if err != nil {
			return err
		}
		return nil
	} else {
		err := OKHandler(w, collectables)
		if err != nil {
			return err
		}
		return nil
	}
}
