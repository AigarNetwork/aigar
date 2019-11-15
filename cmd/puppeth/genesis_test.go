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

package main

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/AigarNetwork/aigar/core"
	"github.com/davecgh/go-spew/spew"
)

// Tests the go-ethereum to Aleth chainspec conversion for the Stureby testnet.
func TestAlethSturebyConverter(t *testing.T) {
	blob, err := ioutil.ReadFile("testdata/stureby_geth.json")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}
	var genesis core.Genesis
	if err := json.Unmarshal(blob, &genesis); err != nil {
		t.Fatalf("failed parsing genesis: %v", err)
	}
	spec, err := newAlethGenesisSpec("stureby", &genesis)
	if err != nil {
		t.Fatalf("failed creating chainspec: %v", err)
	}

	expBlob, err := ioutil.ReadFile("testdata/stureby_aleth.json")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}
	expspec := &alethGenesisSpec{}
	if err := json.Unmarshal(expBlob, expspec); err != nil {
		t.Fatalf("failed parsing genesis: %v", err)
	}
	if !reflect.DeepEqual(expspec, spec) {
		t.Errorf("chainspec mismatch")
		c := spew.ConfigState{
			DisablePointerAddresses: true,
			SortKeys:                true,
		}
		exp := strings.Split(c.Sdump(expspec), "\n")
		got := strings.Split(c.Sdump(spec), "\n")
		for i := 0; i < len(exp) && i < len(got); i++ {
			if exp[i] != got[i] {
				t.Logf("got: %v\nexp: %v\n", exp[i], got[i])
			}
		}
	}
}

// Tests the go-ethereum to Parity chainspec conversion for the Stureby testnet.
func TestParitySturebyConverter(t *testing.T) {
	blob, err := ioutil.ReadFile("testdata/stureby_geth.json")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}
	var genesis core.Genesis
	if err := json.Unmarshal(blob, &genesis); err != nil {
		t.Fatalf("failed parsing genesis: %v", err)
	}
	spec, err := newParityChainSpec("stureby", &genesis, []string{})
	if err != nil {
		t.Fatalf("failed creating chainspec: %v", err)
	}

	expBlob, err := ioutil.ReadFile("testdata/stureby_parity.json")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}
	expspec := &parityChainSpec{}
	if err := json.Unmarshal(expBlob, expspec); err != nil {
		t.Fatalf("failed parsing genesis: %v", err)
	}
	expspec.Nodes = []string{}

	if !reflect.DeepEqual(expspec, spec) {
		t.Errorf("chainspec mismatch")
		c := spew.ConfigState{
			DisablePointerAddresses: true,
			SortKeys:                true,
		}
		exp := strings.Split(c.Sdump(expspec), "\n")
		got := strings.Split(c.Sdump(spec), "\n")
		for i := 0; i < len(exp) && i < len(got); i++ {
			if exp[i] != got[i] {
				t.Logf("got: %v\nexp: %v\n", exp[i], got[i])
			}
		}
	}
}
