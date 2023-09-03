// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceEuropeSimferopolLocation  sync.Once
	cacheEuropeSimferopolLocation *time.Location
)

type EuropeSimferopol struct{}

func (EuropeSimferopol) Location() *time.Location {
	onceEuropeSimferopolLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Simferopol")
		if err != nil {
			panic(err)
		}
		cacheEuropeSimferopolLocation = loc
	})
	return cacheEuropeSimferopolLocation
}
