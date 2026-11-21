# go-whosonfirst

Go package for common Who's On First functionality and interfaces.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/whosonfirst/go-whosonfirst.svg)](https://pkg.go.dev/github.com/whosonfirst/go-whosonfirst)

## Motivation

This package is work in progress and aims to merge a whole bunch of disparate `go-whosonfirst-*` packages in to one.

The original design for the Who's On First (WOF) Go packages was to have a lot of small "composable" packages but over time this has made it difficult to reason about all of the (common) functionality available for working with WOF records. It also makes it a chore to apply security updates.

Since these packages were first started Go has gotten better at excluding unused code in imports so while package itself have many dependencies only those used in any given program should be imported in your tools.

### Versioning

The "v4" version number reflects the largest version number plus one for the set of packages merged in to this one. Many of those packages will not have reached "v3" or even "v2".

Version "4" aims to introduce the least amount of change as possible to existing package names and method signatures. The goal for version 4 is to only require changes to path imports in existing code and nothing else. Once it has seen actual production use and "baked" for a little while then it might be time to consider a version "5" which could potentially introduce backwards-incompatible changes.

For example the `whosonfirst/go-whosonfirst/v4/iterate` and `whosonfirst/go-whosonfirst/v4/github` packages may be refactored into non-WOF-specific packages. This remains "tomorrow's problem" for the time being.

## Command line tools

Consult documentation in [cmd/](cmd/). The "kitchen-sink" Who's On First command line application is [cmd/wof](cmd/wof). The easiest way to build tools is with the handy `cli` Makefile command. For example:

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/wof cmd/wof/main.go
go build -mod vendor -ldflags="-s -w" -o bin/wof-concordances-keys cmd/wof-concordances-keys/main.go
go build -mod vendor -ldflags="-s -w" -tags null -o bin/wof-sql-create cmd/wof-sql-create/main.go
go build -mod vendor -ldflags="-s -w" -tags null -o bin/wof-sql-index cmd/wof-sql-index/main.go
... and so on
```

## Packages

| Directory | Description |
| --- | --- |
| [`app`](app/) | Application code for command line tools in `cmd`. |
| [`concordances`](concordances/) | Lists distinct keys in `wof:concordances`, which links Who's On First records to identifiers from other datasets. |
| [`database`](database/) | `database/sql` opens SQL databases; its `tables`, `indexer`, `prune`, `spr`, and `writer` subpackages define tables, index and remove records, query standard places results, and write records. Under `database/opensearch`, `client`, `document`, `indexer`, `schema/v2`, and `writer` configure clients, transform and index records, define mappings, and write documents. |
| [`derivatives`](derivatives/) | Registers record providers for generating derivative representations. `derivatives/http` parses requests and `derivatives/http/api` serves responses; `derivatives/image` and `derivatives/svg` render raster images and SVG. |
| [`edtf`](edtf/) | Updates legacy Extended Date/Time Format (EDTF) values in Who's On First records. |
| [`errors`](errors/) | The package at this path is named `derivatives`; it defines `ErrNotImplemented` and `ErrNotFound`. |
| [`export`](export/) | Prepares records for storage by applying defaults, validating properties, and formatting output. `export/properties` prepares individual properties, `export/app/export` provides a CLI runner, and `export/wasm` exposes record preparation to JavaScript. |
| [`feature`](feature/) | Defines feature-property errors. `feature/alt` handles alternate geometry records, `feature/constants` defines reserved IDs, and `feature/geometry` and `feature/properties` interpret record geometries and properties. |
| [`fetch`](fetch/) | Retrieves Who's On First records by ID and can recursively fetch their ancestors. |
| [`findingaid`](findingaid/) | Represents the repositories in a finding aid, which maps record IDs to their repositories. `findingaid/provider` supplies iterator sources, `findingaid/producer` builds the mappings (including SQL and document-store writers), `findingaid/resolver` looks them up, and `findingaid/reader` uses them to read records. |
| [`flags`](flags/) | Defines filter interfaces for Who's On First queries. The `custom`, `date`, `existential`, `geometry`, and `placetypes` subpackages implement filters for custom values, dates, record status, geometries, and place types. |
| [`format`](format/) | Formats Who's On First GeoJSON documents for storage. `format/wasm` exposes the formatter to JavaScript. |
| [`geo`](geo/) | Finds a point inside an `orb` geometry. |
| [`geojson`](geojson/) | Configures the JSON encoder and decoder used by `orb`'s GeoJSON package. |
| [`github`](github/) | Groups GitHub API helpers: `github/client` configures clients, `github/organizations` lists and creates organization repositories, `github/repositories` lists changed files in commits, and `github/util` provides a legacy client helper. |
| [`hierarchies`](hierarchies/) | Works with Who's On First record hierarchies, including labeling and merging them. |
| [`http`](http/) | Extracts Who's On First IDs from HTTP requests. `http/webfinger` resolves IDs through WebFinger. |
| [`id`](id/) | Defines interfaces and implementations for generating unique Who's On First IDs. |
| [`iterate`](iterate/) | Defines record iterators and implementations for local sources. The `bucket`, `git` (including `git/github`), `parquet`, `reader`, and `sql` subpackages add blob storage, Git and GitHub, Parquet, reader-backed, and SQL sources. `iterate/filters` selects records, `iterate/flags` configures iterators from the CLI, `iterate/writer` forwards records to writers, and `iterate/fixtures` supports tests. |
| [`names`](names/) | Processes multilingual record names. `names/tags` parses Who's On First language tags, and `names/utils` converts between those tags and RFC 5646. |
| [`parquet`](parquet/) | Converts Who's On First GeoJSON features to Parquet rows, writes them, and finds records by ID in Parquet files. |
| [`placetypes`](placetypes/) | Defines place types and their hierarchical relationships. `placetypes/draw` graphs the relationships with Graphviz, `placetypes/filter` selects place types, and `placetypes/wasm` exposes lookup, validation, and relationship queries to JavaScript. |
| [`properties`](properties/) | Reads machine-readable Who's On First property definitions. `properties/index` catalogs and indexes directories of those definitions. |
| [`reader`](reader/) | Reads records by ID through a `go-reader` backend as bytes, streams, or GeoJSON features. |
| [`sources`](sources/) | Looks up metadata about the data sources cited in Who's On First records. |
| [`spatial`](spatial/) | Defines spatial indexing and filtering interfaces. `spatial/application` configures indexes, `spatial/database` registers backends, and `spatial/filter`, `spatial/query`, `spatial/geo`, and `spatial/maptile` support spatial lookups. `spatial/hierarchy` and `spatial/hierarchy/filter` resolve hierarchies from query results. `spatial/pmtiles` uses PMTiles with a cache subpackage; `spatial/sqlite` uses SQLite and includes a Well-Known Text parser. Under `spatial/grpc` are request helpers, server code, and generated bindings; under `spatial/www` are HTTP middleware, API handlers, and web assets. `spatial/fixtures` and `spatial/fixtures/microhoods` support tests. |
| [`spr`](spr/) | Defines standard places results (SPRs), compact summaries of Who's On First records. `spr/geojson` converts results to GeoJSON, `spr/sort` orders them, `spr/sqlite` queries them from SQLite, and `spr/util` converts a result to a string-valued map. |
| [`travel`](travel/) | Follows record links such as parents, hierarchy entries, and supersession links. `travel/traveler` provides handlers for individual relationship types. |
| [`uri`](uri/) | Parses and constructs record paths and URIs for primary records and alternate geometries. |
| [`validate`](validate/) | Checks record properties and GeoJSON syntax. `validate/wasm` exposes validation to JavaScript. |
| [`wof`](wof/) | Registers and constructs subcommands for `cmd/wof-cli` tool. |
| [`writer`](writer/) | Writes GeoJSON features or bytes through a `go-writer` backend, optionally applying a custom exporter. |

## Tests

Tests are also incomplete at this time. Because of build tags (see below) the easiest way to run test it to use the handy `tests` Makefile tagret which will assign any necessary default tags. For example:

```
$> make test
... time passes, tests happen
```

## Build tags

Build tags are still in flux.

| Tag | Description |
| --- | ---  |
| `sqlite` | Enable the use of one or more `database/sql` SQLite-backed drivers. |
| `sqlite3` | Enable the use of the `mattn/go-sqlite` driver. |
| `modernc` | Enable the use of the `modernc.org/sqlite` driver. |
| `pmtiles` | Enable the use of a `protomaps/go-pmtiles`-backed driver. |

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

### [whosonfirst/go-whosonfirst-iterate-bucket](https://github.com/whosonfirst/go-whosonfirst-iterate-bucket)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/iterate/bucket`.

