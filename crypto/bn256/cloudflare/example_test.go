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

package bn256

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExamplePair(t *testing.T) {
	// This implements the tripartite Diffie-Hellman algorithm from "A One
	// Round Protocol for Tripartite Diffie-Hellman", A. Joux.
	// http://www.springerlink.com/content/cddc57yyva0hburb/fulltext.pdf

	// Each of three parties, a, b and c, generate a private value.
	a, _ := rand.Int(rand.Reader, Order)
	b, _ := rand.Int(rand.Reader, Order)
	c, _ := rand.Int(rand.Reader, Order)

	// Then each party calculates g₁ and g₂ times their private value.
	pa := new(G1).ScalarBaseMult(a)
	qa := new(G2).ScalarBaseMult(a)

	pb := new(G1).ScalarBaseMult(b)
	qb := new(G2).ScalarBaseMult(b)

	pc := new(G1).ScalarBaseMult(c)
	qc := new(G2).ScalarBaseMult(c)

	// Now each party exchanges its public values with the other two and
	// all parties can calculate the shared key.
	k1 := Pair(pb, qc)
	k1.ScalarMult(k1, a)

	k2 := Pair(pc, qa)
	k2.ScalarMult(k2, b)

	k3 := Pair(pa, qb)
	k3.ScalarMult(k3, c)

	// k1, k2 and k3 will all be equal.

	require.Equal(t, k1, k2)
	require.Equal(t, k1, k3)

	require.Equal(t, len(np), 4) //Avoid gometalinter varcheck err on np
}
