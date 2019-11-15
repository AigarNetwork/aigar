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
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/AigarNetwork/aigar/core/asm"
	cli "gopkg.in/urfave/cli.v1"
)

var disasmCommand = cli.Command{
	Action:    disasmCmd,
	Name:      "disasm",
	Usage:     "disassembles evm binary",
	ArgsUsage: "<file>",
}

func disasmCmd(ctx *cli.Context) error {
	if len(ctx.Args().First()) == 0 {
		return errors.New("filename required")
	}

	fn := ctx.Args().First()
	in, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	code := strings.TrimSpace(string(in))
	fmt.Printf("%v\n", code)
	return asm.PrintDisassembled(code)
}
