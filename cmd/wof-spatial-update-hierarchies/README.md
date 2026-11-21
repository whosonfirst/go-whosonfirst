## wof-spatial-update-hierarchies

Rebuilds record hierarchies using spatial query results.

```
$> wof-spatial-update-hierarchies -h
Perform point-in-polygon (PIP), and related update, operations on a set of Who's on First records.
Usage:
	 wof-spatial-update-hierarchies [options] uri(N) uri(N)
Valid options are:

  -exporter-uri string
    	A valid whosonfirst/go-whosonfirst-export URI. (default "whosonfirst://")
  -is-ceased value
    	One or more existential flags (-1, 0, 1) to filter PIP results.
  -is-current value
    	One or more existential flags (-1, 0, 1) to filter PIP results.
  -is-deprecated value
    	One or more existential flags (-1, 0, 1) to filter PIP results.
  -is-superseded value
    	One or more existential flags (-1, 0, 1) to filter PIP results.
  -is-superseding value
    	One or more existential flags (-1, 0, 1) to filter PIP results.
  -mapshaper-server string
    	A valid HTTP URI pointing to a sfomuseum/go-sfomuseum-mapshaper server endpoint.
  -source-iterator-uri value
    	Zero or more URIs denoting data sources to use for indexing the spatial database at startup. URIs take the form of {ITERATOR_URI} + "#" + {PIPE-SEPARATED LIST OF ITERATOR SOURCES}. Where {ITERATOR_URI} is expected to be a registered whosonfirst/go-whosonfirst-iterate/v2 iterator (emitter) URI and {ITERATOR SOURCES} are valid input paths for that iterator. Supported whosonfirst/go-whosonfirst-iterate/v2 iterator schemes are: cwd://, directory://, featurecollection://, file://, filelist://, geojsonl://, null://, repo://.
  -spatial-database-uri string
    	A valid whosonfirst/go-whosonfirst-spatial URI. This is the database of spatial records that will for PIP-ing. (default "rtree://")
  -target-iterator-uri value
    	Zero or more URIs denoting target data sources whose hierarchies need updating. URIs take the form of {ITERATOR_URI} + "#" + {PIPE-SEPARATED LIST OF ITERATOR SOURCES}. Where {ITERATOR_URI} is expected to be a registered whosonfirst/go-whosonfirst-iterate/v2 iterator (emitter) URI and {ITERATOR SOURCES} are valid input paths for that iterator. Supported whosonfirst/go-whosonfirst-iterate/v2 iterator schemes are: cwd://, directory://, featurecollection://, file://, filelist://, geojsonl://, null://, repo://.
  -verbose
    	Enable (verbose) debug logging.
  -writer-uri string
    	A valid whosonfirst/go-writer URI. This is where updated records will be written to. (default "null://")
```
