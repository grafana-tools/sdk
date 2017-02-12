<!--*- mode:markdown;mode:orgtbl -*-->

# Grafana SDK [![Go Report Card](https://goreportcard.com/badge/github.com/grafana-tools/sdk)](https://goreportcard.com/report/github.com/grafana-tools/sdk)

SDK for Go language offers a library for interacting with
[Grafana](http://grafana.org) server from Go applications.  It
realizes many of
[HTTP REST API](http://docs.grafana.org/reference/http_api) calls for
administration, client, organizations. Beside of them it allows
creating of Grafana objects (dashboards, panels, datasources) locally
and manipulating them for constructing dashboards programmatically.
It would be helpful for massive operations on a large set of
dashboards for example.

It was made foremost for
[autograf](https://github.com/grafana-tools/autograf) project but
later separated from it and moved to this new repository because the
library is useful per se.

Grafana operates with Javascript objects on client side so on first
view Go language looks alien thing here.  And Grafana has GUI with
detailed options for panel customization so in many cases you don't
need additional automatization.  But in situations when you operates
on hundreds of dashboards programming generation of them become not
bad idea.  And SDK that allow you import/export, create, modify and
validate Grafana structures is very helpful.  Golang is good enough
choice for operations with JSON though it may be subject of discuss.
Positives of this choice is strong typization in Go that help validate
objects alongside with high speed of execution and nice concurrency
patterns.  Negative aspect the same: the strong typization that add
more verbosity to JSON parsing in comparing with Javascript or for an
example with scripting languages like Python.  But with SDK you
already have ready for use structures and methods so generation of
JSONs become simple.  Anyway Grafana server made in Golang that prove
concept for applicability of Go for that kind of tasks.

And of course if you write applications in Golang and integrate them
with Grafana then client SDK for Go will be uniquely useful.

## Library design principles

1. SDK offers client functionality so it covers Grafana REST API with
   its requests and responses as close as possible.
1. SDK maps Grafana objects (dashboard, row, panel, datasource) to
   similar Go structures but not follows exactly all Grafana
   abstractions.
1. It doesn't use logging, instead API functions can return errors
   where it need.
1. No external deps except Go stdlib. Another exception is URL
   slugify, SDK uses external lib "slug" for algorithm compatibility —
   that is the same package that Grafana server uses.

## Examples [![GoDoc](https://godoc.org/github.com/grafana-tools/sdk?status.svg)](https://godoc.org/github.com/grafana-tools/sdk)

```go
	board := sdk.NewBoard("Sample dashboard title")
	board.ID = 1
	row1 := board.AddRow("Sample row title")
	row1.Add(sdk.NewGraph("Sample graph"))
	graph := sdk.NewGraph("Sample graph 2")
	target := sdk.Target{
		RefID:      "A",
		Datasource: "Sample Source 1",
		Expr:       "sample request 1"}
	graph.AddTarget(&target)
	row1.Add(graph)
	c := sdk.NewClient("http://grafana.host", "grafana-api-key", sdk.DefaultHTTPClient)	
	if err = c.SetDashboard(board, false); err != nil {
		fmt.Printf("error on uploading dashboard %s", board.Title)
	}
```	

The library includes several demo apps for showing API usage:

* [backup-dashboards](cmd/backup-dashboards) — saves all your dashboards as JSON-files.
* [backup-datasources](cmd/backup-datasources) — saves all your datasources as JSON-files.
* [import-datasources](cmd/import-datasources) — imports datasources from JSON-files.
* [import-dashboards](cmd/import-dashboards) — imports dashboards from JSON-files.

You need Grafana API key with _admin rights_ for using these utilities.

## Installation [![Build Status](https://travis-ci.org/grafana-tools/sdk.svg?branch=master)](https://travis-ci.org/grafana-tools/sdk)

Of course Go development environment should be set up first. Then:

    go get github.com/grafana-tools/sdk

Single external dependency required:

    go get github.com/gosimple/slug

The "slugify" for URLs is a simple task but this package used in
Grafana server so it used here for compatibility reasons.

## Grafana server compability

Made mostly for Grafana 3.x, works with Grafana 4.x but need more
tests. Full support for Grafana 4.x is on the way.

## Status of REST API realization [![Coverage Status](https://coveralls.io/repos/github/grafana-tools/sdk/badge.svg?branch=master)](https://coveralls.io/github/grafana-tools/sdk?branch=master)

Work on full API implementation still in progress. Currently
implemented only create/update/delete operations for dashboards and
datasources. State of support for misc API parts noted below.

<!--- 
#+ORGTBL: SEND status orgtbl-to-gfm
| API                    | Status          |
|------------------------+-----------------|
| Authorization          | only API tokens |
| Dashboards             | partially       |
| Datasources            | +               |
| Organization (current) | partially       |
| Organizations          | -               |
| Users                  | partially       |
| User (actual)          | partially       |
| Snapshots              | -               |
| Frontend settings      | -               |
| Admin                  | -               |
-->

<!--- BEGIN RECEIVE ORGTBL status -->
| API | Status |
|---|---|
| Authorization | only API tokens |
| Dashboards | partially |
| Datasources | + |
| Organization (current) | partially |
| Organizations | - |
| Users | partially |
| User (actual) | partially |
| Snapshots | - |
| Frontend settings | - |
| Admin | - |
<!--- END RECEIVE ORGTBL status -->

## Roadmap

* `[DONE]` Realize data structures used in a default Grafana installation for data visualizing (dashboards, datasources, panels, variables, annotations).
* `[PROGRESS]` Support all functions of Grafana REST API for manipulating dashboards and datasources.
* Support functions of Grafana REST API for manipulating users and organizations.


## Collection of Grafana tools in Golang

* [github.com/grafana/grafana-api-golang-client](https://github.com/grafana/grafana-api-golang-client) — official golang client of Grafana project. Currently in realizes parts of the REST API.
* [github.com/adejoux/grafanaclient](https://github.com/adejoux/grafanaclient) — API to manage Grafana 2.0 datasources and dashboards. It lacks features from 2.5 and later Grafana versions.
* [github.com/mgit-at/grafana-backup](https://github.com/mgit-at/grafana-backup) — just saves dashboards localy.
* [github.com/raintank/memo](https://github.com/raintank/memo) — send slack mentions to Grafana annotations.
* [github.com/retzkek/grafctl](https://github.com/retzkek/grafctl) — backup/restore/track dashboards with git.

