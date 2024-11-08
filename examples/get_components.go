package main

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube"
	"github.com/shijl0925/go-sonarqube/sonarqube/components"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"log"
)

func main() {
	ctx := context.Background()
	baseUrl := "https://next.sonarqube.com/sonarqube"
	client := sonarqube.NewClient(baseUrl, "", "", nil)
	req := components.SearchRequest{
		Qualifiers: "TRK",
	}
	p := paging.Params{
		Ps: 100,
		P:  1,
	}
	res, resp, err := client.Components.Search(ctx, req, p)

	//res, err := client.Components.SearchAll(ctx, req)
	if err != nil {
		log.Fatalf("could not search components: %+v", err)
	}

	fmt.Printf("Response status code: %d\n", resp.StatusCode)

	fmt.Printf("%+v\n", res)
}
