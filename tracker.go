/*******************************************************************************
*
* Copyright 2018 Stefan Majewsky <majewsky@gmx.net>
*
* This program is free software: you can redistribute it and/or modify it under
* the terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
*******************************************************************************/

package main

//Uint64Tracker tracks the state of a single uint64 variable (the index of the
//current slide).
type Uint64Tracker struct {
	input   chan<- inquiry
	value   uint64 //only ever accessed from the tracker's own goroutine, so no mutex necessary
	pending []inquiry
}

//An inquiry is sent to Uint64Tracker by other goroutines to inquire about the
//state of the uint64 tracked by it, or to update its value.
//
//If Return is not nil, it works like this: If HasPreviousValue is false, the
//tracker immediately sends the current value into Return. Otherwise, it sends
//the current value into Return as soon as it is different from PreviousValue.
//The tracker always sends into Return exactly once.
//
//If Return is nil, then the NewValue is written into the uint64 tracked by the
//Tracker.
type inquiry struct {
	HasPreviousValue bool
	PreviousValue    uint64
	Return           chan<- uint64
	NewValue         uint64
}

//Run starts the tracker in a separate goroutine.
func (t *Uint64Tracker) Run(initialValue uint64) {
	t.value = initialValue
	t.pending = nil
	inputChan := make(chan inquiry, 10)
	t.input = inputChan
	go func() {
		for i := range inputChan {
			if i.Return == nil {
				t.handleSet(i.NewValue)
			} else {
				t.handleGet(i)
			}
		}
	}()
}

func (t *Uint64Tracker) handleSet(newValue uint64) {
	t.value = newValue

	var stillPending []inquiry
	for _, i := range t.pending {
		if i.HasPreviousValue && i.PreviousValue == t.value {
			stillPending = append(stillPending, i)
		} else {
			i.Return <- t.value
		}
	}
	t.pending = stillPending
}

func (t *Uint64Tracker) handleGet(i inquiry) {
	if i.HasPreviousValue && i.PreviousValue == t.value {
		t.pending = append(t.pending, i)
	} else {
		i.Return <- t.value
	}
}

//Get returns the current value of the tracked uint64.
func (t *Uint64Tracker) Get() uint64 {
	resultChan := make(chan uint64)
	t.input <- inquiry{
		HasPreviousValue: false,
		Return:           resultChan,
	}
	return <-resultChan
}

//GetWhenDifferentFrom is like Get(), but stalls while the value of the tracked uint64 is
//equal to the given value.
func (t *Uint64Tracker) GetWhenDifferentFrom(previousValue uint64) uint64 {
	resultChan := make(chan uint64)
	t.input <- inquiry{
		HasPreviousValue: true,
		PreviousValue:    previousValue,
		Return:           resultChan,
	}
	return <-resultChan
}

//Set updates the current value of the tracked uint64.
func (t *Uint64Tracker) Set(value uint64) {
	t.input <- inquiry{
		Return:   nil,
		NewValue: value,
	}
}
