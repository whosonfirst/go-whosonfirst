## wof-opensearch-index

Iterates over Who's On First records and indexes them in OpenSearch.

```
$> wof-opensearch-index -h
Sends iterated records to a writer.
Usage:
	 wof-opensearch-index [options] path(N) path(N)
Valid options are:

  -forgiving
    	Be "forgiving" of failed writes, logging the issue(s) but not triggering errors
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/v4 URI. (default "repo://")
  -monitor-uri string
    	A valid sfomuseum/go-timings URI. (default "counter://PT60S")
  -verbose
    	Enable verbose (debug) logging
  -writer-uri value
    	One or more valid whosonfirst/go-writer/v4 URIs, each encoded as a gocloud.dev/runtimevar URI.
```
