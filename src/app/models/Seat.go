package models

func (s *Seat) TableName() string {
    return "seats"
}

// Model Struct
type Seat struct {
    Id   int `orm:"size(11)"`
    SeatNumber string `orm:"size(11)"`
    Booked bool `orm:"-"`
}