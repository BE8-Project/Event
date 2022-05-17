package request

type InsertCateg struct {
	Name string `json:"name" validate:"required"`
}
