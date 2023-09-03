// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceEuropeBucharestLocation  sync.Once
	cacheEuropeBucharestLocation *time.Location
)

type EuropeBucharest struct{}

func (EuropeBucharest) Location() *time.Location {
	onceEuropeBucharestLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Bucharest")
		if err != nil {
			panic(err)
		}
		cacheEuropeBucharestLocation = loc
	})
	return cacheEuropeBucharestLocation
}
