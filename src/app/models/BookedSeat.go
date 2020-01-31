package models

import (
	"time"
)

func (bs *BookedSeat) TableName() string {
	return "booked_seats"
}

// Model Struct
type BookedSeat struct {
    Id   int `orm:"size(11)"`
    Seat *Seat `orm:"rel(fk)"` 
    User *User `orm:"rel(fk)"` 
    Status int `orm:"size(11)"`
    BookedAt time.Time `orm:"type(datetime)"`
}