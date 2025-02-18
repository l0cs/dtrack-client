package dtrack

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type Finding struct {
	Attribution   FindingAttribution `json:"attribution"`
	Analysis      Analysis           `json:"analysis"`
	Component     Component          `json:"component"`
	Matrix        string             `json:"matrix"`
	Vulnerability Vulnerability      `json:"vulnerability"`
}

type FindingAttribution struct {
	AlternateIdentifier string    `json:"alternateIdentifier"`
	AnalyzerIdentity    string    `json:"analyzerIdentity"`
	AttributedOn        int       `json:"attributedOn"`
	ReferenceURL        string    `json:"referenceUrl"`
	UUID                uuid.UUID `json:"uuid"`
}

type FindingsPage struct {
	Findings   []Finding
	TotalCount int
}

type FindingService struct {
	client *Client
}

func (f FindingService) GetAll(ctx context.Context, projectUUID uuid.UUID, suppressed bool, po PageOptions) (*FindingsPage, error) {
	params := map[string]string{
		"suppressed": strconv.FormatBool(suppressed),
	}

	req, err := f.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v1/finding/project/%s", projectUUID), withParams(params), withPageOptions(po))
	if err != nil {
		return nil, err
	}

	var findings []Finding
	res, err := f.client.doRequest(req, &findings)
	if err != nil {
		return nil, err
	}

	return &FindingsPage{
		Findings:   findings,
		TotalCount: res.TotalCount,
	}, nil
}
