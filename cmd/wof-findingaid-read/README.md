## wof-findingaid-read

Resolves Who's On First records through a finding aid and reads the corresponding documents.

```
$> wof-findingaid-read -h
Resolve one or more URIs, using a Who's On First finding aid and read their corresponding Who's On First documents, outputting each to STDOUT.
Usage:
	 wof-findingaid-read uri(N) uri(N)
  -data-template string
    	A valid URI template to use for resolving final reader URIs. Default is https://raw.githubusercontent.com/whosonfirst-data/{repo}/master/data/
  -reader-uri string
    	A valid whosonfirst/go-reader-findingaid URI
```
