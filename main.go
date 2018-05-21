package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	fb "github.com/huandu/facebook"
)

func main() {

	type EventData struct {
		id        int
		name      string
		link      string
		likecount int64
	}

	var center = "37.176487,-3.597929"
	res, _ := fb.Get("/search", fb.Params{
		"access_token": "529285107466818|bd7qQ2XQBSvJ3hRwdb7RPg6LTGY",
		"type":         "place",
		"categories":   "[\"ARTS_ENTERTAINMENT\"]",
		"center":       center,
		"distance":     "10000",
		"limit":        "100",
		"fields":       "name,engagement,overall_star_rating,link",
	})

	var items []fb.Result
	err := res.DecodeField("data", &items)

	if err != nil {
		fmt.Println(err)
		return
	}

	var results []*EventData
	var totallikes int64
	var itemlen = int64(len(items))

	for _, item := range items {
		id, _ := strconv.Atoi(item["id"].(string))
		name := item["name"].(string)
		link := item["link"].(string)
		engagement, _ := item["engagement"].(map[string]interface{})

		likecount, _ := engagement["count"].(json.Number).Int64()
		if likecount == 0 {
			likecount = totallikes / itemlen
		} else {
			totallikes += likecount
		}

		// Append results
		results = append(results, &EventData{id, name, link, likecount})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].likecount > results[j].likecount
	})

	for _, result := range results {

		fmt.Println(*result)

	}

}
