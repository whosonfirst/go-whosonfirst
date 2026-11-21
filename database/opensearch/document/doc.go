// Package document transforms Who's On First records into documents suitable
// for indexing in OpenSearch.
//
// The transformation helpers handle both complete GeoJSON features and
// Spelunker v1 documents. A Spelunker v1 document stores fields at the top
// level rather than under a GeoJSON properties object.
package document
