## wof-spr-as-geojson

Reads features referenced by a standard places response and outputs a GeoJSON FeatureCollection.

```
$> wof-spr-as-geojson -h
Reads features referenced by a standard places response and outputs a GeoJSON FeatureCollection.
Usage:
	 wof-spr-as-geojson [options]
Valid options are:

  -as-wof
    	Parse each item returned by -path as though it were a Who's On First style ID/path.
  -path string
    	A valid tidwall/gjson query path for finding the path for each element in your SPR response. (default "places.#.wof:path")
  -reader-uri string
    	A valid whosonfirst/go-reader URI.
```
