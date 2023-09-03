// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAntarcticaTrollLocation  sync.Once
	cacheAntarcticaTrollLocation *time.Location
)

type AntarcticaTroll struct{}

func (AntarcticaTroll) Location() *time.Location {
	onceAntarcticaTrollLocation.Do(func() {
		loc, err := time.LoadLocation("Antarctica/Troll")
		if err != nil {
			panic(err)
		}
		cacheAntarcticaTrollLocation = loc
	})
	return cacheAntarcticaTrollLocation
}
