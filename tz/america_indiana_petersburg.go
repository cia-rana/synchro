// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaIndianaPetersburgLocation  sync.Once
	cacheAmericaIndianaPetersburgLocation *time.Location
)

type AmericaIndianaPetersburg struct{}

func (AmericaIndianaPetersburg) Location() *time.Location {
	onceAmericaIndianaPetersburgLocation.Do(func() {
		loc, err := time.LoadLocation("America/Indiana/Petersburg")
		if err != nil {
			panic(err)
		}
		cacheAmericaIndianaPetersburgLocation = loc
	})
	return cacheAmericaIndianaPetersburgLocation
}
