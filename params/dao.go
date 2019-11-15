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

import (
	"math/big"

	"github.com/AigarNetwork/aigar/common"
)

// DAOForkBlockExtra is the block header extra-data field to set for the DAO fork
// point and a number of consecutive blocks to allow fast/light syncers to correctly
// pick the side they want  ("dao-hard-fork").
var DAOForkBlockExtra = common.FromHex("0x64616f2d686172642d666f726b")

// DAOForkExtraRange is the number of consecutive blocks from the DAO fork point
// to override the extra-data in to prevent no-fork attacks.
var DAOForkExtraRange = big.NewInt(10)

// DAORefundContract is the address of the refund contract to send DAO balances to.
var DAORefundContract = common.HexToAddress("0xbf4ed7b27f1d666546e30d74d50d173d20bca754")

// DAODrainList is the list of accounts whose full balances will be moved into a
// refund contract at the beginning of the dao-fork block.
func DAODrainList() []common.Address {
	return []common.Address{
		common.HexToAddress("0xA9B45025377F00E2903fc962fDe970d8a9F87f56"),
		common.HexToAddress("0xc59Cf2C2E2D142CeD9765B2cc31D871287198DB1"),
		common.HexToAddress("0x2116b681eC5374041c4309965A9a41D1eF31c0A0"),
		common.HexToAddress("0x91b2B1eee3C9FFeF70148fD2ABDf67cEE6c73Ab1"),
		common.HexToAddress("0x2D67e63B2bFC0Ce6DDD5C087e046F276323f7Ea9"),
	}
}
