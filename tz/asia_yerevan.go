// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaYerevanLocation  sync.Once
	cacheAsiaYerevanLocation *time.Location
)

type AsiaYerevan struct{}

func (AsiaYerevan) Location() *time.Location {
	onceAsiaYerevanLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Yerevan")
		if err != nil {
			panic(err)
		}
		cacheAsiaYerevanLocation = loc
	})
	return cacheAsiaYerevanLocation
}
