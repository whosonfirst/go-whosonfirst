## wof-sql-index

Indexes Who's On First records in SQL tables.

```
$> wof-sql-index -h
Indexes Who's On First records in SQL tables.
Usage:
	 wof-sql-index [options] path(N) path(N)
Valid options are:

  -all
    	Index all tables (except the 'search' and 'geometries' tables which you need to specify explicitly)
  -ancestors
    	Index the 'ancestors' tables
  -concordances
    	Index the 'concordances' tables
  -database-uri string
    	A URI in the form of 'sql://{DATABASE_SQL_ENGINE}?dsn={DATABASE_SQL_DSN}'. For example: sql://sqlite3?dsn=test.db
  -geojson
    	Index the 'geojson' table
  -geometries
    	Index the 'geometries' table (requires that libspatialite already be installed if using SQLite)
  -index-alt value
    	Zero or more table names where alt geometry files should be indexed.
  -index-relations
    	Index the records related to a feature, specifically wof:belongsto, wof:depicts and wof:involves. Alt files for relations are not indexed at this time.
  -index-relations-reader-uri string
    	A valid go-reader.Reader URI from which to read data for a relations candidate.
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/v3.Iterator URI. Supported iterator URI schemes are: cwd://,directory://,featurecollection://,file://,filelist://,geojsonl://,null://,repo:// (default "repo://")
  -names
    	Index the 'names' table
  -optimize
    	Attempt to optimize the database before closing connection (default true)
  -processes int
    	The number of concurrent processes to index data with (default 20)
  -properties
    	Index the 'properties' table
  -rtree
    	Index the 'rtree' table
  -search fts5
    	Index the 'search' table. If using the SQLite FTS5 full-text indexer requires the fts5 build tag.
  -spatial-tables
    	If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spatial-sqlite package.
  -spelunker
    	Index the 'spelunker' table
  -spelunker-tables
    	If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spelunker packages
  -spr
    	Index the 'spr' table
  -strict-alt-files
    	Be strict when indexing alt geometries (default true)
  -supersedes
    	Index the 'supersedes' table
  -verbose
    	Enable verbose (debug) logging
```
