// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAfricaNairobiLocation  sync.Once
	cacheAfricaNairobiLocation *time.Location
)

type AfricaNairobi struct{}

func (AfricaNairobi) Location() *time.Location {
	onceAfricaNairobiLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Nairobi")
		if err != nil {
			panic(err)
		}
		cacheAfricaNairobiLocation = loc
	})
	return cacheAfricaNairobiLocation
}
