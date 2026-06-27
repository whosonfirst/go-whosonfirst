# go-whosonfirst-spatial-grpc

gRPC support for the `go-whosonfirst-spatial` interfaces.

## Documentation

Documentation is incomplete.

## Tools

```
$> make cli
go build -mod vendor -o bin/server cmd/server/main.go
go build -mod vendor -o bin/client cmd/client/main.go
```

### client

```
$> ./bin/client -h
  -alternate-geometry value
    	One or more alternate geometry labels (wof:alt_label) values to filter results by.
  -cessation string
    	A valid EDTF date string.
  -geometries string
    	Valid options are: all, alt, default. (default "all")
  -host string
    	The host of the gRPC server to connect to. (default "localhost")
  -inception string
    	A valid EDTF date string.
  -is-ceased value
    	One or more existential flags (-1, 0, 1) to filter results by.
  -is-current value
    	One or more existential flags (-1, 0, 1) to filter results by.
  -is-deprecated value
    	One or more existential flags (-1, 0, 1) to filter results by.
  -is-superseded value
    	One or more existential flags (-1, 0, 1) to filter results by.
  -is-superseding value
    	One or more existential flags (-1, 0, 1) to filter results by.
  -latitude float
    	A valid latitude.
  -longitude float
    	A valid longitude.
  -null
    	Emit results to /dev/null
  -placetype value
    	One or more place types to filter results by.
  -port int
    	The port of the gRPC server to connect to. (default 8082)
  -property value
    	One or more Who's On First properties to append to each result.
  -sort-uri value
    	Zero or more whosonfirst/go-whosonfirst-spr/sort URIs.
  -stdout
    	Emit results to STDOUT (default true)
```

For example:

```
$> ./bin/client -latitude 43.873889 -longitude 18.408611 -inception 199X | jq
{
  "places": [
    {
      "id": "85632609",
      "parent_id": "102191581",
      "placetype": "country",
      "country": "BA",
      "repo": "whosonfirst-data-admin-ba",
      "path": "856/326/09/85632609.geojson",
      "uri": "https://data.whosonfirst.org/856/326/09/85632609.geojson",
      "latitude": 44.091038,
      "longitude": 18.06843,
      "is_current": 1,
      "is_superseding": 1,
      "inception_date": "1992-~03/1992-~04",
      "cessation_date": "..",
      "supersedes": [
        1108955785
      ],
      "belongs_to": [
        102191581
      ],
      "last_modified": 1616353750,
      "name": "Bosnia and Herzegovina"
    }
  ]
}
```

### server

```
$> ./bin/server -h
  -custom-placetypes string
    	A JSON-encoded string containing custom placetypes defined using the syntax described in the whosonfirst/go-whosonfirst-placetypes repository.
  -enable-custom-placetypes
    	Enable wof:placetype values that are not explicitly defined in the whosonfirst/go-whosonfirst-placetypes repository.
  -host string
    	The host to listen for requests on (default "localhost")
  -is-wof
    	Input data is WOF-flavoured GeoJSON. (Pass a value of '0' or 'false' if you need to index non-WOF documents. (default true)
  -iterator-uri value
    	Zero or more URIs denoting data sources to use for indexing the spatial database at startup. URIs take the form of {ITERATOR_URI} + "#" + {PIPE-SEPARATED LIST OF ITERATOR SOURCES}. Where {ITERATOR_URI} is expected to be a registered whosonfirst/go-whosonfirst-iterate/v2 iterator (emitter) URI and {ITERATOR SOURCES} are valid input paths for that iterator. Supported whosonfirst/go-whosonfirst-iterate/v2 iterator schemes are: cwd://, directory://, featurecollection://, file://, filelist://, geojsonl://, null://, repo://.
  -port int
    	The port to listen for requests on (default 8082)
  -properties-reader-uri string
    	A valid whosonfirst/go-reader.Reader URI. Available options are: [fs:// null:// repo:// stdin://]. If the value is {spatial-database-uri} then the value of the '-spatial-database-uri' implements the reader.Reader interface and will be used.
  -spatial-database-uri string
    	A valid whosonfirst/go-whosonfirst-spatial/data.SpatialDatabase URI. options are: [rtree://] (default "rtree://")
```

For example:

```
$> ./bin/server -spatial-database-uri 'rtree://#/usr/local/data/whosonfirst-data-admin-ba/'
2021/03/26 08:59:36 Listening on localhost:8082
08:59:37.890704 [server] STATUS indexing 8792 records indexed
08:59:38.890291 [server] STATUS indexing 16249 records indexed
2021/03/26 08:59:39 time to index paths (1) 2.87549934s
08:59:39.765730 [server] STATUS finished indexing in 2.875533622s
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-spatial
* https://github.com/whosonfirst/go-whosonfirst-spatial-sqlite
* https://github.com/whosonfirst/go-whosonfirst-spatial-pmtiles
* https://github.com/whosonfirst/go-whosonfirst-spr
* https://github.com/grpc/grpc-go