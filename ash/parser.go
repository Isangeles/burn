/*
 * parser.go
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
	"strconv"
	"strings"

	"github.com/isangeles/burn"
	"github.com/isangeles/burn/syntax"
)

// parseBlocks creates script blocks from script
// body text.
func parseBlocks(text string) ([]*ScriptBlock, error) {
	blocks := make([]*ScriptBlock, 0)
	for _, e := range strings.Split(text, "};") {
		e = strings.TrimSpace(e)
		if len(e) < 1 {
			continue
		}
		e += "};"
		b, err := newBlock(e)
		if err != nil {
			return nil, fmt.Errorf("fail to parse inner block: %v", err)
		}
		blocks = append(blocks, b)
	}
	return blocks, nil
}

// parseBody creates script expressions and inner blocks
// from script block body text.
func parseBody(text string) ([]*ScriptExpression, error) {
	exprs := make([]*ScriptExpression, 0)
	for _, l := range strings.Split(text, BodyExprSep) {
		l = strings.TrimSpace(l)
		if len(l) < 1 {
			continue
		}
		switch {
		case strings.HasPrefix(l, EchoKeyword):
			l = textBetween(l, "(", ")")
			expr, err := syntax.NewSTDExpression(l)
			if err != nil {
				return nil, fmt.Errorf("fail to parse echo function: %v", err)
			}
			exprs = append(exprs, NewEchoMacro("", expr))
		case strings.HasPrefix(l, WaitKeyword):
			secText := textBetween(l, "(", ")")
			sec, err := strconv.ParseInt(secText, 32, 64)
			if err != nil {
				return nil, fmt.Errorf("fail to parse wait function: %v", err)
			}
			exprs = append(exprs, NewWaitMacro(sec*1000))
		default:
			expr, err := syntax.NewSTDExpression(l)
			if err != nil {
				return nil, fmt.Errorf("fail to parse expr: %v", err)
			}
			sExpr := NewExpression(expr)
			exprs = append(exprs, sExpr)
		}
	}
	return exprs, nil
}

// parseCaseExpr creates case expression from specified text.
func parseCaseExpr(text string) (burn.Expression, error) {
	switch {
	case strings.HasPrefix(text, RawdisKeyword):
		args := strings.Split(textBetween(text, "(", ")"), ",")
		if len(args) < 2 {
			return nil, fmt.Errorf("not enaught args for rawdis")
		}
		exprText := fmt.Sprintf("objectshow -o range -t %s -a %s",
			strings.TrimSpace(args[0]), strings.TrimSpace(args[1]))
		expr, err := syntax.NewSTDExpression(exprText)
		if err != nil {
			return nil, fmt.Errorf("fail to create rawdis exression: %v", err)
		}
		return expr, nil
	case strings.HasPrefix(text, OutKeyword):
		text = textBetween(text, "(", ")")
		expr, err := syntax.NewSTDExpression(text)
		if err != nil {
			return nil, fmt.Errorf("fail to create std expression: %v", err)
		}
		return expr, nil
	default:
		exprText := fmt.Sprintf("engineshow -o echo -a %s", text)
		expr, err := syntax.NewSTDExpression(exprText)
		if err != nil {
			return nil, fmt.Errorf("fail to create rawdis exression: %v", err)
		}
		return expr, nil
	}
}

// textBetween returns slice from specified text
// between first start and last end sequence or
// the same specified text if start or end sequence
// was not found.
func textBetween(text, start, end string) string {
	startID := strings.Index(text, start)
	if startID < 0 {
		return text
	}
	endID := strings.LastIndex(text, end)
	if endID < 0 {
		return text
	}
	return text[startID+len(start) : endID]
}
