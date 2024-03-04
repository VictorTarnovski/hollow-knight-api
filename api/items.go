package api

import (
	"net/http"

	"github.com/VictorTarnovski/hollow-knight-api/storage"
	"github.com/VictorTarnovski/hollow-knight-api/types"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleItems(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGETItems(w, r)
	}
	err := MethodNotAllowedHandler(w, r.Method)
	if err != nil {
		return err
	}
	return nil
}

func (s *APIServer) handleGETItems(w http.ResponseWriter, r *http.Request) error {
	var items []*types.Collectable
	filters := storage.GetCollectablesFilters{TypeName: "Item"}

	templ, err := mux.CurrentRoute(r).GetPathTemplate()
	if err != nil {
		return err
	}

	usesId := templ == "/items/{id}"
	if usesId == true {
		itemId := mux.Vars(r)["id"]
		filters.ID = &itemId
	}

	items, err = s.store.GetCollectables(filters)
	if err != nil {
		return err
	}

	if usesId {
		if len(items) == 0 {
			err := NotFoundHandler(w, "Item")
			if err != nil {
				return err
			}
			return nil
		}

		err := OKHandler(w, items[0])
		if err != nil {
			return err
		}
		return nil
	} else {
		err := OKHandler(w, items)
		if err != nil {
			return err
		}
		return nil
	}
}
