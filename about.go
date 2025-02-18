package dtrack

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type About struct {
	UUID        uuid.UUID      `json:"uuid"`
	SystemUUID  uuid.UUID      `json:"systemUuid"`
	Application string         `json:"application"`
	Version     string         `json:"version"`
	Timestamp   string         `json:"timestamp"`
	Framework   AboutFramework `json:"framework"`
}

type AboutFramework struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	Timestamp string    `json:"timestamp"`
}

type AboutService struct {
	client *Client
}

func (a AboutService) Get(ctx context.Context) (*About, error) {
	req, err := a.client.newRequest(ctx, http.MethodGet, "/api/version", withoutAuth())
	if err != nil {
		return nil, err
	}

	var about About
	if _, err = a.client.doRequest(req, &about); err != nil {
		return nil, err
	}

	return &about, nil
}
