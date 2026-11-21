## wof-iterate-emit

Emits records from one or more sources using a configured iterator.

```
$> wof-iterate-emit -h
Emit records in one or more whosonfirst/go-whosonfirst/v4/iterate.Iterator sources as structured data.
Usage:
	 wof-iterate-emit [options] uri(N) uri(N)
Valid options are:

  -geojson
    	Emit features as a well-formed GeoJSON FeatureCollection record.
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst/v4/iterate.Iterator URI. Supported iterator URI schemes are: cwd://,directory://,featurecollection://,file://,filelist://,geojsonl://,git://,null://,parquet://,repo:// (default "repo://")
  -json
    	Emit features as a well-formed JSON array.
  -null
    	Publish features to /dev/null
  -stdout
    	Publish features to STDOUT. (default true)
  -verbose
    	Enable verbose (debug) logging.
```
