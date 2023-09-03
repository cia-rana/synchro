// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAfricaLusakaLocation  sync.Once
	cacheAfricaLusakaLocation *time.Location
)

type AfricaLusaka struct{}

func (AfricaLusaka) Location() *time.Location {
	onceAfricaLusakaLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Lusaka")
		if err != nil {
			panic(err)
		}
		cacheAfricaLusakaLocation = loc
	})
	return cacheAfricaLusakaLocation
}
