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

// This test checks the behavior of batches with invalid elements.
// Empty batches are not allowed. Batches may contain junk.

--> []
<-- {"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"empty batch"}}

--> [1]
<-- [{"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"invalid request"}}]

--> [1,2,3]
<-- [{"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"invalid request"}},{"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"invalid request"}},{"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"invalid request"}}]

--> [{"jsonrpc":"2.0","id":1,"method":"test_echo","params":["foo",1]},55,{"jsonrpc":"2.0","id":2,"method":"unknown_method"},{"foo":"bar"}]
<-- [{"jsonrpc":"2.0","id":1,"result":{"String":"foo","Int":1,"Args":null}},{"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"invalid request"}},{"jsonrpc":"2.0","id":2,"error":{"code":-32601,"message":"the method unknown_method does not exist/is not available"}},{"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"invalid request"}}]
