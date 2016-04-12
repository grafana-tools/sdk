package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/grafov/autograf/client"
)

/* This is a simple example of usage of Grafana client
for copying dashboards and saving them to a disk.
Usage:
  backup-dashboards http://grafana.host:3000 api-key-string-here
*/

func main() {
	var (
		boards        []client.FoundBoard
		boardWithMeta client.BoardWithMeta
		data          []byte
		err           error
	)
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage:  backup-dashboards http://grafana.host:3000 api-key-string-here\n")
	}
	c := client.New(os.Args[1], os.Args[2])
	if boards, err = c.SearchDashboards("", false); err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", err))
		os.Exit(1)
	}
	for _, link := range boards {
		if boardWithMeta, err = c.GetDashboard(link.URI); err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("%s for %s\n", err, link.URI))
			continue
		}
		if data, err = json.MarshalIndent(boardWithMeta.Board, "", "  "); err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("%s for %s\n", err, boardWithMeta.Board.Title))
			continue
		}
		if err = ioutil.WriteFile(fmt.Sprintf("%s.json", boardWithMeta.Meta.Slug), data, os.FileMode(int(0666))); err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("%s for %s\n", err, boardWithMeta.Meta.Slug))
		}
	}
}
