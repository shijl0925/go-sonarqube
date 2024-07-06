package main

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube"
	"github.com/shijl0925/go-sonarqube/sonarqube/issues"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"log"
)

func main() {
	ctx := context.Background()
	baseUrl := "https://next.sonarqube.com/sonarqube"
	client := sonarqube.NewClient(baseUrl, "", "", nil)

	req := issues.SearchRequest{}
	p := paging.Params{
		Ps: 25,
		P:  1,
	}

	res, resp, err := client.Issues.Search(ctx, req, p)
	if err != nil {
		log.Fatalf("could not search issues: %+v", err)
	}

	fmt.Printf("Response status code: %d\n", resp.StatusCode)

	fmt.Printf("%+v\n", res)
}
