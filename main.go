// Copyright (c) 2019 Cascade October LLC.

// Module gridunlock provides packages that can query Firebase for unmatched
// riders and drivers for GridUnlock.
package gridunlock

import (
	"GridUnlock-Cloud-Functions.internal/firebaserepo"
)

func Config() string {
  return "gridunlock config"
}

func main() {
	FetchRides()
}
