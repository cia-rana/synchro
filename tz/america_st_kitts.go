// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAmericaSt_KittsLocation  sync.Once
	cacheAmericaSt_KittsLocation *time.Location
)

type AmericaSt_Kitts struct{}

func (AmericaSt_Kitts) Location() *time.Location {
	onceAmericaSt_KittsLocation.Do(func() {
		loc, err := time.LoadLocation("America/St_Kitts")
		if err != nil {
			panic(err)
		}
		cacheAmericaSt_KittsLocation = loc
	})
	return cacheAmericaSt_KittsLocation
}
