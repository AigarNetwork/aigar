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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Aigar network.
var MainnetBootnodes = []string{
	"enode://e648a63eab88bfc158539c1a4aa61b6615791d3191e0023ea6b0b2fdbce0193c318af3a453a9d1c3ffcd8900b1364ccc55d974cd3389d1389a51e05220578c47@185.227.108.100:30301", //Aigar
	"enode://7ca44e04a03c288de4c7528681b3e4cda43eb88b446022cff78831537316ddc08bfce1450c771dfeccb5fbb4fe38d8f3a4d83c9bd99c9551714c42b0e9e554ff@185.143.145.108:30301", //Aigar
	"enode://d29308776f7c93f42e33d2a24d5eaaa3ea79b000cea0780e917e6f0ba2a69a36767bf5989f3b52632a2f94623ec1d5d70abad3b8121a768832ca4c537b540420@45.32.232.42:30301",    //Possum
	"enode://116752351622099f5643d782b297cdb45d4235a0240b27345b58be5c1e8c674f2f994d5c2726acc868b8ad04eb1b4a1d3e16a2e3c855c433f6329783919d93f8@167.179.98.182:30303",  //Guzelka
	"enode://c4f5dc83bb7c0f1c91eb2da56fcd24360a00b9155d4f1db9530e0971f2737d86f11bb30cc2d5408e0d234b712b6871e88591c954ba6d1ac791449263194f69bb@178.57.220.60:30301",   //Bob Slay
	"enode://5675619e577a3efd023087aa5f5ae0828f99b922f49a4b972c66b6e5315606fb846e8f4aaa00ed9fa87691873c02a66259f3a4954120ab1b129044d95ae5b860@185.87.193.151:30301",  //Bob Slay
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes []string

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes []string

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes []string

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes []string
