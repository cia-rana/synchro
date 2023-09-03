// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAtlanticAzoresLocation  sync.Once
	cacheAtlanticAzoresLocation *time.Location
)

type AtlanticAzores struct{}

func (AtlanticAzores) Location() *time.Location {
	onceAtlanticAzoresLocation.Do(func() {
		loc, err := time.LoadLocation("Atlantic/Azores")
		if err != nil {
			panic(err)
		}
		cacheAtlanticAzoresLocation = loc
	})
	return cacheAtlanticAzoresLocation
}
