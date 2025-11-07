package main

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube"
	"github.com/shijl0925/go-sonarqube/sonarqube/sources"
	"log"
)

func main() {
	ctx := context.Background()
	key := "SonarSource_echoes-react:src/index.ts"
	baseUrl := "https://next.sonarqube.com/sonarqube"
	client := sonarqube.NewClient(baseUrl, "", "", nil)
	req := sources.RawRequest{Key: key}

	res, resp, err := client.Sources.Raw(ctx, req)
	if err != nil {
		log.Fatalf("could not get source rsw: %+v", err)
	}

	fmt.Printf("Response status code: %d\n", resp.StatusCode)

	fmt.Printf("%+v\n", *res)
}
