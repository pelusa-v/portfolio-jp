package entities

type Project struct {
	ID          uint   `json:"Id"`
	Name        string `json:"Name"`
	Summary     string `json:"Summary"`
	Description string `json:"Description"`
	ImageURL    string `json:"Image"`
	Tags        []Tag  `json:"Tags"`
	// References  []Reference `json:"URL"`
}
