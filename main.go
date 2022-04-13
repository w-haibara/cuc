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

	type ApiKey struct {
		Key string
	}
	t := new(ApiKey)
	if _, err := toml.DecodeFile(filepath.Join(dir, "key.toml"), t); err != nil {
		panic(err.Error())
	}

	type Config struct {
		Team         string
		Space        string
		Folder       string
		SplintFormat string `toml:"splint_format"`
	}
	config := new(Config)
	if _, err := toml.DecodeFile(filepath.Join(dir, "config.toml"), config); err != nil {
		panic(err.Error())
	}
	pp.Println(config)
	fmt.Println()

	client := clickup.NewClient(nil, t.Key)

	user, _, err := client.Authorization.GetAuthorizedUser(ctx)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("User:", user.Username)

	teams, _, err := client.Teams.GetTeams(ctx)
	if err != nil {
		panic(err.Error())
	}
	for _, team := range teams {
		if team.Name == config.Team {
			fmt.Println("Team:", team.Name)

			spaces, _, err := client.Spaces.GetSpaces(ctx, team.ID)
			if err != nil {
				panic(err.Error())
			}
			for _, space := range spaces {
				if space.Name == config.Space {
					fmt.Println("Space:", space.Name)

					folders, _, err := client.Folders.GetFolders(ctx, space.ID, false)
					if err != nil {
						panic(err.Error())
					}

					for _, folder := range folders {
						if folder.Name == config.Folder {
							fmt.Println("Folder:", folder.Name)

							lists, _, err := client.Lists.GetLists(ctx, folder.ID, false)
							if err != nil {
								panic(err.Error())
							}

							for i := len(lists) - 1; i >= 0; i-- {
								if isCurrentSplint(lists[i].Name, config.SplintFormat, time.Now()) {
									fmt.Println("Current spliint is", lists[i].Name)
									return
								}
							}
							panic("current splint not found")
						}
					}
					panic(fmt.Sprintln("folder not found:", config.Folder))
				}
			}
			panic(fmt.Sprintln("space not found:", config.Space))
		}
	}
	panic(fmt.Sprintln("team not found:", config.Team))
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
