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
	key := "SonarSource_echoes-react_AYvOIyNg-JQvdKIPB6Ig:stories/Icons-stories.tsx"
	baseUrl := "https://next.sonarqube.com/sonarqube"
	client := sonarqube.NewClient(baseUrl, "", "", nil)
	req := sources.RawRequest{Key: key}

	res, err := client.Sources.Raw(ctx, req)
	if err != nil {
		log.Fatalf("could not get source rsw: %+v", err)
	}

	fmt.Printf("%+v\n", *res)
}
