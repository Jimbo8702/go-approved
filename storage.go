package main

import "database/sql"


type Storage interface {
	CreateBooking(*Booking) error
	UpdateBooking(*Booking) error
	DeleteBooking(int) error
	GetBookings() ([]*Booking, error)
	GetBookingsByID(int) (*Booking, error)
	GetBookingsByUserID() (*Booking, error)
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
		id serial primary key,
		start_date varchar(50),
		length_in_days serial,
		user_id serial, 
		product_id serial,
		stripe_invoice_id serial,
		fulfilled boolean,
		extended boolean,
		created_at timestamp
	 )`

	 _, err := s.db.Exec(query)

	 return err
}

func (s *PostgressStore) CreateBooking(book *Booking) error {
	query :=  `insert into booking
	(start_date, length_in_days, user_id, product_id, stripe_invoice_id, fulfilled, extended, created_at)
	values ($1, $2, $3, $4, $5, $6, $7)`


	_, err := s.db.Query(
		query, 
		book.StartDate,
		book.LengthInDays,
		book.UserID, 
		book.ProductID,
		book.StripeInvoiceID, 
		book.Fulfilled, 
		book.Extended,
		book.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgressStore) UpdateBooking(*Booking) error {
	return nil
}

func (s *PostgressStore) DeleteBooking(id int) error {
	return nil
}

func (s *PostgressStore) GetBookings() ([]*Booking, error)  {
	return nil, nil
}

func (s *PostgressStore) GetBookingById(id int) (*Booking, error)  {
	return nil, nil
}

func (s *PostgressStore) GetBookingsByUserId(id int) ([]*Booking, error)  {
	return nil, nil
}

