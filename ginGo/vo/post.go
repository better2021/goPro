package vo

type CreatePostRequest struct {
	Category_id uint `json:"category_id" binding:"required"`
	Title string `json:"title" binding:"required,max=10"`
	HeadImg string `json:"head_img"`
	Content string `json:"content" binding:"required"`
}