package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/raksul/go-clickup/clickup"
)

type Client struct {
	*clickup.Client
	User   *clickup.User
	Config Config
}

func NewClient(ctx context.Context) Client {
	configDir := ConfigDir()
	key := ReadAPIKey(configDir)

	client := clickup.NewClient(nil, key)

	user, _, err := client.Authorization.GetAuthorizedUser(ctx)
	if err != nil {
		panic(err.Error())
	}

	config := ReadConfig(configDir)
	return Client{client, user, config}
}

func (c Client) FetchTeamID(ctx context.Context) string {
	teams, _, err := c.Teams.GetTeams(ctx)
	if err != nil {
		panic(err.Error())
	}

	id := ""
	for _, team := range teams {
		if team.Name == c.Config.Team {
			id = team.ID
			break
		}
	}

	if id == "" {
		panic(fmt.Sprintln("team not found:", c.Config.Team))
	}

	return id
}

func (c Client) FetchSpaceID(ctx context.Context, teamID string) string {
	spaces, _, err := c.Spaces.GetSpaces(ctx, teamID)
	if err != nil {
		panic(err.Error())
	}

	id := ""
	for _, space := range spaces {
		if space.Name == c.Config.Space {
			id = space.ID
			break
		}
	}

	if id == "" {
		panic(fmt.Sprintln("space not found:", c.Config.Space))
	}

	return id
}

func (c Client) FetchFolderID(ctx context.Context, spaceID string) string {
	folders, _, err := c.Folders.GetFolders(ctx, spaceID, false)
	if err != nil {
		panic(err.Error())
	}

	id := ""
	for _, folder := range folders {
		if folder.Name == c.Config.Splint.Folder {
			id = folder.ID
		}
	}

	if id == "" {
		panic(fmt.Sprintln("folder not found:", c.Config.Splint.Folder))
	}

	return id
}

func (c Client) FetchCurrentSplintListID(ctx context.Context, folderID string) clickup.List {
	lists, _, err := c.Lists.GetLists(ctx, folderID, false)
	if err != nil {
		panic(err.Error())
	}

	for i := len(lists) - 1; i >= 0; i-- {
		if isCurrentSplint(lists[i].Name, c.Config.Splint.TimeFormat, time.Now()) {
			return lists[i]
		}
	}

	panic("current splint not found")
}

func isCurrentSplint(listName string, layout string, date time.Time) bool {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	arr := strings.SplitN(listName, "(", 2)
	if len(arr) != 2 {
		return false
	}
	str := strings.TrimSuffix(arr[1], ")")
	arr2 := strings.SplitN(str, " - ", 2)
	if len(arr2) != 2 {
		return false
	}

	t1, err := time.ParseInLocation(layout, strings.TrimSpace(arr2[0]), date.Location())
	if err != nil {
		return false
	}
	if t1.Year() == 0 {
		t1 = t1.AddDate(date.Year(), 0, 0)
	}

	t2, err := time.ParseInLocation(layout, strings.TrimSpace(arr2[1]), date.Location())
	if err != nil {
		return false
	}
	if t2.Year() == 0 {
		t2 = t2.AddDate(date.Year(), 0, 0)
	}

	if (t1.Before(date) || t1.Equal(date)) && (date.Equal(t2) || date.Before(t2)) {
		return true
	}

	return false
}
