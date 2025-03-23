package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/na0x2c6/toggl-client/models"
)

func (c *Client) GetTimeEntries(meta, includeSharing bool, since *int64, before, startDate, endDate *string) ([]models.TimeEntry, error) {
	url := fmt.Sprintf("%s/me/time_entries", c.TimeEntryBaseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if meta {
		q.Add("meta", "true")
	}
	if includeSharing {
		q.Add("incluede_sharing", "true")
	}
	if since != nil {
		q.Add("since", fmt.Sprintf("%d", since))
	}
	if before != nil {
		q.Add("before", *before)
	}
	if startDate != nil {
		q.Add("start_date", *startDate)
	}
	if endDate != nil {
		q.Add("end_date", *endDate)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %d %s", resp.StatusCode, string(body))
	}

	var timeEntries []models.TimeEntry
	if err := json.NewDecoder(resp.Body).Decode(&timeEntries); err != nil {
		return nil, err
	}

	return timeEntries, nil
}
