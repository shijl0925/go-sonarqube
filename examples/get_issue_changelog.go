package main

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube"
	"github.com/shijl0925/go-sonarqube/sonarqube/issues"
	"log"
)

func main() {
	ctx := context.Background()
	baseUrl := "https://next.sonarqube.com/sonarqube/api"
	client := sonarqube.NewClient(baseUrl, "", "", nil)

	req := issues.ChangelogRequest{
		Issue: "63fa8f0c-0a58-42c3-b59e-2a82565e98bf",
	}

	res, err := client.Issues.Changelog(ctx, req)
	if err != nil {
		log.Fatalf("could not get authors: %+v", err)
	}

	fmt.Printf("%+v\n", res)
}
