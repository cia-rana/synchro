// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaFamagustaLocation  sync.Once
	cacheAsiaFamagustaLocation *time.Location
)

type AsiaFamagusta struct{}

func (AsiaFamagusta) Location() *time.Location {
	onceAsiaFamagustaLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Famagusta")
		if err != nil {
			panic(err)
		}
		cacheAsiaFamagustaLocation = loc
	})
	return cacheAsiaFamagustaLocation
}
