// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaBoa_VistaLocation  sync.Once
	cacheAmericaBoa_VistaLocation *time.Location
)

type AmericaBoa_Vista struct{}

func (AmericaBoa_Vista) Location() *time.Location {
	onceAmericaBoa_VistaLocation.Do(func() {
		loc, err := time.LoadLocation("America/Boa_Vista")
		if err != nil {
			panic(err)
		}
		cacheAmericaBoa_VistaLocation = loc
	})
	return cacheAmericaBoa_VistaLocation
}
