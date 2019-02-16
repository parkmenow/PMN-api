package models

import (
	"time"
)

// DBModel contains basic primary key attribute for most entities
type DBModel struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

// Address represents one address
type Address struct {
	//DBModel
	Line1   string
	Line2   string
	Pincode string
}

// User represents entity having all attributes required to represent an app user
type User struct {
	DBModel
	FName    string
	LName    string
	UName    string
	Password string
	Email    string
	// 	Vehicles string
	PhoneNo string
	//Owner    Owner
	Address
}

// Owner specifies if a User has parking space to sublet
type Owner struct {
	DBModel
	Property []Property
	UserID   uint
}

// GpsLocation Location co-Ordinates fetched by API in terms of latitude and longitude
type GpsLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Property represents a single property owned by a owner
type Property struct {
	DBModel
	Address
	GpsLocation
	Spots   []Spot
	OwnerID uint
}

// Spot represent individual parking pots that can be sublet
type Spot struct {
	DBModel
	Type        int
	ImageURL    string
	Description string
	Booking     []TimeSlot
	PropertyID  uint
	PricePHr    int
}

// TimeSlot holds the booking and avaialability details for each Space for one day 12hrs where T variables holds bookingID if booked
type TimeSlot struct {
	DBModel
	T1      uint
	T2      uint
	T3      uint
	T4      uint
	T5      uint
	T6      uint
	T7      uint
	T8      uint
	T9      uint
	T10     uint
	T11     uint
	T12     uint
	SpaceID uint
}
