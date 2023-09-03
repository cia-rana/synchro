// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaGrand_TurkLocation  sync.Once
	cacheAmericaGrand_TurkLocation *time.Location
)

type AmericaGrand_Turk struct{}

func (AmericaGrand_Turk) Location() *time.Location {
	onceAmericaGrand_TurkLocation.Do(func() {
		loc, err := time.LoadLocation("America/Grand_Turk")
		if err != nil {
			panic(err)
		}
		cacheAmericaGrand_TurkLocation = loc
	})
	return cacheAmericaGrand_TurkLocation
}
