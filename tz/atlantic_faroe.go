// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAtlanticFaroeLocation  sync.Once
	cacheAtlanticFaroeLocation *time.Location
)

type AtlanticFaroe struct{}

func (AtlanticFaroe) Location() *time.Location {
	onceAtlanticFaroeLocation.Do(func() {
		loc, err := time.LoadLocation("Atlantic/Faroe")
		if err != nil {
			panic(err)
		}
		cacheAtlanticFaroeLocation = loc
	})
	return cacheAtlanticFaroeLocation
}
