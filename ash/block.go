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
	"fmt"
	"strings"
)

// Struct for script expression block.
type ScriptBlock struct {
	text      string
	condition *ScriptCase
	exprs     []*ScriptExpression
	blocks    []*ScriptBlock
}

// Type for block type.
type BlockType int

const (
	CaseBlock BlockType = iota
	ForBlock
)

// NewBlock creates new script block from script
// conditional block text.
func newBlock(text string) (*ScriptBlock, error) {
	b := new(ScriptBlock)
	b.text = text
	if !strings.Contains(text, "{") {
		return nil, fmt.Errorf("no script body")
	}
	// Main case.
	startBrace := strings.Index(text, "{")
	conText := text[:startBrace]
	conText = strings.ReplaceAll(conText, "{", "")
	con, err := newCase(strings.TrimSpace(conText))
	if err != nil {
		return nil, fmt.Errorf("fail to parse main case: %v", err)
	}
	b.condition = con
	// Body.
	body := textBetween(text, "{", "};")
	exprs := make([]*ScriptExpression, 0)
	innerBlocks := make([]*ScriptBlock, 0)
	if strings.Contains(body, "{") {
		innerBlocks, err = parseBlocks(body)
		if err != nil {
			return nil, fmt.Errorf("fail to parse inner blocks: %v", err)
		}
	} else {
		exprs, err = parseBody(body)
		if err != nil {
			return nil, fmt.Errorf("fail to parse body: %v", err)
		}
	}
	b.blocks = innerBlocks
	b.exprs = exprs
	return b, nil
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

// String returns block text.
func (b *ScriptBlock) String() string {
	return b.text
}

// SetVariable sets variable with specified name
// and value for block expressions.
func (b *ScriptBlock) SetVariable(n, v string) *ScriptBlock {
	baseText := b.text
	baseConText := b.Condition().String()
	text := strings.ReplaceAll(baseText, n, v)
	b, _ = newBlock(text)
	b.text = baseText
	b.condition, _ = newCase(baseConText)
	return b
	/*
	for _, e := range b.exprs {
		if e.Type() != Expr {
			continue
		}
		text := strings.ReplaceAll(e.burnExpr.String(), n, v)
		e.burnExpr, _ = syntax.NewSTDExpression(text)
	}
	for _, ib := range b.blocks {
		for _, e := range ib.exprs {
			if e.Type() != Expr {
				continue
			}
			text := strings.ReplaceAll(e.burnExpr.String(), n, v)
			e.burnExpr, _ = syntax.NewSTDExpression(text)
		}
		text := strings.ReplaceAll(ib.Condition().expr.String(), n, v)
		ib.Condition().expr, _ = syntax.NewSTDExpression(text)
	}
        */
}
