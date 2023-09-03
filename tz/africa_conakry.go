// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAfricaConakryLocation  sync.Once
	cacheAfricaConakryLocation *time.Location
)

type AfricaConakry struct{}

func (AfricaConakry) Location() *time.Location {
	onceAfricaConakryLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Conakry")
		if err != nil {
			panic(err)
		}
		cacheAfricaConakryLocation = loc
	})
	return cacheAfricaConakryLocation
}
