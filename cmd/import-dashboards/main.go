package main

/* This is a simple example of usage of Grafana client
for importing dashboards from a bunch of JSON files (current dir used).
You are can export dashboards with backup-dashboards utitity.
NOTE: old dashboards with same names will be silently overrided!

Usage:
  import-dashboards http://grafana.host:3000 api-key-string-here

You need get API key with Admin rights from your Grafana!
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/grafov/autograf/client"
	"github.com/grafov/autograf/grafana"
)

func main() {
	var (
		filesInDir []os.FileInfo
		rawBoard   []byte
		err        error
	)
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "Usage: import-dashboards http://grafana.host:3000 api-key-string-here\n")
		os.Exit(0)
	}
	c := client.New(os.Args[1], os.Args[2])
	filesInDir, err = ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range filesInDir {
		if strings.HasSuffix(file.Name(), ".json") {
			if rawBoard, err = ioutil.ReadFile(file.Name()); err != nil {
				log.Println(err)
				continue
			}
			var board grafana.Board
			if err = json.Unmarshal(rawBoard, &board); err != nil {
				log.Println(err)
				continue
			}
			c.DeleteDashboard(board.UpdateSlug())
			if err = c.SetDashboard(board, false); err != nil {
				log.Printf("error on importing dashboard %s", board.Title)
				continue
			}
		}
	}
}
