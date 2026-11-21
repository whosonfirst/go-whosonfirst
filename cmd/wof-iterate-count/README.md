## wof-iterate-count

Counts records from one or more sources using a configured iterator.

```
$> wof-iterate-count -h
Count files in one or more whosonfirst/go-whosonfirst/v4/iterate.Iterator sources.
Usage:
	 wof-iterate-count [options] uri(N) uri(N)
Valid options are:

  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst/v4/iterate.Iterator URI. Supported iterator URI schemes are: cwd://,directory://,featurecollection://,file://,filelist://,geojsonl://,git://,null://,parquet://,repo:// (default "repo://")
  -verbose
    	Enable verbose (debug) logging.
```
