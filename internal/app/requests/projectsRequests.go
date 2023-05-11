package requests

type CreateProjectRequest struct {
	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Tags        []uint `json:"tag_ids"`
}
