package main

import (
	"fmt"
)

func main() {
	fmt.Printf("reading histo : \n")
	histo := GetHistory()
	for _, img := range histo {
		fmt.Printf("%s\n\tsrc: %s\n\tthumb: %s\n", img.Date, img.Source, img.Thumbnail)
	}

	fmt.Printf("tiling at : \n")
	source := "https://filmspv.viewsurf.com/hardelot01/1/1/media_1680760809-tiles/ImageProperties.xml"
	tile := GetTilemap(source)
	PrintTilemap(tile)

}
