package models

import "time"

type Movie struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Title     string    `json:"title" form:"title" gorm:"type: varchar(255)"`
	Desc      string    `json:"desc" gorm:"type:text" form:"desc"`
	Rating    float32   `json:"rating" form:"rating"`
	Image     string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type MovieResponse struct {
	ID     int     `json:"id"`
	Title  string  `json:"name"`
	Desc   string  `json:"desc"`
	Rating float32 `json:"rating"`
	Image  string  `json:"image"`
}

func (MovieResponse) TableName() string {
	return "movies"
}
