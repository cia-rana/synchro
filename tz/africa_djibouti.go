// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAfricaDjiboutiLocation  sync.Once
	cacheAfricaDjiboutiLocation *time.Location
)

type AfricaDjibouti struct{}

func (AfricaDjibouti) Location() *time.Location {
	onceAfricaDjiboutiLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Djibouti")
		if err != nil {
			panic(err)
		}
		cacheAfricaDjiboutiLocation = loc
	})
	return cacheAfricaDjiboutiLocation
}
