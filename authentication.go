package plex

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Pin Endpoint
type PinData struct {
	Errors           []Error   `json:"errors"`
	ID               int       `json:"id"`
	Code             string    `json:"code"`
	Product          string    `json:"product"`
	Trusted          bool      `json:"trusted"`
	ClientIdentifier string    `json:"clientIdentifier"`
	Location         Location  `json:"location"`
	ExpiresIn        int       `json:"expiresIn"`
	CreatedAt        time.Time `json:"createdAt"`
	ExpiresAt        time.Time `json:"expiresAt"`
	AuthToken        string    `json:"authToken"`
	NewRegistration  bool      `json:"newRegistration"`
	a                *App
}

type Location struct {
	Code         string `json:"code"`
	Country      string `json:"country"`
	City         string `json:"city"`
	TimeZone     string `json:"time_zone"`
	PostalCode   string `json:"postal_code"`
	Subdivisions string `json:"subdivisions"`
	Coordinates  string `json:"coordinates"`
}

func (a *App) GeneratePin() (*PinData, error) {
	form := url.Values{
		"strong":                   {"true"},
		"X-Plex-product":           {a.o.product},
		"X-Plex-client-Identifier": {a.o.clientIdentifier},
	}
	d, err := a.post(APIBaseURLv2+"/pins", strings.NewReader(form.Encode()))

	var p PinData
	err = json.Unmarshal(d, &p)
	if err != nil {
		return nil, err
	}
	if len(p.Errors) > 0 {
		return nil, p.Errors[0]
	}
	p.a = a
	return &p, nil
}

func (p *PinData) AuthUrl() (authAppURL string) {
	v := url.Values{
		"clientID":                 {p.ClientIdentifier},
		"code":                     {p.Code},
		"context[device][product]": {p.Product},
	}
	return fmt.Sprintf("https://app.plex.tv/auth#?%s", v.Encode())
}

func (p *PinData) Poll() (bool, error) {
	v := url.Values{
		"code":                     {p.Code},
		"X-Plex-client-Identifier": {p.ClientIdentifier},
	}
	d, err := p.a.c.getAnon(APIBaseURLv2+"/pins/"+strconv.Itoa(p.ID), v.Encode())
	if err != nil {
		return false, err
	}

	var pin PinData
	err = json.Unmarshal(d, &pin)
	if err != nil {
		return false, err
	}
	if pin.AuthToken != "" {
		*p = pin
		return true, nil
	}
	return false, nil
}
