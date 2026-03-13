package dto

type CreateNoteRequest struct {
	Title   string `json:"title" binding:"required,min=3,max=100"`
	Content string `json:"content" binding:"required,min=5,max=1000"`
	Pinned  bool   `json:"pinned"`
}