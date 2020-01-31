package models

import (
	"time"
)

func (u *User) TableName() string {
    return "users"
}

// Model Struct
type User struct {
    Id   int
    Email string `orm:"size(50)"`
    Name string `orm:"size(20)"`
    CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}