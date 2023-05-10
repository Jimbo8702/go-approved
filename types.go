package main

import (
	"time"

	"github.com/google/uuid"
)

//create booking request

type CreateBookingRequest struct {
	StartDate string `json:"startDate"`
	LengthInDays int `json:"lengthInDays"`
	UserID string `json:"userID"`
	ProductID string `json:"productID"`
	StripeInvoiceID string `json:"stripeInvoiceID"`
	Dates []string `json:"dates"`
}

type UpdateBookingRequest struct {
	StartDate string `json:"startDate"`
	LengthInDays int `json:"lengthInDays"`
	Dates []string `json:"dates"`
	Fulfilled bool `json:"fulfilled"`
	Extended bool `json:"extended"`	
}

type Booking struct {
	ID string `json:"id"`
	StartDate string `json:"startDate"`
	LengthInDays int `json:"lengthInDays"`
	UserID string `json:"userID"`
	ProductID string `json:"productID"`
	StripeInvoiceID string `json:"stripeInvoiceID"`
	Fulfilled bool `json:"fulfilled"`
	Extended bool `json:"extended"`	
	Dates []string `json:"dates"`	
	CreatedAt time.Time `json:"createdAt"`
}

func NewBooking(startDate,userId, stripeInvoiceId, productId string,  length int, dates []string ) (*Booking, error) {
	uuid := uuid.NewString()

	return &Booking{
		ID: uuid,
		StartDate: startDate,
		LengthInDays: length,
		UserID: userId,
		ProductID: productId,
		StripeInvoiceID:stripeInvoiceId,
		Fulfilled: false,
		Extended: false,
		Dates: dates,
		CreatedAt: time.Now().UTC(),
	}, nil
}