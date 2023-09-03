// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceEuropeSan_MarinoLocation  sync.Once
	cacheEuropeSan_MarinoLocation *time.Location
)

type EuropeSan_Marino struct{}

func (EuropeSan_Marino) Location() *time.Location {
	onceEuropeSan_MarinoLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/San_Marino")
		if err != nil {
			panic(err)
		}
		cacheEuropeSan_MarinoLocation = loc
	})
	return cacheEuropeSan_MarinoLocation
}
