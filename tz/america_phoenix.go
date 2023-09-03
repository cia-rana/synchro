// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaPhoenixLocation  sync.Once
	cacheAmericaPhoenixLocation *time.Location
)

type AmericaPhoenix struct{}

func (AmericaPhoenix) Location() *time.Location {
	onceAmericaPhoenixLocation.Do(func() {
		loc, err := time.LoadLocation("America/Phoenix")
		if err != nil {
			panic(err)
		}
		cacheAmericaPhoenixLocation = loc
	})
	return cacheAmericaPhoenixLocation
}
