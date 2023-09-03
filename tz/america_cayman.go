// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaCaymanLocation  sync.Once
	cacheAmericaCaymanLocation *time.Location
)

type AmericaCayman struct{}

func (AmericaCayman) Location() *time.Location {
	onceAmericaCaymanLocation.Do(func() {
		loc, err := time.LoadLocation("America/Cayman")
		if err != nil {
			panic(err)
		}
		cacheAmericaCaymanLocation = loc
	})
	return cacheAmericaCaymanLocation
}
