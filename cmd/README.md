# cmd

Command-line tools and WebAssembly entry points for working with Who's On First records.

| Command | Description |
| --- | --- |
| [`wof-cli`](wof-cli/README.md) | Kitchen-sink command line application for working with Who's On First documents. |
| [`wof-concordances-keys`](wof-concordances-keys/README.md) | Lists the distinct external identifier namespaces used in Who's On First concordances. |
| [`wof-derivatives-server`](wof-derivatives-server/README.md) | Serves derivative representations of Who's On First records over HTTP. |
| [`wof-edit-wasm`](wof-edit-wasm/README.md) | Registers JavaScript functions for preparing, formatting, and validating records and looking up placetypes in the wof editor. |
| [`wof-edtf-find-invalid`](wof-edtf-find-invalid/README.md) | Scans Who's On First records for invalid Extended Date/Time Format properties. |
| [`wof-edtf-update-unknown-uncertain`](wof-edtf-update-unknown-uncertain/README.md) | Updates legacy open and unknown Extended Date/Time Format values and writes changed records to a configured writer. |
| [`wof-export`](wof-export/README.md) | Command-line exporter for Who's On First records. |
| [`wof-fetch-records`](wof-fetch-records/README.md) | Fetches Who's On First records by ID and can also retrieve their ancestors. |
| [`wof-findingaid-create-dynamodb-import`](wof-findingaid-create-dynamodb-import/README.md) | Converts finding aid archives into data for DynamoDB import. |
| [`wof-findingaid-create-dynamodb-tables`](wof-findingaid-create-dynamodb-tables/README.md) | Creates the DynamoDB tables used by the Who's On First finding aid. |
| [`wof-findingaid-csv2docstore`](wof-findingaid-csv2docstore/README.md) | Loads archived finding aid CSV data into a Go Cloud document store. |
| [`wof-findingaid-csv2sql`](wof-findingaid-csv2sql/README.md) | Loads archived finding aid CSV data into a SQL database. |
| [`wof-findingaid-populate`](wof-findingaid-populate/README.md) | Scans record sources and sends ID-to-repository mappings to a finding aid producer. |
| [`wof-findingaid-read`](wof-findingaid-read/README.md) | Resolves Who's On First records through a finding aid and reads the corresponding documents. |
| [`wof-findingaid-resolve`](wof-findingaid-resolve/README.md) | Looks up the repository containing each requested Who's On First ID. |
| [`wof-findingaid-resolverd`](wof-findingaid-resolverd/README.md) | Exposes a Who's On First finding aid resolver over HTTP. |
| [`wof-findingaid-sources`](wof-findingaid-sources/README.md) | Lists iterator sources supplied by a finding aid provider. |
| [`wof-format`](wof-format/README.md) | Formats a Who's On First GeoJSON feature or checks its formatting. |
| [`wof-format-wasm`](wof-format-wasm/README.md) | Registers the Who's On First GeoJSON formatter as a JavaScript function. |
| [`wof-inner-point`](wof-inner-point/README.md) | Computes a point within a GeoJSON geometry and prints it as a GeoJSON feature. |
| [`wof-iterate-count`](wof-iterate-count/README.md) | Counts records from one or more sources using a configured iterator. |
| [`wof-iterate-emit`](wof-iterate-emit/README.md) | Emits records from one or more sources using a configured iterator. |
| [`wof-iterwriter`](wof-iterwriter/README.md) | Sends records from a configured iterator to a configured writer. |
| [`wof-names-parse`](wof-names-parse/README.md) | Parses Who's On First name language tags. |
| [`wof-opensearch-create-index`](wof-opensearch-create-index/README.md) | Creates an OpenSearch index. |
| [`wof-opensearch-delete-index`](wof-opensearch-delete-index/README.md) | Deletes an OpenSearch index. |
| [`wof-opensearch-get-mapping`](wof-opensearch-get-mapping/README.md) | Retrieves the field mapping for an OpenSearch index. |
| [`wof-opensearch-index`](wof-opensearch-index/README.md) | Iterates over Who's On First records and indexes them in OpenSearch. |
| [`wof-opensearch-indices-stats`](wof-opensearch-indices-stats/README.md) | Retrieves OpenSearch index statistics. |
| [`wof-opensearch-list-indices`](wof-opensearch-list-indices/README.md) | Lists OpenSearch indexes. |
| [`wof-opensearch-put-mapping`](wof-opensearch-put-mapping/README.md) | Updates the field mapping for an OpenSearch index. |
| [`wof-opensearch-put-settings`](wof-opensearch-put-settings/README.md) | Updates the settings for an OpenSearch index. |
| [`wof-opensearch-query`](wof-opensearch-query/README.md) | Submits a search query to OpenSearch and prints the response. |
| [`wof-opensearch-sign-request`](wof-opensearch-sign-request/README.md) | Signs an OpenSearch HTTP request using AWS credentials. |
| [`wof-parquet-export`](wof-parquet-export/README.md) | Exports records from Who's On First repositories to Parquet files. |
| [`wof-parquet-find`](wof-parquet-find/README.md) | Finds a Who's On First record by ID in a Parquet file. |
| [`wof-parquet-write`](wof-parquet-write/README.md) | Converts records from a configured iterator to Parquet rows and writes them to a configured destination. |
| [`wof-placetypes`](wof-placetypes/README.md) | Lists known Who's On First placetypes. |
| [`wof-placetypes-ancestors`](wof-placetypes-ancestors/README.md) | Lists ancestor placetypes, with optional role filtering. |
| [`wof-placetypes-children`](wof-placetypes-children/README.md) | Lists immediate child placetypes, with optional role filtering. |
| [`wof-placetypes-compile-spec`](wof-placetypes-compile-spec/README.md) | Compiles individual placetype definitions into a single specification. |
| [`wof-placetypes-descendants`](wof-placetypes-descendants/README.md) | Lists descendant placetypes, with optional role filtering. |
| [`wof-placetypes-is-valid`](wof-placetypes-is-valid/README.md) | Checks whether a name identifies a known Who's On First placetype. |
| [`wof-placetypes-render-spec`](wof-placetypes-render-spec/README.md) | Renders the placetype specification as a graph. |
| [`wof-placetypes-wasm`](wof-placetypes-wasm/README.md) | Registers JavaScript functions for listing and validating placetypes and following their relationships. |
| [`wof-properties-index`](wof-properties-index/README.md) | Scans Who's On First records and creates missing property definition files. |
| [`wof-properties-report`](wof-properties-report/README.md) | Generates a CSV report of Who's On First property definitions. |
| [`wof-spatial-grpc-client`](wof-spatial-grpc-client/README.md) | Queries a Who's On First spatial gRPC service from the command line. |
| [`wof-spatial-grpc-server`](wof-spatial-grpc-server/README.md) | Serves Who's On First spatial queries over gRPC. |
| [`wof-spatial-http-server`](wof-spatial-http-server/README.md) | Serves the Who's On First spatial HTTP API and web interface. |
| [`wof-spatial-inner-point`](wof-spatial-inner-point/README.md) | Computes a point within a GeoJSON geometry and prints it as a GeoJSON feature. |
| [`wof-spatial-intersects`](wof-spatial-intersects/README.md) | Queries a spatial index for records intersecting a geometry. |
| [`wof-spatial-pip`](wof-spatial-pip/README.md) | Queries a spatial index for records containing a point. |
| [`wof-spatial-update-hierarchies`](wof-spatial-update-hierarchies/README.md) | Rebuilds record hierarchies using spatial query results. |
| [`wof-spr-as-geojson`](wof-spr-as-geojson/README.md) | Reads features referenced by a standard places response and outputs a GeoJSON FeatureCollection. |
| [`wof-sql-create`](wof-sql-create/README.md) | Creates SQL tables for Who's On First data. |
| [`wof-sql-index`](wof-sql-index/README.md) | Indexes Who's On First records in SQL tables. |
| [`wof-sql-iterwriter`](wof-sql-iterwriter/README.md) | Streams records from an iterator into a SQL writer. |
| [`wof-sql-prune`](wof-sql-prune/README.md) | Removes selected Who's On First records from SQL tables. |
| [`wof-travel-belongsto`](wof-travel-belongsto/README.md) | Finds records with matching wof:belongsto values and reports the results. |
| [`wof-travel-id`](wof-travel-id/README.md) | Follows selected record relationships from Who's On First IDs. |
| [`wof-validate`](wof-validate/README.md) | Checks records from an iterator for valid GeoJSON and Who's On First properties. |
| [`wof-validate-wasm`](wof-validate-wasm/README.md) | Registers the Who's On First record validator as a JavaScript function. |
