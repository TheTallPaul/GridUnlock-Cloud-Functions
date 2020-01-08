package ridematch

// A coord is a latitude/longitude pairing
type coord struct {
	lat float64
	lng float64
}

// A riderStartEnd is a starting point and destination pairing for a rider
type riderStartEnd struct {
	riderID string
	start   coord
	end     coord
}

// A driverRoute is the sequential route a driver takes
type driverRoute struct {
	driverID string
	route    []coord
}
