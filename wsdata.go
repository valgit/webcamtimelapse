package main

import "time"

/*
 * data struct ...
 */
type Image struct {
	ID        string    `json:"id"`
	Source    string    `json:"source"`
	Date      time.Time `json:"date"`
	Thumbnail string    `json:"thumbnail"`
}

type TileProperties struct {
	Width     int    `xml:"WIDTH,attr"`
	Height    int    `xml:"HEIGHT,attr"`
	NumTiles  int    `xml:"NUMTILES,attr"`
	NumImages int    `xml:"NUMIMAGES,attr"`
	Version   string `xml:"VERSION,attr"`
	TileSize  int    `xml:"TILESIZE,attr"`
}
