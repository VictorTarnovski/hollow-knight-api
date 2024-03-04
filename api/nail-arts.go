package api

import (
	"net/http"

	"github.com/VictorTarnovski/hollow-knight-api/storage"
	"github.com/VictorTarnovski/hollow-knight-api/types"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleNailArts(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGETNailArts(w, r)
	}
	err := MethodNotAllowedHandler(w, r.Method)
	if err != nil {
		return err
	}
	return nil
}

func (s *APIServer) handleGETNailArts(w http.ResponseWriter, r *http.Request) error {
	var nailArts []*types.Collectable
	filters := storage.GetCollectablesFilters{TypeName: "Nail Art"}

	templ, err := mux.CurrentRoute(r).GetPathTemplate()
	if err != nil {
		return err
	}

	usesId := templ == "/nail-arts/{id}"
	if usesId == true {
		itemId := mux.Vars(r)["id"]
		filters.ID = &itemId
	}

	nailArts, err = s.store.GetCollectables(filters)
	if err != nil {
		return err
	}

	if usesId {
		if len(nailArts) == 0 {
			err := NotFoundHandler(w, "Nail Art")
			if err != nil {
				return err
			}
			return nil
		}

		err := OKHandler(w, nailArts[0])
		if err != nil {
			return err
		}
		return nil
	} else {
		err := OKHandler(w, nailArts)
		if err != nil {
			return err
		}
		return nil
	}
}
