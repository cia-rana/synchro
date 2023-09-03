// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaBogotaLocation  sync.Once
	cacheAmericaBogotaLocation *time.Location
)

type AmericaBogota struct{}

func (AmericaBogota) Location() *time.Location {
	onceAmericaBogotaLocation.Do(func() {
		loc, err := time.LoadLocation("America/Bogota")
		if err != nil {
			panic(err)
		}
		cacheAmericaBogotaLocation = loc
	})
	return cacheAmericaBogotaLocation
}
