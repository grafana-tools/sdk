# Autograf is a dashboard constructor for Grafana [![Build Status](https://drone.io/github.com/grafov/autograf/status.png)](https://drone.io/github.com/grafov/autograf/latest)

[Grafana](http://grafana.org) is flexible and usable for exploring and visualizing data. But UI of Grafana is not very suitable for repetitive operations with large number of objects on multiple dashboards. Aim of Autograf project is help with maintaining a large set of dashboards and datasources in an automated way. Autograf will not try to be a replacement for native Grafana methods of automation (templating variables, repeatable panels and scripted dashboards) but it complement them with own way.

This project is in early stage of development. Firstly it offers a way for processing dashboards in Go apps and interacting with Grafana instances. This part of project is already usable. Secondly it will offer DSL for constructing dashboards. I think plain blocks of text without complex nesting will enough for representing Grafana board-row-panel concept.

Currently there are two packages:

* `grafana` [![GoDoc](https://godoc.org/github.com/grafov/autograf/grafana?status.svg)](https://godoc.org/github.com/grafov/autograf/grafana) defines structures of Grafana and it may be used separately in Go apps for custom dashboards handling.
* `client` [![GoDoc](https://godoc.org/github.com/grafov/autograf/client?status.svg)](https://godoc.org/github.com/grafov/autograf/client) realizes [HTTP REST API](http://docs.grafana.org/reference/http_api). It also may be used separately for integrating Go apps with Grafana. It uses `grafana` package for keeping loaded dashboards and defines its own types for keeping users/orgs and other auxilary structures used in Grafana API.

DSL part is unfinished yet so it will be published later.

## Demo utilities

Autograf includes several of demo apps for show how to use `client` and `grafana` API:

* [backup-dashboards](cmd/backup-dashboards) — saves all your dashboards as JSON-files.
* [backup-datasources](cmd/backup-datasources) — saves all your datasources as JSON-files.
* [import-datasources](cmd/import-datasources) — imports datasources from JSON-files.
* [import-dashboards](cmd/import-dashboards) — imports dashboards from JSON-files.

You need Grafana API key with _admin rights_ for using these utilities.

## Thoughts about DSL

Work on DSL syntax and translator in progress not much yet to say about it. I want something simple in plain text
for describing dashboards instead of mapping them 1:1 to Grafana objects. Short sample how it may look (syntax may be
changed):

    # Example of a board with a panel
    board The sample dashboard

	# Example how define new sources in the loop:
	repeat (0..3) as $srv:
      source PromSrc$srv
      prometheus http://127.1:9090

	# This example defines the set of panels with different sources.
	var n = 0
	for $handlers as handler:
	  $n++
	  panel Sample number $n
	  type graph
	  query my_response_time_metric{instance="host$n",handler="$handler"}
	  if $n < 4:
	    datasource PromSrc$n
	  else:
		datasource Prom0


## Installation [![Build Status](https://travis-ci.org/grafov/autograf.svg?branch=master)](https://travis-ci.org/grafov/autograf)

For use in your Go apps just install packages separately:

    go get github.com/grafov/autograf/grafana
    go get github.com/grafov/autograf/client

Single external dependency required for `grafana` package:

    go get github.com/gosimple/slug

— "slugify" URLs is a simple task but this package used in Grafana server so it used
here for compatibility reasons.

Sample utility that just saves all dashboards from Grafana to JSON files in a current dir:

    go install github.com/grafov/autograf/grafana/cmd/backup-dashboards

## Roadmap [![Coverage Status](https://coveralls.io/repos/github/grafov/autograf/badge.svg?branch=master)](https://coveralls.io/github/grafov/autograf?branch=master)

### Major targets

* `[DONE]` Realize data structures used in a default Grafana installation for data visualizing (dashboards, datasources, panels, variables, annotations).
* `[PROGRESS]` Support all functions of Grafana REST API for manipulating dashboards and datasources.
* `[PROGRESS]` Realize DSL for defining dashboards in a plain text format.
* Import dashboards or single panels from running Grafana instances and convert them to DSL.

### Minor targets

* Support functions of Grafana REST API for manipulating users and organizations.

## Collection of Grafana tools in Golang

* [github.com/grafana/grafana-api-golang-client](https://github.com/grafana/grafana-api-golang-client) — official golang client of Grafana project. Currently in realizes parts of the REST API.
* [github.com/adejoux/grafanaclient](https://github.com/adejoux/grafanaclient) — API to manage Grafana 2.0 datasources and dashboards. It lacks features from 2.5 and later Grafana versions.
* [github.com/mgit-at/grafana-backup](https://github.com/mgit-at/grafana-backup) — just saves dashboards localy.
* [github.com/raintank/memo](https://github.com/raintank/memo) — send slack mentions to Grafana annotations.
* [github.com/retzkek/grafctl](https://github.com/retzkek/grafctl) — backup/restore/track dashboards with git.

### Projects offered DSL or helper tools for Grafana in other languages

* [github.com/jakubplichta/grafana-dashboard-builder](https://github.com/jakubplichta/grafana-dashboard-builder) Python tool for building Grafana dashboards in YAML.
* [github.com/m110/grafcli](https://github.com/m110/grafcli) Python tool for managing Grafana in CLI. It querying Grafana backends directly. The project abandoned in alpha state for a long time.
