package main

import (
	"fmt"
	"strings"

	"github.com/mmcdole/gofeed/json"
)

func main() {

	feedData := `{"version":"1.0", "home_page_url": "https://daringfireball.net"}`
	fp := json.Parser{}
	jsonFeed, _ := fp.Parse(strings.NewReader(feedData))
	fmt.Println(jsonFeed.Version)
}
