package plex

type data struct {
	Size            int    `json:"size"`
	AllowSync       bool   `json:"allowSync"`
	Identifier      string `json:"identifier"`
	MediaTagPrefix  string `json:"mediaTagPrefix"`
	MediaTagVersion int    `json:"mediaTagVersion"`
	Title1          string `json:"title1"`
}
