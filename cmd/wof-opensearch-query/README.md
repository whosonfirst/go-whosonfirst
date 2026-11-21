## wof-opensearch-query

Submits a search query to OpenSearch and prints the response.

```
$> wof-opensearch-query -h
Submits a search query to OpenSearch and prints the response.
Usage:
	 wof-opensearch-query [options] path(N) path(N)
Valid options are:

  -opensearch-aws-credentials-uri string
    	...
  -opensearch-endpoint string
    	... (default "https://localhost:9200")
  -opensearch-index string
    	The name of the index to query.
  -opensearch-insecure
    	...
  -opensearch-password string
    	...
  -opensearch-query-all
    	Convenience flag to auto-generate a match_all query.
  -opensearch-query-id string
    	Convenience flag to auto-generate a ids/values ID query.
  -opensearch-username string
    	...
```
