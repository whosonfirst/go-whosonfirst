## wof-validate

Checks records from an iterator for valid GeoJSON and Who's On First properties.

```
$> wof-validate -h
Validate the contents of one or more whosonfirst/go-whosonfirst-iterate/v3 data sources.
Usage:
	 wof-validate path(N) path(N)
Valid arguments are:
  -all
    	Enable all validation checks.
  -edtf
    	Validate edtf: properties. (default true)
  -id
    	Validate wof:id property. (default true)
  -is-current
    	Validate mz:is_current property.
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst/v4/iterate.Iterator URI (default "repo://")
  -name
    	Validate wof:name property. (default true)
  -names
    	Validate WOF/RFC 5646 names.
  -placetype
    	Validate wof:placetype property. (default true)
  -repo
    	Validate wof:repo property. (default true)
  -verbose
    	Be chatty about what's happening.
```
