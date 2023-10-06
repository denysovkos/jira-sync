package jira

import (
	"encoding/json"
	"errors"
	"fmt"
	"jira-sync/pkg/db"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

type Jira struct {
	client *resty.Client
}

func NewJiraClient(user, token string) *Jira {
	client := resty.New()
	client.
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(20 * time.Second)

	// Set Basic Authentication credentials
	client.SetBasicAuth(user, token)

	return &Jira{
		client: client,
	}
}

func (j *Jira) GetPendingTasks(params db.Params) (PendingTaskResponse, error) {
	// Make a GET request
	queryTemplate := "jql=project=\"JOIN\" AND assignee=currentUser() AND (remainingEstimate > 0 OR timespent = 0) AND updated >= -1w  order by created DESC"
	tasksTemplate := "https://%s.atlassian.net/rest/api/2/search"

	// Use fmt.Sprintf to format the template string with values
	url := fmt.Sprintf(tasksTemplate, params.Domain)

	resp, _ := j.client.R().SetQueryString(queryTemplate).Get(url)

	if resp.StatusCode() != 200 {
		r := fmt.Sprintf("HTTP request failed with status code: %d", resp.StatusCode())
		return PendingTaskResponse{}, errors.New(r)
	}

	var pendingTaskResponse PendingTaskResponse
	if err := json.Unmarshal(resp.Body(), &pendingTaskResponse); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	return pendingTaskResponse, nil
}

func (j *Jira) UpdateTime(params db.Params, issueKey, timeSpent string) error {
	updateSpentTimeTemplate := "https://%s.atlassian.net/rest/api/2/issue/%s/worklog"

	url := fmt.Sprintf(updateSpentTimeTemplate, params.Domain, issueKey)

	currentTime := time.Now().UTC()
	formattedTime := currentTime.Format("2006-01-02T15:04:05.000-0700")

	payload := map[string]interface{}{
		"timeSpent": timeSpent,
		"started":   formattedTime,
	}

	_, err := j.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(nil).
		Post(url)

	if err != nil {
		return err
	}

	return nil
}
