package validator

type CreatePostRequest struct {
	CategoryId uint   `json:"category_id" binding:"required"`
	Tittle     string `json:"tittle" binding:"required,max=10"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" required:"required"`
}
