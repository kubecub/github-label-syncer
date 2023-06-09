// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package exporter

import (
	"context"

	"github.com/google/go-github/v28/github"
)

type Label struct {
	// The name of the label.
	Name string `json:"name"`
	// An optional description of the label.
	Description string `json:"description"`
	// The hexadecimal color code for the label.
	Color string `json:"color"`

	// True if the label is the default label for the repository, false otherwise.
	Default bool `json:"default"`
	// The unique ID of the label assigned by GitHub.
	ID int64 `json:"id"`
	// The unique node ID of the label assigned by GitHub.
	NodeID string `json:"node_id"`
	// The URL to query information about the label.
	URL string `json:"url"`
}

/*
	Kubernetes API
URL： https://api.github.com/repos/Kubernetes/Kubernetes/labels?page=1&sort=name-asc
	{
	  "id": 1242861616,
	  "node_id": "MDU6TGFiZWwxMjQyODYxNjE2",
	  "url": "https://api.github.com/repos/kubernetes/kubernetes/labels/api-review",
	  "name": "api-review",
	  "color": "e11d21",
	  "default": false,
	  "description": "Categorizes an issue or PR as actively needing an API review."
	},
*/

func (c *githubClient) ListLabels(ctx context.Context, owner, repo string) ([]*Label, error) {
	opt := &github.ListOptions{PerPage: 10}
	var labels []*Label
	for {
		ghLabels, resp, err := c.client.Issues.ListLabels(ctx, owner, repo, opt)
		if err != nil {
			return nil, err
		}
		for _, l := range ghLabels {
			labels = append(labels, &Label{
				ID:          l.GetID(),
				NodeID:      l.GetNodeID(),
				URL:         l.GetURL(),
				Name:        l.GetName(),
				Color:       l.GetColor(),
				Default:     l.GetDefault(),
				Description: l.GetDescription(),
			})
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return labels, nil
}
