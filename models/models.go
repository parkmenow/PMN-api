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
	UName    string `gorm:"type:varchar(40); not null`
	Password string `gorm:"type:varchar(40); not null` // TODO: Change this to an encrypted version
	Email    string
	// 	Vehicles string
	PhoneNo string
	Wallet  int64
	Address // TODO: Once requirements are satisfied, then change this
}

// Owner specifies if a User has parking space to sublet
type Owner struct {
	DBModel
	Property []Property
	UserID   uint
}

// GpsLocation Location co-Ordinates fetched by API in terms of latitude and longitude
type GpsLocation struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// Property represents a single property owned by a owner
type Property struct {
	DBModel
	Address // TODO: Once requirements are satisfied, then change this
	GpsLocation // TODO: Once requirements are satisfied, then change this
	Spots   []Spot
	OwnerID uint
}

// Spot represent individual parking pots that can be sublet
type Spot struct { // TODO: Refactor this spot name to ParkingSpot
	DBModel
	Type        int // TODO: What is this actually, can be transformed to string type.
	ImageURL    string
	Description string
	Slots       []Slot // TODO: Refactor
	PropertyID  uint
}

// Slot holds the booking and avaialability details for each Space for one day 12hrs where T variables holds bookingID if booked
//TODO: change data type of start and end time
type Slot struct { // TODO: Refactor this spot name to TimeSlot
	DBModel
	StartTime time.Time `gorm:"type:timestamp with time zone"`
	EndTime   time.Time `gorm:"type:timestamp with time zone"`
	// StartTime time.Time `gorm:"type:datetime"`
	// EndTime   time.Time `gorm:"type:datetime"`
	Price     int
	SpotID    uint
	Available bool
}

//SearchInput is the input details from user to search parking spots
type SearchInput struct {
	Type      int
	StartTime time.Time `gorm:"type:timestamp with time zone"`
	EndTime   time.Time `gorm:"type:timestamp with time zone"`
	GpsLocation
}

// Booking is
type Booking struct {
	DBModel
	UserID  uint
	OwnerID uint
	SlotID  uint
	Price   int64
	Status string
}
