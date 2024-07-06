package main

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube"
	"github.com/shijl0925/go-sonarqube/sonarqube/server"
	"log"
)

func main() {
	ctx := context.Background()
	baseUrl := "https://next.sonarqube.com/sonarqube"
	client := sonarqube.NewClient(baseUrl, "", "", nil)
	req := server.VersionRequest{}

	res, err := client.Server.Version(ctx, req)
	if err != nil {
		log.Fatalf("could not get version: %+v", err)
	}
	fmt.Printf("version: %s\n", *res)
}
