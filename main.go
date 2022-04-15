package main

import (
	"context"
	"fmt"
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
