package flags

// ExistentialFlag represents a Who's On First existential flag: 1 means true,
// 0 means false, and -1 means unknown or to be determined.
type ExistentialFlag interface {
	StringFlag() string
	Flag() int64
	// Return a boolean value indicating whether the flag is true.
	IsTrue() bool
	// Return a boolean value indicating whether the flag is false.
	IsFalse() bool
	// Return a boolean value indicating whether the flag is true or false (not -1).
	IsKnown() bool
	// Return a boolean value indicating whether the flag matches any record in a set of ExistentialFlag instances.
	MatchesAny(...ExistentialFlag) bool
	// Return a boolean value indicating whether the flag matches all records in a set of ExistentialFlag instances.
	MatchesAll(...ExistentialFlag) bool
	String() string
}

// AlternateGeometryFlag identifies whether to match alternate geometries and which sources to allow.
type AlternateGeometryFlag interface {
	// Return a boolean value indicating whether the flag matches any record in a set of AlternateGeometryFlag instances.
	MatchesAny(...AlternateGeometryFlag) bool
	// Return a boolean value indicating whether the flag matches all records in a set of AlternateGeometryFlag instances.
	MatchesAll(...AlternateGeometryFlag) bool
	IsAlternateGeometry() bool
	Label() string
	String() string
}

// DateFlag provides a common interface for querying Who's On First records using date ranges.
type DateFlag interface {
	// Return a boolean value indicating whether the flag matches any record in a set of DateFlag instances.
	MatchesAll(...DateFlag) bool
	// Return a boolean value indicating whether the flag matches all records in a set of DateFlag instances.
	MatchesAny(...DateFlag) bool
	// Returns min, max numeric values representing to inner range of allowable dates for a DateFlag instance.
	InnerRange() (*int64, *int64)
	// Returns min, max numeric values representing to outer range of allowable dates for a DateFlag instance.
	OuterRange() (*int64, *int64)
	String() string
}

// PlacetypeFlag provides a common interface for querying Who's On First records using well-defined placetypes.
type PlacetypeFlag interface {
	// Return a boolean value indicating whether the flag matches any record in a set of PlacetypeFlag instances.
	MatchesAny(...PlacetypeFlag) bool
	// Return a boolean value indicating whether the flag matches all records in a set of PlacetypeFlag instances.
	MatchesAll(...PlacetypeFlag) bool
	Placetype() string
	String() string
}

// CustomFlag provides a common interface for querying Who's On First records with non-standard (custom) flag types.
type CustomFlag interface {
	// Return a boolean value indicating whether the flag matches any record in a set of CustomFlag instances.
	MatchesAny(...CustomFlag) bool
	// Return a boolean value indicating whether the flag matches all records in a set of CustomFlag instances.
	MatchesAll(...CustomFlag) bool
	String() string
}
