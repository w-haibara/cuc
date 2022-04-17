package main

import (
	"context"
	"fmt"

	"github.com/w-haibara/cuc/db"
)

func main() {
	ctx := context.Background()
	SaveSplintInfo(ctx)
	PrintCurrentSplintList(ctx)
}

func SaveSplintInfo(ctx context.Context) {
	client := NewClient(ctx)

	team := client.FetchTeam(ctx)
	space := client.FetchSpace(ctx, team.ID)
	folder := client.FetchSplintFolder(ctx, space.ID)
	db.Updates(db.ClickUp{
		TeamID:   team.ID,
		SpaceID:  space.ID,
		FolderID: folder.ID,
	})

}

func PrintCurrentSplintList(ctx context.Context) {
	client := NewClient(ctx)

	folderID := db.FetchFolderID()
	curList := client.FetchCurrentSplintList(ctx, folderID)
	fmt.Println("Current spliint is", curList.Name)
}
