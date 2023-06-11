package entity

type Photo struct {
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`
	Location    string `json:"location"`
	Date        string `json:"date"`
}

type ListPhotos []ListPhoto

type ListPhoto struct {
	PhotoID     string `json:"photo_id"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`
	Location    string `json:"location"`
	Date        string `json:"date"`
}

type PhotoID struct {
	PhotoID string `json:"photo_id"`
}
