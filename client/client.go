package client

import (
	"net/http"
	"time"
)

const (
	DefaultAPIBaseURL       = "https://api.track.toggl.com"
	DefaultReportBaseURL    = "https://api.track.toggl.com/reports/api/v3"
	DefaultTimeEntryBaseURL = "https://api.track.toggl.com/api/v9"
)

type Client struct {
	Email    string
	Password string

	APIBaseURL       string
	ReportBaseURL    string
	TimeEntryBaseURL string

	HTTPClient *http.Client
	UserAgent  string
}

type Option func(*Client)

func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = hc
	}
}

func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.UserAgent = ua
	}
}

func WithCustomURLs(apiBaseURL, reportBaseURL, timeEntryBaseURL string) Option {
	return func(c *Client) {
		if apiBaseURL != "" {
			c.APIBaseURL = apiBaseURL
		}
		if reportBaseURL != "" {
			c.ReportBaseURL = reportBaseURL
		}
		if timeEntryBaseURL != "" {
			c.TimeEntryBaseURL = timeEntryBaseURL
		}
	}
}

func NewClient(email, password string, options ...Option) *Client {
	client := &Client{
		Email:    email,
		Password: password,

		APIBaseURL:       DefaultAPIBaseURL,
		ReportBaseURL:    DefaultReportBaseURL,
		TimeEntryBaseURL: DefaultTimeEntryBaseURL,

		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},

		UserAgent: "TogglClient/" + "0.1.0",
	}

	for _, opt := range options {
		opt(client)
	}

	return client
}
