package dtrack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Component struct {
	UUID       uuid.UUID `json:"uuid"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	Group      string    `json:"group"`
	Name       string    `json:"name"`
	Version    string    `json:"version"`
	Classifier string    `json:"classifier"`
	FileName   string    `json:"filename"`
	Extension  string    `json:"extension"`

	MD5         string `json:"md5"`
	SHA1        string `json:"sha1"`
	SHA256      string `json:"sha256"`
	SHA384      string `json:"sha384"`
	SHA512      string `json:"sha512"`
	SHA3_256    string `json:"sha3_256"`
	SHA3_384    string `json:"sha3_384"`
	SHA3_512    string `json:"sha3_512"`
	BLAKE2b_256 string `json:"blake2b_256"`
	BLAKE2b_384 string `json:"blake2b_384"`
	BLAKE2b_512 string `json:"blake2b_512"`
	BLAKE3      string `json:"blake3"`

	CPE       string `json:"cpe"`
	PURL      string `json:"purl"`
	SWIDTagID string `json:"swidTagId"`

	Internal           bool     `json:"isInternal"`
	Description        string   `json:"description"`
	Copyright          string   `json:"copyright"`
	License            string   `json:"license"`
	ResolvedLicense    *License `json:"resolvedLicense"`
	DirectDependencies string   `json:"directDependencies"`
	Notes              string   `json:"notes"`
}

type ComponentsPage struct {
	Components []Component
	TotalCount int
}

type ComponentService struct {
	client *Client
}

func (c ComponentService) Get(ctx context.Context, componentUUID uuid.UUID) (*Component, error) {
	req, err := c.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v1/component/%s", componentUUID))
	if err != nil {
		return nil, err
	}

	var component Component
	_, err = c.client.doRequest(req, &component)
	if err != nil {
		return nil, err
	}

	return &component, nil
}

func (c ComponentService) GetAll(ctx context.Context, projectUUID uuid.UUID, po PageOptions) (*ComponentsPage, error) {
	req, err := c.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v1/component/project/%s", projectUUID), withPageOptions(po))
	if err != nil {
		return nil, err
	}

	var components []Component
	res, err := c.client.doRequest(req, &components)
	if err != nil {
		return nil, err
	}

	return &ComponentsPage{
		Components: components,
		TotalCount: res.TotalCount,
	}, nil
}
