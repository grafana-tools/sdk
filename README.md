# Autograf is a dashboard constructor for Grafana

[Grafana](http://grafana.org) is flexible and usable for exploring and visualizing data. But UI of Grafana is not very suitable for repetitive operations with large number of objects on multiple dashboards. Aim of Autograf project is help with maintaining a large set of dashboards and datasources in an automated way. Autograf will not try to be a replacement for native Grafana methods of automation (templating variables, repeatable panels and scripted dashboards) but it complement them with own way.

This project is in early stage of development. Firstly it offers a way for processing dashboards in Go apps and interacting with Grafana instances. This part of project is already usable. Secondly it will become DSL for constructing dashboards. 

Currently there are two packages:

* `grafana` [![GoDoc](https://godoc.org/github.com/grafov/autograf/grafana?status.svg)](https://godoc.org/github.com/grafov/autograf/grafana) defines structures of Grafana and it may be used separately in Go apps for custom dashboards handling. 
* `client` [![GoDoc](https://godoc.org/github.com/grafov/autograf/client?status.svg)](https://godoc.org/github.com/grafov/autograf/client) realizes [HTTP REST API](http://docs.grafana.org/reference/http_api). It also may be used separately for integrating Go apps with Grafana.

DSL part is unfinished yet so it will be published later.

# Installation

For use in your Go apps just install packages separately:

    go get github.com/grafov/autograf/grafana
    go get github.com/grafov/autograf/client

Single external dependency required for `grafana` package:

    go get github.com/gosimple/slug

— "slugify" URLs is a simple task but this package used in Grafana server so it used
here for compatibility reasons.

Sample utility that just saves all dashboards from Grafana to JSON files in a current dir:

    go install github.com/grafov/autograf/grafana/cmd/backup-dashboards

## Roadmap

### Major targets

* `[DONE]` Realize data structures used in a default Grafana installation for data visualizing (dashboards, datasources, panels, variables, annotations).
* `[PROGRESS]` Support all functions of Grafana REST API for manipulating dashboards and datasources.
* `[PROGRESS]` Realize DSL for defining dashboards in a plain text format.
* Import dashboards or single panels from running Grafana instances and convert them to DSL.

### Minor targets

* Support functions of Grafana REST API for manipulating users and organizations.

## Related works in Go

* [github.com/grafana/grafana-api-golang-client](https://github.com/grafana/grafana-api-golang-client) — official golang client of Grafana project. Currently in realizes parts of the REST API.
* [github.com/adejoux/grafanaclient](https://github.com/adejoux/grafanaclient) — API to manage Grafana 2.0 datasources and dashboards. It lacks features from 2.5 and later Grafana versions.

### Other projects offered DSL or helper tools for Grafana

* [github.com/jakubplichta/grafana-dashboard-builder](https://github.com/jakubplichta/grafana-dashboard-builder) Python tool for building Grafana dashboards in YAML.
* [github.com/m110/grafcli](https://github.com/m110/grafcli) Python tool for managing Grafana in CLI. It querying Grafana backends directly. The project abandoned in alpha state for a long time.
