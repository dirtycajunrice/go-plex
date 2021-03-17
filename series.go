package plex

type seriesResponse struct {
	Data seriesData `json:"MediaContainer"`
}

type seriesData struct {
	data
	Art                 string   `json:"art"`
	LibrarySectionID    int      `json:"librarySectionID"`
	LibrarySectionTitle string   `json:"librarySectionTitle"`
	LibrarySectionUUID  string   `json:"librarySectionUUID"`
	Nocache             bool     `json:"nocache"`
	Thumb               string   `json:"thumb"`
	Title2              string   `json:"title2"`
	ViewGroup           string   `json:"viewGroup"`
	ViewMode            int      `json:"viewMode"`
	Series              []Series `json:"Metadata"`
}

type Series struct {
	RatingKey             int        `json:"ratingKey,string" `
	Key                   string     `json:"key" `
	SkipChildren          bool       `json:"skipChildren,omitempty" `
	GUID                  string     `json:"guid" `
	Studio                string     `json:"studio" `
	Type                  string     `json:"type" `
	Title                 string     `json:"title" `
	ContentRating         string     `json:"contentRating,omitempty" `
	Summary               string     `json:"summary" `
	Index                 int        `json:"index" `
	Rating                float64    `json:"rating,omitempty" `
	ViewCount             int        `json:"viewCount,omitempty" `
	LastViewedAt          int        `json:"lastViewedAt,omitempty" `
	Year                  int        `json:"year" `
	Thumb                 string     `json:"thumb" `
	Art                   string     `json:"art" `
	Banner                string     `json:"banner" `
	Theme                 string     `json:"theme,omitempty" `
	Duration              int        `json:"duration" `
	OriginallyAvailableAt string     `json:"originallyAvailableAt" `
	LeafCount             int        `json:"leafCount" `
	ViewedLeafCount       int        `json:"viewedLeafCount" `
	ChildCount            int        `json:"childCount" `
	AddedAt               int        `json:"addedAt" `
	UpdatedAt             int        `json:"updatedAt" `
	Genre                 []Metadata `json:"Genre" `
	Role                  []Metadata `json:"Role,omitempty" `
	TitleSort             string     `json:"titleSort,omitempty" `
	Collection            []Metadata `json:"Collection,omitempty" `
	l                     *Library
}

type Metadata struct {
	Tag string `json:"tag"`
}

func (s *Series) Scrobble() error {
	return s.l.s.scrobble(s.RatingKey)
}
