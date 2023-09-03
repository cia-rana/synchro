// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaFortalezaLocation  sync.Once
	cacheAmericaFortalezaLocation *time.Location
)

type AmericaFortaleza struct{}

func (AmericaFortaleza) Location() *time.Location {
	onceAmericaFortalezaLocation.Do(func() {
		loc, err := time.LoadLocation("America/Fortaleza")
		if err != nil {
			panic(err)
		}
		cacheAmericaFortalezaLocation = loc
	})
	return cacheAmericaFortalezaLocation
}
