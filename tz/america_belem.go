// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaBelemLocation  sync.Once
	cacheAmericaBelemLocation *time.Location
)

type AmericaBelem struct{}

func (AmericaBelem) Location() *time.Location {
	onceAmericaBelemLocation.Do(func() {
		loc, err := time.LoadLocation("America/Belem")
		if err != nil {
			panic(err)
		}
		cacheAmericaBelemLocation = loc
	})
	return cacheAmericaBelemLocation
}
