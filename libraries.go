package plex

import (
	"encoding/json"
	"fmt"
)

type libraryResponse struct {
	Data libraryData `json:"MediaContainer"`
}

type libraryData struct {
	data
	Sections []Library `json:"Directory"`
}

type Library struct {
	AllowSync        bool              `json:"-" `
	Art              string            `json:"-" `
	Composite        string            `json:"-" `
	Filters          bool              `json:"-" `
	Refreshing       bool              `json:"-" `
	Thumb            string            `json:"-" `
	Key              int               `json:"key,string" `
	Type             string            `json:"type" `
	Title            string            `json:"title" `
	Agent            string            `json:"agent" `
	Scanner          string            `json:"scanner" `
	Language         string            `json:"-" `
	UUID             string            `json:"uuid" `
	UpdatedAt        int               `json:"-" `
	CreatedAt        int               `json:"-" `
	ScannedAt        int               `json:"-" `
	Content          bool              `json:"-" `
	Directory        bool              `json:"-" `
	ContentChangedAt int               `json:"-" `
	Hidden           int               `json:"-" `
	Location         []LibraryLocation `json:"-" `
	s                *Server
}

type LibraryLocation struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
}

func (l *Library) Series(filter bool) ([]Series, error) {
	f := ""
	if filter {
		f = "?type=2unwatchedLeaves=1"
	}
	d, err := l.s.get(fmt.Sprintf("/library/sections/%d/all%s", l.Key, f))
	if err != nil {
		return nil, err
	}

	var sr seriesResponse
	err = json.Unmarshal(d, &sr)
	if err != nil {
		return nil, err
	}

	return sr.Data.Series, nil
}
