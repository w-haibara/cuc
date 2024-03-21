package client

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/raksul/go-clickup/clickup"
)

var apiTokenFilePath = defaultApiTokenFilePath()

func defaultApiTokenFilePath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		panic(fmt.Sprint("faild to get user config dir: ", err.Error()))
	}

	return filepath.Join(dir, "cuc", "api_token.txt")
}

type Client struct {
	*clickup.Client
	User  *clickup.User
	Teams []clickup.Team
}

func NewClient(ctx context.Context) (Client, error) {
	b, err := os.ReadFile(apiTokenFilePath)
	if err != nil {
		return Client{}, err
	}

	str := string(b)
	token := str[:strings.Index(str, "\n")]

	client, err := NewClientWithoutFallback(ctx, token)
	if err != nil {
		return Client{}, err
	}

	return client, nil
}

func NewClientWithoutFallback(ctx context.Context, token string) (Client, error) {
	client := clickup.NewClient(nil, token)

	teams, _, err := client.Authorization.GetAuthorizedTeams(ctx)
	if err != nil {
		return Client{}, err
	}

	user, _, err := client.Authorization.GetAuthorizedUser(ctx)
	if err != nil {
		return Client{}, err
	}

	return Client{client, user, teams}, nil
}

func SetupApiToken() error {
	if _, err := os.Stat(apiTokenFilePath); err != nil {
		if err := os.MkdirAll(filepath.Dir(apiTokenFilePath), 0755); err != nil {
			return err
		}
	}

	var qs = []*survey.Question{
		{
			Name:     "ApiToken",
			Prompt:   &survey.Password{Message: "API Token"},
			Validate: survey.Required,
		},
	}

	answers := struct {
		ApiToken string
	}{}

	if err := survey.Ask(qs, &answers); err != nil {
		return err
	}

	if err := os.WriteFile(apiTokenFilePath, []byte(answers.ApiToken+"\n"), 0666); err != nil {
		return err
	}

	return nil
}
