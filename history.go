package main

import (
	"encoding/json"
	"log"
	"net/http"
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
