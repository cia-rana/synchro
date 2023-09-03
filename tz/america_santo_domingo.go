// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaSanto_DomingoLocation  sync.Once
	cacheAmericaSanto_DomingoLocation *time.Location
)

type AmericaSanto_Domingo struct{}

func (AmericaSanto_Domingo) Location() *time.Location {
	onceAmericaSanto_DomingoLocation.Do(func() {
		loc, err := time.LoadLocation("America/Santo_Domingo")
		if err != nil {
			panic(err)
		}
		cacheAmericaSanto_DomingoLocation = loc
	})
	return cacheAmericaSanto_DomingoLocation
}
