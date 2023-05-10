package main

import "database/sql"

type Storage interface {
	CreateBooking(*Booking) error
	UpdateBooking(string, *UpdateBookingRequest) error
	DeleteBooking(string) error
	GetBookings() ([]*Booking, error)
	GetBookingByID(string) (*Booking, error)
	GetBookingsByUserID(string) ([]*Booking, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	psqlInfo := "host=localhost port=55000 user=postgres password=postgrespw dbname=postgres sslmode=disable"
    
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) Init() error {
	return s.createBookingTable()
}

func (s *PostgressStore) createBookingTable() error {
	query := `create table if not exists booking (
		id varchar(50) primary key,
		start_date varchar(50),
		length_in_days serial,
		user_id varchar(50), 
		product_id varchar(50),
		stripe_invoice_id varchar(50),
		dates list,
		fulfilled boolean,
		extended boolean,
		created_at timestamp
	 )`

	 _, err := s.db.Exec(query)

	 return err
}

func (s *PostgressStore) CreateBooking(book *Booking) error {
	query :=  `insert into booking
	(id, start_date, length_in_days, user_id, product_id, stripe_invoice_id, dates, fulfilled, extended, created_at)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := s.db.Query(
		query, 
		book.ID,
		book.StartDate,
		book.LengthInDays,
		book.UserID, 
		book.ProductID,
		book.StripeInvoiceID, 
		book.Dates,
		book.Fulfilled, 
		book.Extended,
		book.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgressStore) UpdateBooking(id string, fields *UpdateBookingRequest) error {
	return nil
}

func (s *PostgressStore) DeleteBooking(id string) error {
	return nil
}

func (s *PostgressStore) GetBookings() ([]*Booking, error)  {
	return nil, nil
}

func (s *PostgressStore) GetBookingByID(id string) (*Booking, error)  {
	return  &Booking{}, nil
}

func (s *PostgressStore) GetBookingsByUserID(id string) ([]*Booking, error)  {
	return nil, nil
}

