package main

import (
	"fmt"
	"log"

	"github.com/Juby210/gplayapi-go"
)

const sessionFile = "_session.json"

func main() {
	client, err := gplayapi.LoadSession(sessionFile)
	if err != nil {
		client, err = gplayapi.NewClient(
			"email",
			"aasToken",
		)
		if err != nil {
			log.Fatal(err)
		}
		client.SaveSession(sessionFile)
	}
	client.SessionFile = sessionFile
	app, err := client.GetAppDetails("com.discord")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s (%d)\n", app.VersionName, app.VersionCode)
	deliveryData, err := client.Purchase(app.PackageName, app.VersionCode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deliveryData.GetDownloadUrl())
}
