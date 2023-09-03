// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	oncePacificTarawaLocation  sync.Once
	cachePacificTarawaLocation *time.Location
)

type PacificTarawa struct{}

func (PacificTarawa) Location() *time.Location {
	oncePacificTarawaLocation.Do(func() {
		loc, err := time.LoadLocation("Pacific/Tarawa")
		if err != nil {
			panic(err)
		}
		cachePacificTarawaLocation = loc
	})
	return cachePacificTarawaLocation
}
