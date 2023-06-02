// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kubecub/github-label-syncer/pkg/github"
	"go.uber.org/multierr"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	manifest := os.Getenv("INPUT_MANIFEST")
	labels, err := github.FromManifestToLabels(manifest)
	if err != nil {
		return fmt.Errorf("unable to load manifest: %w", err)
	}

	prune, err := strconv.ParseBool(os.Getenv("INPUT_PRUNE"))
	if err != nil {
		return fmt.Errorf("unable to parse prune: %w", err)
	}

	token := os.Getenv("INPUT_TOKEN")
	if len(token) == 0 {
		token = os.Getenv("GITHUB_TOKEN")
	}
	client := github.NewClient(token)

	repository := os.Getenv("INPUT_REPOSITORY")
	if len(repository) == 0 {
		repository = os.Getenv("GITHUB_REPOSITORY")
	}

	// Doesn't run concurrently to avoid GitHub API rate limit.
	for _, r := range strings.Split(repository, "\n") {
		if len(r) == 0 {
			continue
		}

		s := strings.Split(r, "/")
		if len(s) != 2 {
			err = multierr.Append(err, fmt.Errorf("invalid repository: %s", repository))
		}
		owner, repo := s[0], s[1]

		if err := client.SyncLabels(ctx, owner, repo, labels, prune); err != nil {
			err = multierr.Append(err, fmt.Errorf("unable to sync labels: %w", err))
		}
	}

	return err
}
