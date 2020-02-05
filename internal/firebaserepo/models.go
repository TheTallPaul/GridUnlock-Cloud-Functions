package firebaserepo

import (
	"time"

	"google.golang.org/genproto/googleapis/type/latlng"
)

// A Ride is a Firestore document that stores information about a GridUnlock
// ride
type Ride struct {
	ID                      string         `firestore:"-"`
	Active                  bool           `firestore:"active"`
	ActualDropoffTime       time.Time      `firestore:"actual_dropoff_time"`
	ActualPickupTime        time.Time      `firestore:"actual_pickup_time"`
	DriverCancelled         bool           `firestore:"driver_cancelled"`
	DriverComments          string         `firestore:"driver_comments"`
	DriverID                string         `firestore:"driver_id"`
	DriverMarkedCompletion  bool           `firestore:"driver_marked_completion"`
	DriverRatingOfRider     float64        `firestore:"driver_rating_of_rider"`
	DropoffLocation         *latlng.LatLng `firestore:"dropoff_location"`
	PickupLocation          *latlng.LatLng `firestore:"pickup_location"`
	RequestTime             time.Time      `firestore:"request_time"`
	RiderCancelled          bool           `firestore:"rider_cancelled"`
	RiderComments           string         `firestore:"rider_comments"`
	RiderID                 string         `firestore:"rider_id"`
	RiderMarkedCompletion   bool           `firestore:"rider_marked_completion"`
	RiderRatingOfDriver     float64        `firestore:"rider_rating_of_driver"`
	RiderReachedDestination bool           `firestore:"rider_reached_destination"`
	ScheduledDropoffTime    time.Time      `firestore:"scheduled_dropoff_time"`
	ScheduledPickupTime     time.Time      `firestore:"scheduled_pickup_time"`
}

// A Drive is a Firestore document that stores information about a GridUnlock
// drive
type Drive struct {
	ID           string         `firestore:"-"`
	DepatureTime time.Time      `firestore:"departure_time"`
	Origin       *latlng.LatLng `firestore:"origin"`
	Destination  *latlng.LatLng `firestore:"destination"`
	DriverID     string         `firestore:"driver_id"`
}

// A User is a Firestore document that stores information about a GridUnlock
// user
type User struct {
	AcceptedRides  map[string]time.Time `firestore:"accepted_rides"`
	Description    string               `firestore:"description"`
	Email          string               `firestore:"email"`
	FirstName      string               `firestore:"first_name"`
	LastName       string               `firestore:"last_name"`
	Networks       map[string]time.Time `firestore:"networks"`
	RequestedRides map[string]time.Time `firestore:"requested_rides"`
}
