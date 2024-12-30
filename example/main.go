package main

import (
	"fmt"

	"github.com/sarumaj/go-wallpaper"
)

func main() {
	background, err := wallpaper.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println("Current wallpaper:", background)

	err = wallpaper.SetFromFile("/usr/share/backgrounds/gnome/adwaita-day.jpg")
	if err != nil {
		panic(err)
	}

	err = wallpaper.SetFromURL("https://picsum.photos/1920/1080")
	if err != nil {
		panic(err)
	}
}
