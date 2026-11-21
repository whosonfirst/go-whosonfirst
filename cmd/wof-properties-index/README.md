## wof-properties-index

Scans Who's On First records and creates missing property definition files.

```
$> wof-properties-index -h
Scans Who's On First records and creates missing property definition files.
Usage:
	 wof-properties-index [options] path(N) path(N)
Valid options are:

  -alternate value
    	One or more paths to alternate properties directories that will be crawled to check for existing properties (that will not be duplicated).
  -debug
    	Go through all the motions but don't write any new files.
  -exclude value
    	One or more valid regular expressions to use for excluding property names you don't want to index
  -iterator-uri string
    	A valid go-whosonfirst-iterate/v2 URI. (default "repo://")
  -properties string
    	The path to your whosonfirst-properties/properties directory
```
