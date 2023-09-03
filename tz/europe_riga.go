// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceEuropeRigaLocation  sync.Once
	cacheEuropeRigaLocation *time.Location
)

type EuropeRiga struct{}

func (EuropeRiga) Location() *time.Location {
	onceEuropeRigaLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Riga")
		if err != nil {
			panic(err)
		}
		cacheEuropeRigaLocation = loc
	})
	return cacheEuropeRigaLocation
}
