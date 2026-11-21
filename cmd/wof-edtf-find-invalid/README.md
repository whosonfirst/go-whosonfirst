## wof-edtf-find-invalid

Scans Who's On First records for invalid Extended Date/Time Format properties.

```
$> wof-edtf-find-invalid -h
Scans Who's On First records for invalid Extended Date/Time Format properties.
Usage:
	 wof-edtf-find-invalid [options] path(N) path(N)
Valid options are:

  -include-key
    	Include edtf: property of relevant Who's On First record in output. (default true)
  -include-path
    	Include path of relevant Who's On First record in output. (default true)
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst/v4/iterate.Iterator URI. Supported emitter URI schemes are:  (default "repo://")
```
