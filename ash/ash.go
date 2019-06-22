/*
 * ash.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
 * MA 02110-1301, USA.
 *
 *
 */

package ash

import (
	"fmt"
	"time"

	"github.com/isangeles/burn"
)

const (
	SCRIPT_FILE_EXT = ".ash"
)

// Run runs specified script.
func Run(scr *Script) error {
	for meet(scr.MainBlock().Condition()) {
		err := runBlock(scr.MainBlock())
		if err != nil {
			return fmt.Errorf("fail_to_run_expr_block:%v", err)
		}
	}
	return nil
}

// runBlock runs specfied script block.
func runBlock(blk *ScriptBlock) error {
	if !meet(blk.Condition()) {
		return nil
	}
	for _, e := range blk.Expressions() {
		if e.Type() == Wait_macro {
				time.Sleep(time.Duration(e.WaitTime()) * time.Millisecond)
				continue
		}
		r, o := burn.HandleExpression(e.BurnExpr())
		if r != 0 {
				return fmt.Errorf("fail_to_run_expr:'%s':[%d]%s",
					e.BurnExpr().String(), r, o)
		}
		if e.Type() == Echo_macro {
			fmt.Printf("%s\n", o)
		}
	}
	for _, b := range blk.Blocks() {
		err := runBlock(b)
		if err != nil {
			return err
		}
	}
	return nil
}

// meet checks if specified case is meet.
func meet(c *ScriptCase) bool {
	_, o := burn.HandleExpression(c.Expression())
	return c.CorrectRes(o)
}
