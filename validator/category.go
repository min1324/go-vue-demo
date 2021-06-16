package validator

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
