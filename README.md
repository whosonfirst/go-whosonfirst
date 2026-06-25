# go-whosonfirst

Go package for common Who's On First operations.

## Motivation

This package is work in progress and aims to merge a whole bunch of disparate `go-whosonfirst-*` packages in to one.

The original design for the Who's On First (WOF) Go packages was to have a lot of small "composable" packages but over time this has made it difficult to reason about all of the (common) functionality available for working with WOF records. It also makes it a chore to apply updates.

Since these packages were first started Go has gotten better at excluding unused code in imports so while package itself have many dependencies only those used in any given program should be imported in your tools.

### Versioning

The "v4" version number reflects the largest version number plus one for the set of packages merged in to this one. Many of those packages will not have reached "v3" or even "v2".

## Documentation

Documentation is incomplete at this time.

## Tests

Tests are also incomplete at this time.

## Packages that have merged/superseded in to this package

* [whosonfirst/go-reader-findingaid](https://github.com/whosonfirst/go-reader-findingaid)
* [whosonfirst/go-whosonfirst-concordances](https://github.com/whosonfirst/go-whosonfirst-concordances)
* [whosonfirst/go-whosonfirst-database](https://github.com/whosonfirst/go-whosonfirst-database)
* [whosonfirst/go-whosonfirst-edtf](https://github.com/whosonfirst/go-whosonfirst-edtf)
* [whosonfirst/go-whosonfirst-export](https://github.com/whosonfirst/go-whosonfirst-export)
* [whosonfirst/go-whosonfirst-features](https://github.com/whosonfirst/go-whosonfirst-feature)
* [whosonfirst/go-whosonfirst-findingaid](https://github.com/whosonfirst/go-whosonfirst-findingaid)
* [whosonfirst/go-whosonfirst-flags](https://github.com/whosonfirst/go-whosonfirst-flags)
* [whosonfirst/go-whosonfirst-format](https://github.com/whosonfirst/go-whosonfirst-format)
* [whosonfirst/go-whosonfirst-id](https://github.com/whosonfirst/go-whosonfirst-id)
* [whosonfirst/go-whosonfirst-image](https://github.com/whosonfirst/go-whosonfirst-image)
* [whosonfirst/go-whosonfirst-iterate](https://github.com/whosonfirst/go-whosonfirst-iterate)
* [whosonfirst/go-whosonfirst-iterwriter](https://github.com/whosonfirst/go-whosonfirst-iterwriter)
* [whosonfirst/go-whosonfirst-names](https://github.com/whosonfirst/go-whosonfirst-names)
* [whosonfirst/go-whosonfirst-placetypes](https://github.com/whosonfirst/go-whosonfirst-placetypes)
* [whosonfirst/go-whosonfirst-reader](https://github.com/whosonfirst/go-whosonfirst-reader)
* [whosonfirst/go-whosonfirst-sources](https://github.com/whosonfirst/go-whosonfirst-sources)
* [whosonfirst/go-whosonfirst-spatial](https://github.com/whosonfirst/go-whosonfirst-spatial)
* [whosonfirst/go-whosonfirst-spr](https://github.com/whosonfirst/go-whosonfirst-spr)
* [whosonfirst/go-whosonfirst-travel](https://github.com/whosonfirst/go-whosonfirst-travel)
* [whosonfirst/go-whosonfirst-uri](https://github.com/whosonfirst/go-whosonfirst-uri)
* [whosonfirst/go-whosonfirst-validate](https://github.com/whosonfirst/go-whosonfirst-validate)
* [whosonfirst/go-whosonfirst-writer](https://github.com/whosonfirst/go-whosonfirst-writer)
