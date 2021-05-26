package hashtree

import "time"

var timeLocation = time.Local

// SetTimeLocation sets time location to hashtree
// Must be call before all Hashtree created
func SetTimeLocation(l *time.Location) {
	timeLocation = l
}
