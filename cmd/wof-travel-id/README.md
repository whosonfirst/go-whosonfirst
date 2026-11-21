## wof-travel-id

Follows selected record relationships from Who's On First IDs.

```
$> wof-travel-id -h
Follows selected record relationships from Who's On First IDs.
Usage:
	 wof-travel-id [options] path(N) path(N)
Valid options are:

  -csv
    	Emit results formatted as a comma-separated values.
  -hierarchies
    	Travel the hierarchies for each ID specified.
  -ids
    	Emit results as a line-separated list of IDs (traveled).
  -markdown
    	Emit results formatted as Markdown.
  -parent
    	Travel the parent ID of each ID specified.
  -singleton
    	... (default true)
  -source value
    	One or more valid whosonfirst/go-reader URIs to use to read WOF ID data, which may or may not be part of the sources to graph. This is work in progress.
  -stdin
    	Read IDs to travel from STDIN
  -superseded-by
    	Travel records that supersede each ID specified.
  -supersedes
    	Travel records that are superseded by each ID specified.
  -timings
    	Display timing information
```
