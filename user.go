package plex

import (
	"encoding/xml"
	"time"
)

// User endpoint
type User struct {
	ID                      int          `json:"id" `
	UUID                    string       `json:"uuid" `
	Username                string       `json:"username" `
	Title                   string       `json:"title" `
	Email                   string       `json:"email" `
	Locale                  string       `json:"locale" `
	Confirmed               bool         `json:"confirmed" `
	EmailOnlyAuth           bool         `json:"emailOnlyAuth" `
	HasPassword             bool         `json:"hasPassword" `
	Protected               bool         `json:"protected" `
	Thumb                   string       `json:"thumb" `
	AuthToken               string       `json:"authToken"`
	MailingListStatus       string       `json:"mailingListStatus" `
	MailingListActive       bool         `json:"mailingListActive" `
	ScrobbleTypes           string       `json:"scrobbleTypes" `
	Country                 string       `json:"country" `
	Pin                     string       `json:"pin"`
	Subscription            Subscription `json:"subscription" `
	SubscriptionDescription string       `json:"subscriptionDescription" `
	Restricted              bool         `json:"restricted" `
	Anonymous               interface{}  `json:"anonymous" `
	Home                    bool         `json:"home" `
	Guest                   bool         `json:"guest" `
	HomeSize                int          `json:"homeSize" `
	HomeAdmin               bool         `json:"homeAdmin" `
	MaxHomeSize             int          `json:"maxHomeSize" `
	CertificateVersion      int          `json:"certificateVersion" `
	RememberExpiresAt       int          `json:"rememberExpiresAt" `
	Profile                 Profile      `json:"profile" `
	Entitlements            []string     `json:"entitlements" `
	Roles                   []string     `json:"roles" `
	Services                []Services   `json:"services" `
	AdsConsent              interface{}  `json:"adsConsent" `
	AdsConsentSetAt         interface{}  `json:"adsConsentSetAt" `
	AdsConsentReminderAt    interface{}  `json:"adsConsentReminderAt" `
	ExperimentalFeatures    bool         `json:"experimentalFeatures" `
	TwoFactorEnabled        bool         `json:"twoFactorEnabled" `
	BackupCodesCreated      bool         `json:"backupCodesCreated"`
	app                     *App
}

type Subscription struct {
	Active         bool      `json:"active"`
	SubscribedAt   time.Time `json:"subscribedAt"`
	Status         string    `json:"status"`
	PaymentService string    `json:"paymentService"`
	Plan           string    `json:"plan"`
	Features       []string  `json:"features"`
}

type Profile struct {
	AutoSelectAudio              bool   `json:"autoSelectAudio"`
	DefaultAudioLanguage         string `json:"defaultAudioLanguage"`
	DefaultSubtitleLanguage      string `json:"defaultSubtitleLanguage"`
	AutoSelectSubtitle           int    `json:"autoSelectSubtitle"`
	DefaultSubtitleAccessibility int    `json:"defaultSubtitleAccessibility"`
	DefaultSubtitleForced        int    `json:"defaultSubtitleForced"`
}

type Services struct {
	Identifier string `json:"identifier"`
	Endpoint   string `json:"endpoint"`
	Token      string `json:"token,omitempty"`
	Status     string `json:"status"`
	Secret     string `json:"secret,omitempty"`
}

func (u *User) Servers() ([]Server, error) {
	d, err := u.app.c.get(APIBaseURL+"/servers", u.AuthToken)
	if err != nil {
		return nil, err
	}

	var mcXML mediaContainerXML
	err = xml.Unmarshal(d, &mcXML)
	if err != nil {
		return nil, err
	}
	s := make([]Server, 0)
	for _, i := range mcXML.Servers {
		i.app = u.app
		s = append(s, i)
	}
	return s, nil
}

func (a *App) AttachUser(user *User) {
	user.app = a
}
