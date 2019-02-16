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
	DBModel
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
	Email    string
// 	Vehicles string
	PhoneNo  string
	Owner    Owner
	Address  Address
	Password string
}

// Owner specifies if a User has parking space to sublet
type Owner struct {
	DBModel
	Property []Property
	UserID	  int
}

// GpsLocation Location co-Ordinates fetched by API in terms of latitude and longitude
type GpsLocation struct {
	lat  string
	long string
}

// Property represents a single property owned by a owner
type Property struct {
	DBModel
	Address      Address
	Location     GpsLocation
	OwnerID      uint
	ParkingSpots []Space
}

// Space represent individual parking pots that can be sublet
type Space struct {
	DBModel
	Type        int
	Photos      string
	Description string
	Booking     []TimeSlots
	PropertyID  uint
	PricePHr    uint
}

// TimeSlots holds the booking and avaialability details for each Space for one day 12hrs where T variables holds bookingID if booked
type TimeSlots struct {
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

type UserList struct {
	List []User `json:"users"`
}