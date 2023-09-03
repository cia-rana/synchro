// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	onceAsiaJerusalemLocation  sync.Once
	cacheAsiaJerusalemLocation *time.Location
)

type AsiaJerusalem struct{}

func (AsiaJerusalem) Location() *time.Location {
	onceAsiaJerusalemLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Jerusalem")
		if err != nil {
			panic(err)
		}
		cacheAsiaJerusalemLocation = loc
	})
	return cacheAsiaJerusalemLocation
}
