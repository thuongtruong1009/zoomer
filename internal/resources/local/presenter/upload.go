package presenter

type SingleUploadResponse struct {
	Image string `json:"image"`
	Size  int64  `json:"size"`
}

type MultipleUploadResponse struct {
	Images []SingleUploadResponse `json:"images"`
}
