# go-whosonfirst

Go package for common Who's On First functionality and interfaces.

## Motivation

This package is work in progress and aims to merge a whole bunch of disparate `go-whosonfirst-*` packages in to one.

The original design for the Who's On First (WOF) Go packages was to have a lot of small "composable" packages but over time this has made it difficult to reason about all of the (common) functionality available for working with WOF records. It also makes it a chore to apply updates.

Since these packages were first started Go has gotten better at excluding unused code in imports so while package itself have many dependencies only those used in any given program should be imported in your tools.

### Versioning

The "v4" version number reflects the largest version number plus one for the set of packages merged in to this one. Many of those packages will not have reached "v3" or even "v2".

Version "4" aims to introduce the least amount of change as possible to existing package names and method signatures. The goal for version 4 is to only require changes to path imports in existing code and nothing else. Once it has seen actual production use and "baked" for a little while then it might be time to consider a version "5" which could potentially introduce backwards-incompatible changes.

For example the `whosonfirst/go-whosonfirst/v4/iterate` and `whosonfirst/go-whosonfirst/v4/github` packages may be refactored into non-WOF-specific packages. This remains "tomorrow's problem" for the time being.

## Documentation

Documentation is incomplete at this time.

Documentation for individual packages which have been merged in to this one have also been copied in to their respective folders. There is outstanding work to unify all that documentation into a coherent whole.

## Tests

Tests are also incomplete at this time.

## Packages that have merged/superseded in to this package

### [whosonfirst/go-reader-findingaid](https://github.com/whosonfirst/go-reader-findingaid)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/findingaid/reader`.

### [whosonfirst/go-whosonfirst-concordances](https://github.com/whosonfirst/go-whosonfirst-concordances)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/concordances`.

### [whosonfirst/go-whosonfirst-database](https://github.com/whosonfirst/go-whosonfirst-database)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/database`.

### [whosonfirst/go-whosonfirst-edtf](https://github.com/whosonfirst/go-whosonfirst-edtf)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/edtf`.

### [whosonfirst/go-whosonfirst-export](https://github.com/whosonfirst/go-whosonfirst-export)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/export`.

### [whosonfirst/go-whosonfirst-feature](https://github.com/whosonfirst/go-whosonfirst-feature)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/feature`.

### [whosonfirst/go-whosonfirst-findingaid](https://github.com/whosonfirst/go-whosonfirst-findingaid)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/findingaid`.

### [whosonfirst/go-whosonfirst-flags](https://github.com/whosonfirst/go-whosonfirst-flags)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/flags`.

### [whosonfirst/go-whosonfirst-format](https://github.com/whosonfirst/go-whosonfirst-format)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/format`.

### [whosonfirst/go-whosonfirst-id](https://github.com/whosonfirst/go-whosonfirst-id)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/id`.

### [whosonfirst/go-whosonfirst-image](https://github.com/whosonfirst/go-whosonfirst-image)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/derivatives/image`.

### [whosonfirst/go-whosonfirst-iterate](https://github.com/whosonfirst/go-whosonfirst-iterate)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/iterate`.

### [whosonfirst/go-whosonfirst-iterate](https://github.com/whosonfirst/go-whosonfirst-iterate-git)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/iterate/git`.

### [whosonfirst/go-whosonfirst-iterwriter](https://github.com/whosonfirst/go-whosonfirst-iterwriter)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/iterate/writer`.

### [whosonfirst/go-whosonfirst-names](https://github.com/whosonfirst/go-whosonfirst-names)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/names`.

### [whosonfirst/go-whosonfirst-placetypes](https://github.com/whosonfirst/go-whosonfirst-placetypes)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/placetypes`.

### [whosonfirst/go-whosonfirst-properties](https://github.com/whosonfirst/go-whosonfirst-properties)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/properties`.

### [whosonfirst/go-whosonfirst-reader](https://github.com/whosonfirst/go-whosonfirst-reader)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/reader`.

### [whosonfirst/go-whosonfirst-sources](https://github.com/whosonfirst/go-whosonfirst-sources)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/sources`.

### [whosonfirst/go-whosonfirst-spatial](https://github.com/whosonfirst/go-whosonfirst-spatial)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/spatial`.

### [whosonfirst/go-whosonfirst-spr](https://github.com/whosonfirst/go-whosonfirst-spr)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/spr`.

### [whosonfirst/go-whosonfirst-spr](https://github.com/whosonfirst/go-whosonfirst-spr-geojson)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/spr/geojson`.

### [whosonfirst/go-whosonfirst-travel](https://github.com/whosonfirst/go-whosonfirst-travel)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/travel`.

### [whosonfirst/go-whosonfirst-uri](https://github.com/whosonfirst/go-whosonfirst-uri)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/uri`.

### [whosonfirst/go-whosonfirst-validate](https://github.com/whosonfirst/go-whosonfirst-validate)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/validate`.

### [whosonfirst/go-whosonfirst-writer](https://github.com/whosonfirst/go-whosonfirst-writer)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/writer`.
