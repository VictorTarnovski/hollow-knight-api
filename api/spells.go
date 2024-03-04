package api

import (
	"net/http"

	"github.com/VictorTarnovski/hollow-knight-api/storage"
	"github.com/VictorTarnovski/hollow-knight-api/types"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleSpells(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGETSpells(w, r)
	}
	err := MethodNotAllowedHandler(w, r.Method)
	if err != nil {
		return err
	}
	return nil
}

func (s *APIServer) handleGETSpells(w http.ResponseWriter, r *http.Request) error {
	var spells []*types.Collectable
	filters := storage.GetCollectablesFilters{TypeName: "Spell"}

	templ, err := mux.CurrentRoute(r).GetPathTemplate()
	if err != nil {
		return err
	}

	usesId := templ == "/spells/{id}"
	if usesId == true {
		spellId := mux.Vars(r)["id"]
		filters.ID = &spellId
	}

	spells, err = s.store.GetCollectables(filters)
	if err != nil {
		return err
	}

	if usesId {
		if len(spells) == 0 {
			err := NotFoundHandler(w, "Spell")
			if err != nil {
				return err
			}
			return nil
		}

		err := OKHandler(w, spells[0])
		if err != nil {
			return err
		}
		return nil
	} else {
		err := OKHandler(w, spells)
		if err != nil {
			return err
		}
		return nil
	}
}
