// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceEuropeIsle_of_ManLocation  sync.Once
	cacheEuropeIsle_of_ManLocation *time.Location
)

type EuropeIsle_of_Man struct{}

func (EuropeIsle_of_Man) Location() *time.Location {
	onceEuropeIsle_of_ManLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Isle_of_Man")
		if err != nil {
			panic(err)
		}
		cacheEuropeIsle_of_ManLocation = loc
	})
	return cacheEuropeIsle_of_ManLocation
}
