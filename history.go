package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetHistory() []Image {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pv.viewsurf.com/1246/Neufchatel-Hardelot?a=medias&c=5318", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://pv.viewsurf.com/1246/Neufchatel-Hardelot?i=NTMxODp1bmRlZmluZWQ")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Brave";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	/* marshal
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	*/
	var history []Image

	//use alternativ unmarshal
	err = json.NewDecoder(resp.Body).Decode(&history)
	if err != nil {
		log.Fatal(err)
	}
	return history
}

/*
 * download thumbnail
 */
func GetThumbnail(t Image) error {
	url := "https:" + t.Thumbnail
	fileName := fmt.Sprintf("%s_%s.jpg", t.Date.Format("2006-01-02"), t.ID)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading thumbnail:", err)
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	fmt.Printf("Thumbnail downloaded for tile with ID %s\n", t.ID)
	return nil
}

/*
 * download a tile
 */
func GetOneTile() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://filmspv.viewsurf.com/hardelot01/1/1/media_1680796809-tiles/TileGroup0/4-7-1.jpg", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8")
	req.Header.Set("Accept-Language", "fr-FR,fr;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Origin", "https://pv.viewsurf.com")
	req.Header.Set("Referer", "https://pv.viewsurf.com/")
	req.Header.Set("Sec-Fetch-Dest", "image")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Chromium";v="112", "Brave";v="112", "Not:A-Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
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
}
