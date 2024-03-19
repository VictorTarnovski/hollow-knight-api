package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictorTarnovski/hollow-knight-api/storage"
	"github.com/gorilla/mux"
)

const UUIDRegex = "[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}"

type APIServer struct {
	listenAddress string
	store         storage.Storage
}

func NewAPIServer(listenAddress string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	/* /kingdoms */
	router.HandleFunc("/kingdoms", makeHTTPHandleFunc(s.handleKingdoms))
	router.HandleFunc("/kingdoms/{id}", makeHTTPHandleFunc(s.withKingdom(s.handleKingdoms)))
	router.HandleFunc("/kingdoms/{id}/map", makeHTTPHandleFunc(s.withKingdom(s.handleKingdoms)))

	/* /items */
	router.HandleFunc("/items", makeHTTPHandleFunc(s.handleCollectables))
	router.HandleFunc("/items/{id}", makeHTTPHandleFunc(s.handleCollectables))

	/* /spells */
	router.HandleFunc("/spells", makeHTTPHandleFunc(s.handleCollectables))
	router.HandleFunc("/spells/{id}", makeHTTPHandleFunc(s.handleCollectables))

	/* /nail-arts */
	router.HandleFunc("/nail-arts", makeHTTPHandleFunc(s.handleCollectables))
	router.HandleFunc("/nail-arts/{id}", makeHTTPHandleFunc(s.handleCollectables))

	err := http.ListenAndServe(s.listenAddress, router)

	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

func NewNotFoundError(resource string) APIError {
	return APIError{
		StatusCode: http.StatusNotFound,
		Error:      fmt.Sprint(resource, " was not found"),
	}
}

func NewMethodNotAllowedError(method string) APIError {
	return APIError{
		StatusCode: http.StatusMethodNotAllowed,
		Error:      fmt.Sprint("Method ", method, " is not allowed for this resource"),
	}
}

func NewInternalServerError() APIError {
	return APIError{
		StatusCode: http.StatusInternalServerError,
		Error:      "Internal Server Error",
	}
}

func makeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			InternalServerHandler(w)
			fmt.Println(err)
		}
	}
}

func OKHandler(w http.ResponseWriter, v any) error {
	err := WriteJSON(w, http.StatusOK, v)
	if err != nil {
		return err
	}
	return nil
}

func NotFoundHandler(w http.ResponseWriter, resource string) error {
	err := WriteJSON(w, http.StatusNotFound, NewNotFoundError(resource))
	if err != nil {
		return err
	}
	return nil
}

func MethodNotAllowedHandler(w http.ResponseWriter, method string) error {
	err := WriteJSON(w, http.StatusMethodNotAllowed, NewMethodNotAllowedError(method))
	if err != nil {
		return err
	}
	return nil
}

func InternalServerHandler(w http.ResponseWriter) error {
	err := WriteJSON(w, http.StatusInternalServerError, NewInternalServerError())
	if err != nil {
		return err
	}
	return nil
}
