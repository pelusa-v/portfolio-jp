package entities

type Project struct {
	Id          string      `json:"Id"`
	Name        string      `json:"Name"`
	References  []Reference `json:"URL"`
	Summary     string      `json:"Summary"`
	Description string      `json:"Description"`
	ImageURL    string      `json:"Image"`
	Tags        []Tag       `json:"Tags"`
}
