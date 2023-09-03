// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaAdakLocation  sync.Once
	cacheAmericaAdakLocation *time.Location
)

type AmericaAdak struct{}

func (AmericaAdak) Location() *time.Location {
	onceAmericaAdakLocation.Do(func() {
		loc, err := time.LoadLocation("America/Adak")
		if err != nil {
			panic(err)
		}
		cacheAmericaAdakLocation = loc
	})
	return cacheAmericaAdakLocation
}
