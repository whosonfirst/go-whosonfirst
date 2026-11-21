package geo

// IsValidMinLatitude reports whether lat is at least -90.
func IsValidMinLatitude(lat float64) bool {
	return lat >= -90.00
}

// IsValidMaxLatitude reports whether lat is at most 90.
func IsValidMaxLatitude(lat float64) bool {
	return lat <= 90.00
}

// IsValidLatitude reports whether lat is between -90 and 90, inclusive.
func IsValidLatitude(lat float64) bool {
	return IsValidMinLatitude(lat) && IsValidMaxLatitude(lat)
}

// IsValidMinLongitude reports whether lon is at least -180.
func IsValidMinLongitude(lon float64) bool {
	return lon >= -180.00
}

// IsValidMaxLongitude reports whether lon is at most 180.
func IsValidMaxLongitude(lon float64) bool {
	return lon <= 180.00
}

// IsValidLongitude reports whether lon is between -180 and 180, inclusive.
func IsValidLongitude(lon float64) bool {
	return IsValidMinLongitude(lon) && IsValidMaxLongitude(lon)
}
