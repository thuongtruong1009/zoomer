package presenter

// import "mime/multipart"

// type SingleUploadRequest struct {
// 	Image *multipart.FileHeader `json:"image" form:"image" validate:"required"`
// }

type SingleUploadResponse struct {
	Image string `json:"image"`
	Size  int64  `json:"size"`
}

// type MultipleUploadRequest struct {
// 	Images []*multipart.FileHeader `json:"images" form:"images" validate:"required"`
// }

type MultipleUploadResponse struct {
	Images []SingleUploadResponse `json:"images"`
}
