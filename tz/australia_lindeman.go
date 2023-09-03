// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAustraliaLindemanLocation  sync.Once
	cacheAustraliaLindemanLocation *time.Location
)

type AustraliaLindeman struct{}

func (AustraliaLindeman) Location() *time.Location {
	onceAustraliaLindemanLocation.Do(func() {
		loc, err := time.LoadLocation("Australia/Lindeman")
		if err != nil {
			panic(err)
		}
		cacheAustraliaLindemanLocation = loc
	})
	return cacheAustraliaLindemanLocation
}
