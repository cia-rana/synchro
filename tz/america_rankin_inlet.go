// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaRankin_InletLocation  sync.Once
	cacheAmericaRankin_InletLocation *time.Location
)

type AmericaRankin_Inlet struct{}

func (AmericaRankin_Inlet) Location() *time.Location {
	onceAmericaRankin_InletLocation.Do(func() {
		loc, err := time.LoadLocation("America/Rankin_Inlet")
		if err != nil {
			panic(err)
		}
		cacheAmericaRankin_InletLocation = loc
	})
	return cacheAmericaRankin_InletLocation
}
