package main

import "time"

type Booking struct {
	ID int `json:"id"`
	StartDate string `json:"startDate"`
	LengthInDays int `json:"lengthInDays"`
	UserID int `json:"userID"`
	ProductID int `json:"productID"`
	StripeInvoiceID int `json:"stripeInvoiceID"`
	Fulfilled bool `json:"fulfilled"`
	Extended bool `json:"extended"`	
	CreatedAt time.Time `json:"createdAt"`
}


func NewBooking(startDate string, stripeInvoiceId,productId, length, userId int ) (*Booking, error) {
	return &Booking{
		StartDate: startDate,
		LengthInDays: length,
		UserID: userId,
		ProductID: productId,
		StripeInvoiceID:stripeInvoiceId,
		Fulfilled: false,
		Extended: false,
		CreatedAt: time.Now().UTC(),
	}, nil
}