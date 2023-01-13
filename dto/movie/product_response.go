package moviesdto

type MovieResponse struct {
	ID     int     `json:"id"`
	Title  string  `json:"title" form:"title" validate:"required"`
	Desc   string  `json:"desc" form:"desc" gorm:"type:varchar(255)"`
	Rating float32 `json:"rating" form:"rating" gorm:"type: float32" validate:"required"`
	Image  string  `json:"image" form:"id" validate:"required"`
}
