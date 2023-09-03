// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaHalifaxLocation  sync.Once
	cacheAmericaHalifaxLocation *time.Location
)

type AmericaHalifax struct{}

func (AmericaHalifax) Location() *time.Location {
	onceAmericaHalifaxLocation.Do(func() {
		loc, err := time.LoadLocation("America/Halifax")
		if err != nil {
			panic(err)
		}
		cacheAmericaHalifaxLocation = loc
	})
	return cacheAmericaHalifaxLocation
}
