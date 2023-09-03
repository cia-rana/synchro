// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaAqtauLocation  sync.Once
	cacheAsiaAqtauLocation *time.Location
)

type AsiaAqtau struct{}

func (AsiaAqtau) Location() *time.Location {
	onceAsiaAqtauLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Aqtau")
		if err != nil {
			panic(err)
		}
		cacheAsiaAqtauLocation = loc
	})
	return cacheAsiaAqtauLocation
}
