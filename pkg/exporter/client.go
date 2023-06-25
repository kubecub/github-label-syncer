// Copyright © 2023 KubeCub open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package exporter

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/go-github/v28/github"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type githubClient struct {
	client *github.Client
}

// NewClient creates a new github client.
func NewClient() (*githubClient, error) {
	// Check if .env file exists, if not, remind user to copy .env.template
	envPath := filepath.Join(".", ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		envTemplatePath := filepath.Join(".", ".env.template")
		fmt.Println("envPath: ", envPath)
		if _, err := os.Stat(envTemplatePath); os.IsNotExist(err) {
			return nil, errors.New("missing .env.template file")
		}
		cmd := exec.Command("cp", envTemplatePath, envPath)
		if err := cmd.Run(); err != nil {
			return nil, errors.New("failed to copy .env.template to .env file")
		}
	}

	// Load environment variables from .env file
	if err := godotenv.Load(envPath); err != nil {
		return nil, errors.New("failed to load .env file")
	}

	// Getting the token
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		// If the token is not defined in the .env file, use the configuration file
		viper.SetConfigName("github-key") // The file name is github-key.yaml
		// The file path is the  ~/.config/kubecub/github-key.yaml
		token = viper.GetString("github.github_token")
	}

	// Creating a github client
	cli := newClient(token)
	return &githubClient{
		client: cli,
	}, nil
}

func newClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})
	tc := oauth2.NewClient(context.Background(), ts)
	return github.NewClient(tc)
}
