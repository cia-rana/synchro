// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaKathmanduLocation  sync.Once
	cacheAsiaKathmanduLocation *time.Location
)

type AsiaKathmandu struct{}

func (AsiaKathmandu) Location() *time.Location {
	onceAsiaKathmanduLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Kathmandu")
		if err != nil {
			panic(err)
		}
		cacheAsiaKathmanduLocation = loc
	})
	return cacheAsiaKathmanduLocation
}
