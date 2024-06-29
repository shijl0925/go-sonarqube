package main

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube"
	"github.com/shijl0925/go-sonarqube/sonarqube/server"
	"log"
)

func main() {
	baseUrl := "https://next.sonarqube.com/sonarqube/api"
	client := sonarqube.NewClient(baseUrl, "", "", nil)
	req := server.VersionRequest{}

	res, err := client.Server.Version(req)
	if err != nil {
		log.Fatalf("could not get version: %+v", err)
	}
	fmt.Printf("version: %s\n", *res)

	fmt.Printf("%+v\n", res)
}
