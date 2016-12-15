# Grafana SDK

*These libraries just moved out from https://github.com/grafov/autograf repository.*
*Badges of external services broken yet because of old paths!*

SDK for Go language offers a library for interacting with [Grafana](http://grafana.org) server from Go applications.
It realizes the many of [HTTP REST API](http://docs.grafana.org/reference/http_api) calls but beside them
it allows create Grafana objects (dashboards, panels, datasources) locally and manipulate them for
constructing dashboards programmatically.
It was made foremost for [autograf](https://github.com/grafana-tools/autograf) project but later separated 
from it and moved to this new repository because the library is useful per se.

The library includes several demo apps for showing API usage:

* [backup-dashboards](cmd/backup-dashboards) — saves all your dashboards as JSON-files.
* [backup-datasources](cmd/backup-datasources) — saves all your datasources as JSON-files.
* [import-datasources](cmd/import-datasources) — imports datasources from JSON-files.
* [import-dashboards](cmd/import-dashboards) — imports dashboards from JSON-files.

You need Grafana API key with _admin rights_ for using these utilities.

## Installation [![Build Status](https://travis-ci.org/grafov/autograf.svg?branch=master)](https://travis-ci.org/grafov/autograf) [![Build Status](https://drone.io/github.com/grafov/autograf/status.png)](https://drone.io/github.com/grafov/autograf/latest)

For use in your Go apps just install packages separately:

    go get github.com/grafana-tools/sdk/grafana
    go get github.com/grafana-tools/sdk/client

Single external dependency required for `grafana` package:

    go get github.com/gosimple/slug

— "slugify" URLs is a simple task but this package used in Grafana server so it used
here for compatibility reasons.

## Status of REST API realization

Currently implemented only create/update/delete operations for dashboards and datasources. Other functions in progress. State of misc API parts noted below.

### Authorization

Only API tokens implemented.

### Dashboards

Partially implemented.

### Datasources

Partially implemented.

### Organizations

Not implemented.

### Users

Partially implemented.

### Snapshots

Not implemented.

### Frontend settings

Not implemented.

### Login

Not implemented.

### Admin

Not implemented.


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

