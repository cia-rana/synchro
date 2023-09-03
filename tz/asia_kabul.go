// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaKabulLocation  sync.Once
	cacheAsiaKabulLocation *time.Location
)

type AsiaKabul struct{}

func (AsiaKabul) Location() *time.Location {
	onceAsiaKabulLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Kabul")
		if err != nil {
			panic(err)
		}
		cacheAsiaKabulLocation = loc
	})
	return cacheAsiaKabulLocation
}
