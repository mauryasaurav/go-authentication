package dto

/* Question Validation*/
type QuestionValidator struct {
	Question   string `json:"question" binding:"required"`
	Answer     string `json:"answer" binding:"required"`
	CategoryID string `json:"category_id"`
}
