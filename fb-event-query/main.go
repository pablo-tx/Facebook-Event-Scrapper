package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
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
		website   string
		address   string
	}

	var center = "37.176487,-3.597929"
	res, _ := fb.Get("/search", fb.Params{
		"access_token": "529285107466818|bd7qQ2XQBSvJ3hRwdb7RPg6LTGY",
		"type":         "place",
		"categories":   "[\"ARTS_ENTERTAINMENT\"]",
		"center":       center,
		"distance":     "10000",
		"limit":        "100",
		"fields":       "name,engagement, link, website, single_line_address",
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
		website := ""
		if item["website"] != nil {
			website = item["website"].(string)
		}
		address := ""
		if item["single_line_address"] != nil {
			address = item["single_line_address"].(string)
		}
		engagement, _ := item["engagement"].(map[string]interface{})
		likecount, _ := engagement["count"].(json.Number).Int64()
		if likecount == 0 {
			likecount = totallikes / itemlen
		} else {
			totallikes += likecount
		}

		// Append results
		results = append(results, &EventData{id, name, link, likecount, website, address})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].likecount > results[j].likecount
	})

	var out bytes.Buffer
	multi := io.MultiWriter(os.Stdout, &out)

	for _, result := range results {
		fmt.Println(*result)
		cmd := exec.Command("node", "../fb-event-scrapper/index.js", result.link)
		cmd.Stdout = multi
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
		cmd.Wait()
		//fmt.Printf("%s\n", out.String())
		out.Reset()
	}

}
