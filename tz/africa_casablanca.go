// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAfricaCasablancaLocation  sync.Once
	cacheAfricaCasablancaLocation *time.Location
)

type AfricaCasablanca struct{}

func (AfricaCasablanca) Location() *time.Location {
	onceAfricaCasablancaLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Casablanca")
		if err != nil {
			panic(err)
		}
		cacheAfricaCasablancaLocation = loc
	})
	return cacheAfricaCasablancaLocation
}
