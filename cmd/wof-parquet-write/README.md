## wof-parquet-write

Converts records from a configured iterator to Parquet rows and writes them to a configured destination.

```
$> wof-parquet-write -h
Converts records from a configured iterator to Parquet rows and writes them to a configured destination.
Usage:
	 wof-parquet-write [options] path(N) path(N)
Valid options are:

  -iterator-uri string
    	A registered whosonfirst/go-whosonfirst/v4/iterate.Iterator URI. (default "repo://")
  -writer-uri string
    	The path to a file where Parquet data should be written. If "-" then data will be written to STDOUT.
```
