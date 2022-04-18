package main

import (
	"fmt"
	"log"

	"github.com/Juby210/gplayapi-go"
)

const sessionFile = "_session.arm64_v8a.json"

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
	if split := deliveryData.SplitDeliveryData; split != nil {
		for _, s := range split {
			fmt.Printf("%s %s\n", s.GetName(), s.GetDownloadUrl())
		}
	}
}
