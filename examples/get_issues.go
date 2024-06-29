package main

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube"
	"github.com/shijl0925/go-sonarqube/sonarqube/issues"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"log"
)

func main() {
	baseUrl := "https://next.sonarqube.com/sonarqube/api"
	client := sonarqube.NewClient(baseUrl, "", "", nil)

	req := issues.SearchRequest{}
	p := paging.Params{
		Ps: 25,
		P:  1,
	}

	res, err := client.Issues.Search(req, p)
	if err != nil {
		log.Fatalf("could not search issues: %+v", err)
	}

	fmt.Printf("%+v\n", res)
}
