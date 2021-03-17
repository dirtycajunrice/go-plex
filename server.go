package plex

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	. "github.com/DirtyCajunRice/go-utility/types"
)

type mediaContainerXML struct {
	XMLName xml.Name `xml:"MediaContainer"`
	Servers []Server `xml:"Server"`
	Size    string   `xml:"size,attr"`
}

type Server struct {
	XMLName           xml.Name      `xml:"Server" json:"-"`
	AccessToken       string        `xml:"accessToken,attr" `
	Address           string        `xml:"address,attr" `
	CreatedAt         UnixTimestamp `xml:"createdAt,attr" json:"-" `
	Host              string        `xml:"host,attr" `
	LocalAddresses    string        `xml:"localAddresses,attr" `
	MachineIdentifier string        `xml:"machineIdentifier,attr" `
	Name              string        `xml:"name,attr" `
	Owned             bool          `xml:"owned,attr" `
	Port              int           `xml:"port,attr" `
	Scheme            string        `xml:"scheme,attr" `
	Synced            bool          `xml:"synced,attr" `
	UpdatedAt         UnixTimestamp `xml:"updatedAt,attr" json:"-" `
	Version           string        `xml:"version,attr" `
	OwnerId           int           `xml:"ownerId,attr" `
	app               App
}

func (s *Server) URL() string {
	return fmt.Sprintf("%s://%s:%d", s.Scheme, s.Host, s.Port)
}
func (s *Server) get(endpoint string) ([]byte, error) {
	return s.app.get(s.URL()+endpoint, s.AccessToken)
}

func (s *Server) command(endpoint string) error {
	return s.app.command(s.URL()+endpoint, s.AccessToken)
}

func (s *Server) Libraries() ([]Library, error) {
	d, err := s.get("/library/sections")
	if err != nil {
		return nil, err
	}

	var lr libraryResponse
	err = json.Unmarshal(d, &lr)
	if err != nil {
		return nil, err
	}

	return lr.Data.Sections, nil
}

func (s *Server) Scrobble(ratingKey int) error {
	return s.command(fmt.Sprintf("/:/Scrobble?identifier=com.plexapp.plugins.library&key=%d", ratingKey))
}
