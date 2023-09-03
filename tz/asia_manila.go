// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaManilaLocation  sync.Once
	cacheAsiaManilaLocation *time.Location
)

type AsiaManila struct{}

func (AsiaManila) Location() *time.Location {
	onceAsiaManilaLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Manila")
		if err != nil {
			panic(err)
		}
		cacheAsiaManilaLocation = loc
	})
	return cacheAsiaManilaLocation
}
