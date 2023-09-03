// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaDawsonLocation  sync.Once
	cacheAmericaDawsonLocation *time.Location
)

type AmericaDawson struct{}

func (AmericaDawson) Location() *time.Location {
	onceAmericaDawsonLocation.Do(func() {
		loc, err := time.LoadLocation("America/Dawson")
		if err != nil {
			panic(err)
		}
		cacheAmericaDawsonLocation = loc
	})
	return cacheAmericaDawsonLocation
}
