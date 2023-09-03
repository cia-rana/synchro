// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaMazatlanLocation  sync.Once
	cacheAmericaMazatlanLocation *time.Location
)

type AmericaMazatlan struct{}

func (AmericaMazatlan) Location() *time.Location {
	onceAmericaMazatlanLocation.Do(func() {
		loc, err := time.LoadLocation("America/Mazatlan")
		if err != nil {
			panic(err)
		}
		cacheAmericaMazatlanLocation = loc
	})
	return cacheAmericaMazatlanLocation
}
