package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

func main() {

	var center = "37.176487,-3.597929"
	//searchList := []string{"Party", "Discoteca", "Club", "Pub", "Nightclub"}
	//for _, search := range searchList {
	res, _ := fb.Get("/search", fb.Params{
		"access_token": "529285107466818|bd7qQ2XQBSvJ3hRwdb7RPg6LTGY",
		"type":         "place",
		"categories":   "[\"ARTS_ENTERTAINMENT\"]",
		"center":       center,
		"distance":     "10000",
		"limit":        "1000",
	})

	var items []fb.Result

	err := res.DecodeField("data", &items)

	if err != nil {
		// err can be a Facebook API error.
		// if so, the Error struct contains error details.
		fmt.Println(err)

		return
	}

	for _, item := range items {
		fmt.Println(item["name"])
	}
	// read my last feed story.
	//fmt.Println("Sitios que contienen "+search+" :", res.Get("data"))

}

//}
