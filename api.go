package main

import (
	"encoding/json"
	"fmt"
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
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}
	
	req := new(CreateBookingRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	booking, err := NewBooking(req.StartDate, req.UserID, req.StripeInvoiceID, req.ProductID, req.LengthInDays, req.Dates)
	if err != nil {
		return err
	}

	if err := s.store.CreateBooking(booking); err != nil {
		 return err 
	}

	return WriteJSON(w, http.StatusOK, booking)
}

func (s *APIServer) handleUpdateBooking(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "UPDATE" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	id := getID(r)
	

	req := new(UpdateBookingRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	if err := s.store.UpdateBooking(id, req); err != nil {
		return err 
   }

   return WriteJSON(w, http.StatusOK, nil)
}

func (s *APIServer) handleDeleteBooking(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "DELETE" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}
	
	id := getID(r)
	

	if err := s.store.DeleteBooking(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]string{"deleted": id})
}

func (s *APIServer) handleGetBookings(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	bookings, err := s.store.GetBookings()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, bookings)
}

func (s *APIServer) handleGetBookingById(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	id := getID(r)
	

	booking, err := s.store.GetBookingByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, booking)
}

func (s *APIServer) handleGetBookingsByUserId(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	id := getID(r)

	booking, err := s.store.GetBookingsByUserID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, booking)
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

func getID(r *http.Request) (string) {
	id := mux.Vars(r)["id"]

	return id
}