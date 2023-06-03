// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package exporter

import (
	"context"
	"errors"
	"os"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
	"github.com/spf13/viper"
)

type githubClient struct {
	client *github.Client
}


// 1. New client
func NewClient() (*githubClient, error) {
	// Read the yaml configuration file
	viper.SetConfigName("github-key") // The file name is github-key.yaml
	// The file path is the  ~/.config/kubecub/github-key.yaml
	viper.AddConfigPath("$HOME/.config/kubecub")
	err := viper.ReadInConfig() // Read configuration file
	if err != nil {
		return nil, err
	}

	// Getting the token
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		token = viper.GetString("github.github_token")
		if token == "" {
			return nil, errors.New("missing GITHUB_TOKEN")
		}
	}

	// 创建 github 客户端
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
