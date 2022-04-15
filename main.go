package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	client := NewClient(ctx)
	team := client.FetchTeam(ctx)
	space := client.FetchSpace(ctx, team.ID)
	folder := client.FetchFolder(ctx, space.ID)
	curList := client.FetchCurrentSplintList(ctx, folder.ID)
	fmt.Println("Current spliint is", curList.Name)
}
