# pefi-cli

[![Build Status](https://travis-ci.org/SimonSchneider/pefi-cli.svg?branch=master)](https://travis-ci.org/SimonSchneider/pefi-cli) [![Go Report Card](https://goreportcard.com/badge/github.com/simonschneider/dyntab)](https://goreportcard.com/report/github.com/simonschneider/dyntab)

PErsonal FInance cli client it is a simple clie client for the pefi api. It has some basic functionality for getting, listing, adding, modifying against the pefi api.

The output is both table or json and graphs, queries towards the API can also be added.

## TODO

* Add support for reading from `stdin` when putting a dash (`-`) in the `--file` flag (`-f -`). This will read from `stdin` so that it can be used in bash scripts (very good for adding massive data to the API).
