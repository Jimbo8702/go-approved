package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type APIServer struct {
	listenAddr string
	store Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}
}


func (s *APIServer) Run() {
	router := mux.NewRouter()


	router.HandleFunc("/create", makeHTTPHandleFunc(s.handleCreateBooking))
	router.HandleFunc("/update", makeHTTPHandleFunc(s.handleUpdateBooking))
	router.HandleFunc("/delete", makeHTTPHandleFunc(s.handleDeleteBooking))
	router.HandleFunc("/bookings/all", makeHTTPHandleFunc(s.handleGetBookings))
	router.HandleFunc("/bookings/{id}", makeHTTPHandleFunc(s.handleGetBookingById))
	router.HandleFunc("/bookings/user/{id}", makeHTTPHandleFunc(s.handleGetBookingsByUserId))

	
	log.Println("JSON API server running on port:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}


func (s *APIServer) handleCreateBooking(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleUpdateBooking(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteBooking(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetBookings(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetBookingById(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetBookingsByUserId(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error  {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}


type apiFunc func(http.ResponseWriter, *http.Request) error 

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w,r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}