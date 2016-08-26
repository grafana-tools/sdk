package main

/* This is a simple example of usage of Grafana client
for copying datasources and saving them to a disk.
It is useful for Grafana backups!

Usage:
  backup-datasources http://grafana.host:3000 api-key-string-here

You need get API key with Admin rights from your Grafana!
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gosimple/slug"
	"github.com/grafov/autograf/client"
	"github.com/grafov/autograf/grafana"
)

func main() {
	var (
		datasources []grafana.Datasource
		dsPacked    []byte
		meta        client.BoardProperties
		err         error
	)
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "Usage:  backup-datasources http://grafana.host:3000 api-key-string-here\n")
		os.Exit(0)
	}
	c := client.New(os.Args[1], os.Args[2], client.DefaultHTTPClient)
	if datasources, err = c.GetAllDatasources(); err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", err))
		os.Exit(1)
	}
	for _, ds := range datasources {
		if dsPacked, err = json.Marshal(ds); err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("%s for %s\n", err, ds.Name))
			continue
		}
		if err = ioutil.WriteFile(fmt.Sprintf("%s.json", slug.Make(ds.Name)), dsPacked, os.FileMode(int(0666))); err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("%s for %s\n", err, meta.Slug))
		}
	}
}
