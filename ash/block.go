/*
 * block.go
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

)

// Struct for script expression block.
type ScriptBlock struct {
  condition *ScriptCase
  exprs     []*ScriptExpression
  blocks    []*ScriptBlock
}

// NewBlock creates new script block.
func NewBlock(con *ScriptCase, exprs []*ScriptExpression, blocks []*ScriptBlock) *ScriptBlock {
  b := new(ScriptBlock)
  b.condition = con
  b.exprs = exprs
  b.blocks = blocks
  return b
}

// Conditon returns block condition case.
func (b *ScriptBlock) Condition() *ScriptCase {
  return b.condition
}

// Expressions returns all expression within block.
func (b *ScriptBlock) Expressions() []*ScriptExpression {
  return b.exprs
}

// Blocks returns all inner blocks.
func (b *ScriptBlock) Blocks() []*ScriptBlock {
  return b.blocks
}
