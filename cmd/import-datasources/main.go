package main

/* This is a simple example of usage of Grafana client
for importing datasources from a bunch of JSON files (current dir used).
You are can export datasources with backup-datasources utitity.
NOTE: old datasources with same names will be silently overrided!

Usage:
  import-datasousces http://grafana.host:3000 api-key-string-here

You need get API key with Admin rights from your Grafana!
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/grafov/autograf/client"
	"github.com/grafov/autograf/grafana"
	"strings"
)

func main() {
	var (
		datasources []grafana.Datasource
		filesInDir  []os.FileInfo
		rawDS       []byte
		status      client.StatusMessage
		err         error
	)
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "Usage:  import-datasources http://grafana.host:3000 api-key-string-here\n")
		os.Exit(0)
	}
	c := client.New(os.Args[1], os.Args[2])
	if datasources, err = c.GetAllDatasources(); err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", err))
		os.Exit(1)
	}
	filesInDir, err = ioutil.ReadDir(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", err))
	}
	for _, file := range filesInDir {
		if strings.HasSuffix(file.Name(), ".json") {
			if rawDS, err = ioutil.ReadFile(file.Name()); err != nil {
				fmt.Fprint(os.Stderr, fmt.Sprintf("%s\n", err))
				continue
			}
			var newDS grafana.Datasource
			if err = json.Unmarshal(rawDS, &newDS); err != nil {
				fmt.Fprint(os.Stderr, fmt.Sprintf("%s\n", err))
				continue
			}
			for _, existingDS := range datasources {
				if existingDS.Name == newDS.Name {
					c.DeleteDatasource(existingDS.ID)
					break
				}
			}
			if status, err = c.CreateDatasource(newDS); err != nil {
				fmt.Fprint(os.Stderr, fmt.Sprintf("error on importing datasource %s with %s (%s)", newDS.Name, err, status.Message))
			}
		}
	}
}
