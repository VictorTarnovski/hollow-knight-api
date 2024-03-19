package api

import (
	"net/http"
	"regexp"

	"github.com/VictorTarnovski/hollow-knight-api/storage"
	"github.com/VictorTarnovski/hollow-knight-api/types"
	"github.com/gorilla/mux"
)

func (s *APIServer) withKingdom(f APIFunc) APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		kingdomId := mux.Vars(r)["id"]

		kingdoms, err := s.store.GetKingdoms(storage.GetKingdomsFilters{ID: &kingdomId})
		if err != nil {
			return err
		}

		if len(kingdoms) == 0 {
			err := NotFoundHandler(w, "Kingdom")
			if err != nil {
				return err
			}
			return nil
		}

		err = f(w, r)
		if err != nil {
			return err
		}
		return nil
	}
}

func (s *APIServer) handleKingdoms(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		getAreas, err := regexp.Match("^/kingdoms/"+"("+UUIDRegex+")"+"/map"+"$", []byte(r.URL.Path))
		if err != nil {
			return err
		}

		if getAreas {
			return s.handleGETKingdomAreas(w, r)
		} else {
			return s.handleGETKingdoms(w, r)
		}
	}

	err := MethodNotAllowedHandler(w, r.Method)
	if err != nil {
		return err
	}
	return nil
}

func (s *APIServer) handleGETKingdoms(w http.ResponseWriter, r *http.Request) error {
	var kingdoms []*types.Kingdom
	filters := storage.GetKingdomsFilters{}

	templ, err := mux.CurrentRoute(r).GetPathTemplate()
	if err != nil {
		return err
	}

	usesId := templ == "/kingdoms/{id}"
	if usesId {
		kingdomId := mux.Vars(r)["id"]
		filters.ID = &kingdomId
	}

	kingdoms, err = s.store.GetKingdoms(filters)
	if err != nil {
		return err
	}

	if len(kingdoms) == 1 && usesId {
		err := OKHandler(w, kingdoms[0])
		if err != nil {
			return err
		}
		return nil
	} else {
		err := OKHandler(w, kingdoms)
		if err != nil {
			return err
		}
		return nil
	}
}

func (s *APIServer) handleGETKingdomAreas(w http.ResponseWriter, r *http.Request) error {
	areas, err := s.store.GetKingdomAreas(mux.Vars(r)["id"])

	if err != nil {
		return err
	}

	err = OKHandler(w, areas)
	if err != nil {
		return err
	}

	return nil
}
