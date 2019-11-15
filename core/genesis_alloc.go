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

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xe2\xe1\x94\x1eN4a\xe0\x87;om!\xda\xccQp\\J\ucc29\xfd\x8b\t\x05\x00\x00\x00\x00\x00\x00\x00\x00\x00"
const testnetAllocData = "\xc0"
const rinkebyAllocData = "\xc0"
const goerliAllocData = "\xc0"
