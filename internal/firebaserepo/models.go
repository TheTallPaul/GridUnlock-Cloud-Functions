package firebaserepo

type Ride struct {
	ActualDropoffTime       string  `json:"actual_dropoff_time"`
	ActualPickupTime        string  `json:"actual_pickup_time"`
	DriverCancelled         bool    `json:"driver_cancelled"`
	DriverComments          string  `json:"driver_comments"`
	DriverMarkedCompletion  bool    `json:"driver_marked_completion"`
	DriverRatingOfRider     float64 `json:"driver_rating_of_rider"`
	DropoffLocation         string  `json:"dropoff_location"`
	PickupLocation          string  `json:"pickup_location"`
	RequestTime             string  `json:"request_time"`
	RiderCancelled          bool    `json:"rider_cancelled"`
	RiderComments           string  `json:"rider_comments"`
	RiderMarkedCompletion   bool    `json:"rider_marked_completion"`
	RiderRatingOfDriver     float64 `json:"rider_rating_of_driver"`
	RiderReachedDestination bool    `json:"rider_reached_destination"`
	ScheduledDropoffTime    string  `json:"scheduled_dropoff_time"`
	ScheduledPickupTime     string  `json:"scheduled_pickup_time"`
}