### [whosonfirst/go-whosonfirst-iterate-git](https://github.com/whosonfirst/go-whosonfirst-iterate-git)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/iterate/git`.

### [whosonfirst/go-whosonfirst-iterate-reader](https://github.com/whosonfirst/go-whosonfirst-iterate-reader)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/iterate/reader`.

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

### [whosonfirst/go-whosonfirst-spatial-grpc](https://github.com/whosonfirst/go-whosonfirst-spatial-grpc)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/spatial/grpc`.

### [whosonfirst/go-whosonfirst-spatial-pmtiles](https://github.com/whosonfirst/go-whosonfirst-spatial-pmtiles)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/spatial/pmtiles`.

### [whosonfirst/go-whosonfirst-spatial-sqlite](https://github.com/whosonfirst/go-whosonfirst-spatial-sqlite)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/spatial/sqlite`.

### [whosonfirst/go-whosonfirst-spatial-www](https://github.com/whosonfirst/go-whosonfirst-spatial-www)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/spatial/www`.

### [whosonfirst/go-whosonfirst-spr](https://github.com/whosonfirst/go-whosonfirst-spr)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/spr`.

### [whosonfirst/go-whosonfirst-spr-geojson](https://github.com/whosonfirst/go-whosonfirst-spr-geojson)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/spr/geojson`.

### [whosonfirst/go-whosonfirst-sqlite-spr](https://github.com/whosonfirst/go-whosonfirst-sqlite-spr)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/spr/sqlite`.

### [whosonfirst/go-whosonfirst-travel](https://github.com/whosonfirst/go-whosonfirst-travel)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/travel`.

### [whosonfirst/go-whosonfirst-uri](https://github.com/whosonfirst/go-whosonfirst-uri)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/uri`.

### [whosonfirst/go-whosonfirst-validate](https://github.com/whosonfirst/go-whosonfirst-validate)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/validate`.

### [whosonfirst/go-whosonfirst-writer](https://github.com/whosonfirst/go-whosonfirst-writer)

Replace with `github.com/whosonfirst/go-whosonfirst/v4/writer`.
