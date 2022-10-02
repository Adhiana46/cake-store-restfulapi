package requests

type GetCakeRequest struct {
	ID int `uri:"id" binding:"required"`
}

// /cakes?limit=10&page=1&s=bla bla&rating_min=2.5&rating_max=4.5&sort_by=rating.desc,title.asc
type GetCakeListRequest struct {
	Limit     *int     `query:"limit" validate:"omitempty,numeric,min=1,max=100"`
	Page      *int     `query:"page" validate:"omitempty,numeric,min=1"`
	S         *string  `query:"s" validate:"omitempty,min=3,max=100"`
	RatingMin *float32 `query:"rating_min" validate:"omitempty,numeric,gte=0,lte=10"`
	RatingMax *float32 `query:"rating_max" validate:"omitempty,numeric,gte=0,lte=10"`
	SortBy    *string  `query:"sort_by" validate:"omitempty"`
}

type CreateCakeRequest struct {
	Title       string  `json:"title" validate:"required,min=3,max=100"`
	Description string  `json:"description" validate:"max=255"`
	Rating      float32 `json:"rating" validate:"numeric,gte=0,lte=10"`
	Image       string  `json:"image" validate:"max=255"`
}

type UpdateCakeRequest struct {
	ID          int      `uri:"id" binding:"required"`
	Title       *string  `json:"title" validate:"omitempty,min=3,max=100"`
	Description *string  `json:"description" validate:"omitempty,max=255"`
	Rating      *float32 `json:"rating" validate:"omitempty,numeric,gte=0,lte=10"`
	Image       *string  `json:"image" validate:"omitempty,max=255"`
}

type DeleteCakeRequest struct {
	ID int `uri:"id" binding:"required"`
}
