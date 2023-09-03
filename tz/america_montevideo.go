// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaMontevideoLocation  sync.Once
	cacheAmericaMontevideoLocation *time.Location
)

type AmericaMontevideo struct{}

func (AmericaMontevideo) Location() *time.Location {
	onceAmericaMontevideoLocation.Do(func() {
		loc, err := time.LoadLocation("America/Montevideo")
		if err != nil {
			panic(err)
		}
		cacheAmericaMontevideoLocation = loc
	})
	return cacheAmericaMontevideoLocation
}
