## wof-findingaid-populate

Scans record sources and sends ID-to-repository mappings to a finding aid producer.

```
$> wof-findingaid-populate -h
Scans record sources and sends ID-to-repository mappings to a finding aid producer.
Usage:
	 wof-findingaid-populate [options] path(N) path(N)
Valid options are:

  -atomic
    	Produce atomic findingaids for each item in a source list. If true then -producer URI must be a valid URI template containing a '{source}' variable to expand with findingaid name.
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/v3.Iterator URI. (default "repo://")
  -producer-uri string
    	A valid whosonfirst/go-whosonfirst-findingaid/v2/producer URI. (default "csv://?archive=archive.tar.gz")
  -provider-uri string
    	An optional whosonfirst/go-whosonfirst-findingaid/v2/provider URI to use for deriving additional sources.
  -verbose
    	Enable verbose (debug) logging.
```
