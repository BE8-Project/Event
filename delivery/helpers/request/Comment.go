package request

type CommentInsert struct {
	EventID uint   `json:"eventId" validate:"required"`
	Field   string `json:"field" validate:"required"`
}
type Commentget struct {
	EventID uint `json:"eventId" validate:"required"`
}
