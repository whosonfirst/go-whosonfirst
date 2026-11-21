## wof-fetch-records

Fetches Who's On First records by ID and can also retrieve their ancestors.

```
$> wof-fetch-records -h
Fetch one or more Who's on First records and, optionally, their ancestors.

Usage:
  wof-fetch-records [options] [path1 path2 ... pathN]

Options:
  -belongs-to value
    	One or more placetypes that a given ID may belong to to also fetch. You may also pass 'all' as a short-hand to fetch the entire hierarchy for a place.
  -max-clients int
    	The maximum number of concurrent requests for multiple Who's On First records. (default 10)
  -reader-uri string
    	A valid whosonfirst/go-reader URI. (default "findingaid://https/data.whosonfirst.org/findingaid?template=https://raw.githubusercontent.com/whosonfirst-data/{repo}/master/data/")
  -retries int
    	The maximum number of attempts to try fetching a record. (default 3)
  -verbose
    	Enable verbose (debug) logging.
  -writer-uri string
    	A valid whosonfirst/go-writer URI. (default "stdout://")

Notes:

pathN may be any valid Who's On First ID or URI that can be parsed by the
go-whosonfirst-uri package.
```
