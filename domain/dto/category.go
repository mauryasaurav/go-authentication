package dto

/* Category Validation*/
type CategoryValidator struct {
	Name string `json:"name" binding:"required"`
}
