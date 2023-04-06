package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetTilemap(url string) TileProperties {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

	var imgProps TileProperties
	err = xml.Unmarshal(bodyText, &imgProps)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return TileProperties{}
	}

	return imgProps
}

func PrintTilemap(imgProps TileProperties) {
	fmt.Printf("Width: %d\n", imgProps.Width)
	fmt.Printf("Height: %d\n", imgProps.Height)
	fmt.Printf("NumTiles: %d\n", imgProps.NumTiles)
	fmt.Printf("NumImages: %d\n", imgProps.NumImages)
	fmt.Printf("Version: %s\n", imgProps.Version)
	fmt.Printf("TileSize: %d\n", imgProps.TileSize)
}
