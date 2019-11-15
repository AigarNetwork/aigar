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

// This test calls the test_echo method.

--> {"jsonrpc": "2.0", "id": 2, "method": "test_echo", "params": []}
<-- {"jsonrpc":"2.0","id":2,"error":{"code":-32602,"message":"missing value for required argument 0"}}

--> {"jsonrpc": "2.0", "id": 2, "method": "test_echo", "params": ["x"]}
<-- {"jsonrpc":"2.0","id":2,"error":{"code":-32602,"message":"missing value for required argument 1"}}

--> {"jsonrpc": "2.0", "id": 2, "method": "test_echo", "params": ["x", 3]}
<-- {"jsonrpc":"2.0","id":2,"result":{"String":"x","Int":3,"Args":null}}

--> {"jsonrpc": "2.0", "id": 2, "method": "test_echo", "params": ["x", 3, {"S": "foo"}]}
<-- {"jsonrpc":"2.0","id":2,"result":{"String":"x","Int":3,"Args":{"S":"foo"}}}

--> {"jsonrpc": "2.0", "id": 2, "method": "test_echoWithCtx", "params": ["x", 3, {"S": "foo"}]}
<-- {"jsonrpc":"2.0","id":2,"result":{"String":"x","Int":3,"Args":{"S":"foo"}}}
