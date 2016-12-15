# Grafana SDK

*These libraries just moved out from https://github.com/grafov/autograf repository.*
*Paths not fixed yet so they in non working state!*

SDK for Go language offers a way for interacting with [Grafana](http://grafana.org) server from Go applications.

Currently it consists of two packages:

* `grafana` [![GoDoc](https://godoc.org/github.com/grafov/autograf/grafana?status.svg)](https://godoc.org/github.com/grafov/autograf/grafana) defines structures of Grafana and it may be used separately in Go apps for custom dashboards handling.
* `client` [![GoDoc](https://godoc.org/github.com/grafov/autograf/client?status.svg)](https://godoc.org/github.com/grafov/autograf/client) realizes [HTTP REST API](http://docs.grafana.org/reference/http_api). It also may be used separately for integrating Go apps with Grafana. It uses `grafana` package for keeping loaded dashboards and defines its own types for keeping users/orgs and other auxilary structures used in Grafana API.

## Demo utilities

The library includes several demo apps for showing how to use `client` and `grafana` API:

* [backup-dashboards](cmd/backup-dashboards) — saves all your dashboards as JSON-files.
* [backup-datasources](cmd/backup-datasources) — saves all your datasources as JSON-files.
* [import-datasources](cmd/import-datasources) — imports datasources from JSON-files.
* [import-dashboards](cmd/import-dashboards) — imports dashboards from JSON-files.

You need Grafana API key with _admin rights_ for using these utilities.

## Installation [![Build Status](https://travis-ci.org/grafov/autograf.svg?branch=master)](https://travis-ci.org/grafov/autograf) [![Build Status](https://drone.io/github.com/grafov/autograf/status.png)](https://drone.io/github.com/grafov/autograf/latest)

For use in your Go apps just install packages separately:

    go get github.com/grafana-tools/grafana
    go get github.com/grafana-tools/client

Single external dependency required for `grafana` package:

    go get github.com/gosimple/slug

— "slugify" URLs is a simple task but this package used in Grafana server so it used
here for compatibility reasons.

## Roadmap [![Coverage Status](https://coveralls.io/repos/github/grafov/autograf/badge.svg?branch=master)](https://coveralls.io/github/grafov/autograf?branch=master)

* `[DONE]` Realize data structures used in a default Grafana installation for data visualizing (dashboards, datasources, panels, variables, annotations).
* `[PROGRESS]` Support all functions of Grafana REST API for manipulating dashboards and datasources.
* Support functions of Grafana REST API for manipulating users and organizations.

## Collection of Grafana tools in Golang

* [github.com/grafana/grafana-api-golang-client](https://github.com/grafana/grafana-api-golang-client) — official golang client of Grafana project. Currently in realizes parts of the REST API.
* [github.com/adejoux/grafanaclient](https://github.com/adejoux/grafanaclient) — API to manage Grafana 2.0 datasources and dashboards. It lacks features from 2.5 and later Grafana versions.
* [github.com/mgit-at/grafana-backup](https://github.com/mgit-at/grafana-backup) — just saves dashboards localy.
* [github.com/raintank/memo](https://github.com/raintank/memo) — send slack mentions to Grafana annotations.
* [github.com/retzkek/grafctl](https://github.com/retzkek/grafctl) — backup/restore/track dashboards with git.

