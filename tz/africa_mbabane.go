// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAfricaMbabaneLocation  sync.Once
	cacheAfricaMbabaneLocation *time.Location
)

type AfricaMbabane struct{}

func (AfricaMbabane) Location() *time.Location {
	onceAfricaMbabaneLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Mbabane")
		if err != nil {
			panic(err)
		}
		cacheAfricaMbabaneLocation = loc
	})
	return cacheAfricaMbabaneLocation
}
