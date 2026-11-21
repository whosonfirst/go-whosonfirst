## wof-derivatives-server

Serves derivative representations of Who's On First records over HTTP.

```
$> wof-derivatives-server -h
A simple HTTP server-based interface for serving different machine-reabable representations (derivatives) of Who's On First documents.

Usage:
	 wof-derivatives-server[options]

Valid options are:
  -cors-allowed-origin value
    	Zero or more allowed origins for CORS requests.
  -enable-cors
    	Enable CORS support.
  -navplace-max-features int
    	The maximum number of WOF IDs allowed in a NavPlace request. (default 10)
  -path-geojson string
    	The default path to serve GeoJSON requests from. (default "/id/{id}/geojson")
  -path-geojson-alt value
    	Zero or more alternate paths to serve GeoJSON requests from.
  -path-geojsonld string
    	The default path to serve GeoJSONLD requests from. (default "/id/{id}/geojsonld")
  -path-geojsonld-alt value
    	Zero or more alternate paths to serve GeoJSONLD requests from.
  -path-navaplace-alt value
    	Zero or more alternate paths to serve IIIF NavPlace requests from.
  -path-navplace string
    	The default path to serve IIIF NavPlace requests from. (default "/id/{id}/navplace")
  -path-select string
    	The default path to serve select requests from. (default "/id/{id}/select")
  -path-select-alt value
    	Zero or more alternate paths to serve select requests from.
  -path-spr string
    	The default path to serve standard place result (SPR) requests from. (default "/id/{id}/spr")
  -path-spr-alt value
    	Zero or more alternate paths to serve standard place result (SPR) requests from.
  -path-svg string
    	The default path to serve SVG requests from. (default "/id/{id}/svg")
  -path-svg-alt value
    	Zero or more alternate paths to serve SVG requests from.
  -path-wkt string
    	The default path to serve WKT requests from. (default "/id/{id}/wkt")
  -path-wkt-alt value
    	Zero or more alternate paths to serve WKT requests from.
  -provider-uri string
    	A registered whosonfirst/go-whosonfirst-derivatives.Provider URI. (default "reader://?reader-uri=https://data.whosonfirst.org")
  -reader-uri string
    	If not-empty and the -provider-uri flag contains the string '{reader_uri}' the value of this flag will be used to replace the '{reader_uri}' placeholder in the -provider-uri flag.
  -server-uri string
    	A registered aaronland/go-http-server.Server URI. (default "http://localhost:8080")
  -verbose
    	Enable verbose (debug) logging.
```
