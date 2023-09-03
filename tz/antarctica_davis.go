// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAntarcticaDavisLocation  sync.Once
	cacheAntarcticaDavisLocation *time.Location
)

type AntarcticaDavis struct{}

func (AntarcticaDavis) Location() *time.Location {
	onceAntarcticaDavisLocation.Do(func() {
		loc, err := time.LoadLocation("Antarctica/Davis")
		if err != nil {
			panic(err)
		}
		cacheAntarcticaDavisLocation = loc
	})
	return cacheAntarcticaDavisLocation
}
