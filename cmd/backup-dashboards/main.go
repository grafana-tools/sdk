package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/grafov/autograf/client"
)

/* This is a simple example of usage of Grafana client
for copying dashboards and saving them to a disk.
Ie really useful for Grafana backups!

Usage:
  backup-dashboards http://grafana.host:3000 api-key-string-here
*/

func main() {
	var (
		boardLinks []client.FoundBoard
		rawBoard   []byte
		meta       client.BoardProperties
		err        error
	)
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "Usage:  backup-dashboards http://grafana.host:3000 api-key-string-here\n")
		os.Exit(0)
	}
	c := client.New(os.Args[1], os.Args[2])
	if boardLinks, err = c.SearchDashboards("", false); err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", err))
		os.Exit(1)
	}
	for _, link := range boardLinks {
		if rawBoard, meta, err = c.GetRawDashboard(link.URI); err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("%s for %s\n", err, link.URI))
			continue
		}
		if err = ioutil.WriteFile(fmt.Sprintf("%s.json", meta.Slug), rawBoard, os.FileMode(int(0666))); err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("%s for %s\n", err, meta.Slug))
		}
	}
}
