//  Copyright 2018 The go-ethereum Authors
//  Copyright 2019 The go-aigar Authors
//  This file is part of the go-aigar library.
//
//  The go-aigar library is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Lesser General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  The go-aigar library is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
//  GNU Lesser General Public License for more details.
//
//  You should have received a copy of the GNU Lesser General Public License
//  along with the go-aigar library. If not, see <http://www.gnu.org/licenses/>.

package metrics

import (
	"math"
	"sync"
	"sync/atomic"
)

// EWMAs continuously calculate an exponentially-weighted moving average
// based on an outside source of clock ticks.
type EWMA interface {
	Rate() float64
	Snapshot() EWMA
	Tick()
	Update(int64)
}

// NewEWMA constructs a new EWMA with the given alpha.
func NewEWMA(alpha float64) EWMA {
	return &StandardEWMA{alpha: alpha}
}

// NewEWMA1 constructs a new EWMA for a one-minute moving average.
func NewEWMA1() EWMA {
	return NewEWMA(1 - math.Exp(-5.0/60.0/1))
}

// NewEWMA5 constructs a new EWMA for a five-minute moving average.
func NewEWMA5() EWMA {
	return NewEWMA(1 - math.Exp(-5.0/60.0/5))
}

// NewEWMA15 constructs a new EWMA for a fifteen-minute moving average.
func NewEWMA15() EWMA {
	return NewEWMA(1 - math.Exp(-5.0/60.0/15))
}

// EWMASnapshot is a read-only copy of another EWMA.
type EWMASnapshot float64

// Rate returns the rate of events per second at the time the snapshot was
// taken.
func (a EWMASnapshot) Rate() float64 { return float64(a) }

// Snapshot returns the snapshot.
func (a EWMASnapshot) Snapshot() EWMA { return a }

// Tick panics.
func (EWMASnapshot) Tick() {
	panic("Tick called on an EWMASnapshot")
}

// Update panics.
func (EWMASnapshot) Update(int64) {
	panic("Update called on an EWMASnapshot")
}

// NilEWMA is a no-op EWMA.
type NilEWMA struct{}

// Rate is a no-op.
func (NilEWMA) Rate() float64 { return 0.0 }

// Snapshot is a no-op.
func (NilEWMA) Snapshot() EWMA { return NilEWMA{} }

// Tick is a no-op.
func (NilEWMA) Tick() {}

// Update is a no-op.
func (NilEWMA) Update(n int64) {}

// StandardEWMA is the standard implementation of an EWMA and tracks the number
// of uncounted events and processes them on each tick.  It uses the
// sync/atomic package to manage uncounted events.
type StandardEWMA struct {
	uncounted int64 // /!\ this should be the first member to ensure 64-bit alignment
	alpha     float64
	rate      float64
	init      bool
	mutex     sync.Mutex
}

// Rate returns the moving average rate of events per second.
func (a *StandardEWMA) Rate() float64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.rate * float64(1e9)
}

// Snapshot returns a read-only copy of the EWMA.
func (a *StandardEWMA) Snapshot() EWMA {
	return EWMASnapshot(a.Rate())
}

// Tick ticks the clock to update the moving average.  It assumes it is called
// every five seconds.
func (a *StandardEWMA) Tick() {
	count := atomic.LoadInt64(&a.uncounted)
	atomic.AddInt64(&a.uncounted, -count)
	instantRate := float64(count) / float64(5e9)
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.init {
		a.rate += a.alpha * (instantRate - a.rate)
	} else {
		a.init = true
		a.rate = instantRate
	}
}

// Update adds n uncounted events.
func (a *StandardEWMA) Update(n int64) {
	atomic.AddInt64(&a.uncounted, n)
}
