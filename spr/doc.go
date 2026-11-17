// Package spr defines the standard places result interfaces for Who's On First records.
//
// A [StandardPlacesResult] exposes the common data needed to describe one
// record. It takes inspiration from Flickr's [standard photo response],
// which provides enough information to display a photo with attribution
// and a link back to its page. [StandardPlacesResults] represents a
// collection of these results.
//
// These interfaces do not define a data exchange format. Implementations
// may provide serialization, but are not required to do so.
//
// [standard photo response]: https://code.flickr.net/2008/08/19/standard-photos-response-apis-for-civilized-age
package spr
