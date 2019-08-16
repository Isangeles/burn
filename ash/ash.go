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
	"strings"

	"github.com/isangeles/burn"
)

const (
	SCRIPT_FILE_EXT = ".ash"
)

// Run runs specified script.
func Run(scr *Script) error {
	for _, b := range scr.Blocks() {
		if scr.Stopped() {
			return nil
		}
		err := runBlock(scr, b)
		if err != nil {
			return fmt.Errorf("fail to run script block: %v", err)
		}
	}
	return nil
}

// runBlock runs specfied script block.
func runBlock(scr *Script, blk *ScriptBlock) error {
	for {
		if scr.Stopped() {
			return nil
		}
		// Check condition.
		m, out, err := meet(blk.Condition())
		if err != nil {
			return fmt.Errorf("fail to check block condition: %v", err)
		}
		if !m {
			break
		}
		vars := strings.Fields(out)
		if blk.Condition().compType != For {
			vars = append(vars, "")
		}
		for _, v := range vars {
			// Condition variable.
			argid := blk.Condition().argID
			if argid > 0 {
				blk = blk.SetVariable(fmt.Sprintf("@%d", argid), v)
			}
			// Execute expressions.
			for _, e := range blk.Expressions() {
				if e.Type() == WaitMacro {
					time.Sleep(time.Duration(e.WaitTime()) * time.Millisecond)
					continue
				}
				r, o := burn.HandleExpression(e.BurnExpr())
				if r != 0 {
					return fmt.Errorf("fail to run expr: '%s': [%d]%s",
						e.BurnExpr().String(), r, o)
				}
				if e.Type() == EchoMacro {
					fmt.Printf("%s\n", o)
				}
			}
			// Inner blocks.
			for _, b := range blk.Blocks() {
				err := runBlock(scr, b)
				if err != nil {
					return err
				}
			}	
		}
	}
	return nil
}

// meet checks if specified case is meet and returns
// condition command output.
// Returns error if burn return error result(!=0)
// for case expression.
func meet(c *ScriptCase) (bool, string, error) {
	if c.compType == True {
		return true, "", nil
	}
	r, o := burn.HandleExpression(c.Expression())
	if r != 0 {
		return false, "", fmt.Errorf("fail to run condition exp: '%s': [%d]%s\n",
			c.Expression().String(), r, o)
	}
	meet, err := c.CorrectRes(o)
	if err != nil {
		return false, "", fmt.Errorf("fail to check result: %v", err)
	}
	return meet, o, nil
}
