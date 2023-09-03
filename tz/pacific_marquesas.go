// Code generated by tzgen. DO NOT EDIT.
// Package model contains the types.
package tz

import "time"
import "sync"

var (
	oncePacificMarquesasLocation  sync.Once
	cachePacificMarquesasLocation *time.Location
)

type PacificMarquesas struct{}

func (PacificMarquesas) Location() *time.Location {
	oncePacificMarquesasLocation.Do(func() {
		loc, err := time.LoadLocation("Pacific/Marquesas")
		if err != nil {
			panic(err)
		}
		cachePacificMarquesasLocation = loc
	})
	return cachePacificMarquesasLocation
}
