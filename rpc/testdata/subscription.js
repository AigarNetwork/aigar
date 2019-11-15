/*
 * // Copyright 2018 The go-ethereum Authors
 * // Copyright 2019 The go-aigar Authors
 * // This file is part of the go-aigar library.
 * //
 * // The go-aigar library is free software: you can redistribute it and/or modify
 * // it under the terms of the GNU Lesser General Public License as published by
 * // the Free Software Foundation, either version 3 of the License, or
 * // (at your option) any later version.
 * //
 * // The go-aigar library is distributed in the hope that it will be useful,
 * // but WITHOUT ANY WARRANTY; without even the implied warranty of
 * // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * // GNU Lesser General Public License for more details.
 * //
 * // You should have received a copy of the GNU Lesser General Public License
 * // along with the go-aigar library. If not, see <http://www.gnu.org/licenses/>.
 */

// This test checks basic subscription support.

--> {"jsonrpc":"2.0","id":1,"method":"nftest_subscribe","params":["someSubscription",5,1]}
<-- {"jsonrpc":"2.0","id":1,"result":"0x1"}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":1}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":2}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":3}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":4}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":5}}

--> {"jsonrpc":"2.0","id":2,"method":"nftest_echo","params":[11]}
<-- {"jsonrpc":"2.0","id":2,"result":11}
