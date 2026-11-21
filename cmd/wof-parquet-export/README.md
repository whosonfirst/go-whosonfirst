## wof-parquet-export

Exports records from Who's On First repositories to Parquet files.

```
$> wof-parquet-export -h
Write data from one or more Who's On First style repositories to Parquet files.
Usage:
	wof-parquet-export [options]
Valid options are:
  -git-exclude value
    	If -iterator-uri is git://, zero or more repository prefixe to exclude by (if -repo flag is empty). Default is [whosonfirst-data-admin-alt].
  -git-org string
    	If -iterator-uri is git://, the organization to use when cloning repos. (default "whosonfirst-data")
  -git-prefix value
    	If -iterator-uri is git://, zero or more repository prefixe to filter by (if -repo flag is empty). Default is [whosonfirst-data-admin-].
  -git-repo value
    	If -iterator-uri is git://, zero or more repos to clone and create Parquet exports for. If empty then repos will be limited to those starting in whosonfirst-data-admin.
  -iterator-uri string
    	One of two registered whosonfirst/go-whosonfirst/v4/iterate.Iterator URIs: repo:// or git:// (default "git:///tmp")
  -verbose
    	Enable verbose (debug) logging.
  -writer-root string
    	The path on the local filesystem where Parquet files should be written to. (default "/usr/local/data/whosonfirst-parquet")
```
