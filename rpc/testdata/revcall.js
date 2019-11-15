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

// This test checks reverse calls.

--> {"jsonrpc":"2.0","id":2,"method":"test_callMeBack","params":["foo",[1]]}
<-- {"jsonrpc":"2.0","id":1,"method":"foo","params":[1]}
--> {"jsonrpc":"2.0","id":1,"result":"my result"}
<-- {"jsonrpc":"2.0","id":2,"result":"my result"}
