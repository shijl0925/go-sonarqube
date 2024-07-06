package main

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube"
	"github.com/shijl0925/go-sonarqube/sonarqube/duplications"
	"log"
)

func main() {
	ctx := context.Background()
	baseUrl := "https://next.sonarqube.com/sonarqube"
	client := sonarqube.NewClient(baseUrl, "", "", nil)
	req := duplications.ShowRequest{
		Key: "sonar-scanner-azdo-sc:src/common/sonarqube-v5/helpers/api.ts",
	}
	res, err := client.Duplications.Show(ctx, req)

	if err != nil {
		log.Fatalf("could not show duplications: %+v", err)
	}

	fmt.Printf("%+v\n", res)
}
