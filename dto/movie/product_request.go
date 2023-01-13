package moviesdto

type MovieRequest struct {
	Title  string  `json:"title" form:"title" gorm:"type:varchar(255)" validate:"required"`
	Desc   string  `json:"desc" form:"desc" gorm:"text"`
	Rating float32 `json:"rating" gorm:"type: float32" form:"rating" validate:"required"`
	Image  string  `json:"image" form:"image" gorm:"type:varchar(255)"`
}

type UpdateMovie struct {
	Title  string  `json:"title" form:"title"`
	Desc   string  `json:"desc" form:"desc" gorm:"text"`
	Rating float32 `json:"rating" gorm:"type: float32" form:"rating"`
	Image  string  `json:"image" form:"image"`
}
