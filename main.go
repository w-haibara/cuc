package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/k0kubun/pp"
	"github.com/raksul/go-clickup/clickup"
)

func main() {
	ctx := context.Background()
	client := NewClient(ctx)
	teamID := client.FetchTeamID(ctx)
	spaceID := client.FetchSpaceID(ctx, teamID)
	folderID := client.FetchFolderID(ctx, spaceID)
	curList := client.FetchCurrentSplintListID(ctx, folderID)
	fmt.Println("Current spliint is", curList.Name)
}

func ConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}

	dir := filepath.Join(home, ".config", "cuc")
	if _, err := os.Stat(dir); err != nil {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	return dir
}

func ReadAPIKey(configDir string) string {
	type tml struct {
		Key string
	}
	t := new(tml)
	if _, err := toml.DecodeFile(filepath.Join(configDir, "key.toml"), t); err != nil {
		panic(err.Error())
	}
	return t.Key
}

type Config struct {
	Team         string
	Space        string
	Folder       string
	SplintFormat string `toml:"splint_format"`
}

func ReadConfig(configDir string) Config {
	config := new(Config)
	if _, err := toml.DecodeFile(filepath.Join(configDir, "config.toml"), config); err != nil {
		panic(err.Error())
	}
	pp.Println(config)
	fmt.Println()
	return *config
}

type Client struct {
	*clickup.Client
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

	fmt.Println("User:", user.Username)

	config := ReadConfig(configDir)
	return Client{client, config}
}

func (c Client) FetchTeamID(ctx context.Context) string {
	teams, _, err := c.Teams.GetTeams(ctx)
	if err != nil {
		panic(err.Error())
	}

	id := ""
	for _, team := range teams {
		if team.Name == c.Config.Team {
			fmt.Println("Team:", team.Name)
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
			fmt.Println("Space:", space.Name)
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
		if folder.Name == c.Config.Folder {
			fmt.Println("Folder:", folder.Name)
			id = folder.ID
		}
	}

	if id == "" {
		panic(fmt.Sprintln("folder not found:", c.Config.Folder))
	}

	return id
}

func (c Client) FetchCurrentSplintListID(ctx context.Context, folderID string) clickup.List {
	lists, _, err := c.Lists.GetLists(ctx, folderID, false)
	if err != nil {
		panic(err.Error())
	}

	for i := len(lists) - 1; i >= 0; i-- {
		if isCurrentSplint(lists[i].Name, c.Config.SplintFormat, time.Now()) {
			return lists[i]
		}
	}

	panic("current splint not found")
}

func isCurrentSplint(listName string, layout string, date time.Time) bool {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	arr := strings.SplitN(listName, "(", 2)
	if len(arr) != 2 {
		log.Println("isCurrentSplint 1", listName)
		return false
	}
	str := strings.TrimSuffix(arr[1], ")")
	arr2 := strings.SplitN(str, " - ", 2)
	if len(arr2) != 2 {
		log.Println("isCurrentSplint 2", str)
		return false
	}

	t1, err := time.ParseInLocation(layout, strings.TrimSpace(arr2[0]), date.Location())
	if err != nil {
		log.Println("isCurrentSplint 3", err.Error())
		return false
	}
	if t1.Year() == 0 {
		t1 = t1.AddDate(date.Year(), 0, 0)
	}

	t2, err := time.ParseInLocation(layout, strings.TrimSpace(arr2[1]), date.Location())
	if err != nil {
		log.Println("isCurrentSplint 4", err.Error())
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
