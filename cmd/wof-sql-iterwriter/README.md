## wof-sql-iterwriter

Streams records from an iterator into a SQL writer.

```
$> wof-sql-iterwriter -h
Sends iterated records to a writer.
Usage:
	 wof-sql-iterwriter [options] path(N) path(N)
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
