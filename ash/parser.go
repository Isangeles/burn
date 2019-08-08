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

// parseBlock creates script block from script expressions
// conditional block.
func parseBlock(text string) (*ScriptBlock, error) {
	if !strings.Contains(text, "{") {
		return nil, fmt.Errorf("no_script_body")
	}
	// Main case.
	startBrace := strings.Index(text, "{")
	conText := text[:startBrace]
	conText = strings.ReplaceAll(conText, "{", "")
	con, err := parseCase(strings.TrimSpace(conText))
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_main_case:%v", err)
	}
	// Body.
	body := textBetween(text, "{", "}")
	exprs := make([]*ScriptExpression, 0)
	innerBlocks := make([]*ScriptBlock, 0)
	if strings.Contains(body, "{") {
		innerBlocks, err = parseInnerBlocks(body)
		if err != nil {
			return nil, fmt.Errorf("fail_to_parse_inner_blocks:%v", err)
		}
	} else {
		exprs, err = parseBody(body)
		if err != nil {
			return nil, fmt.Errorf("fail_to_parse_body:%v", err)
		}
	}
	block := NewBlock(con, exprs, innerBlocks)
	return block, nil
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
				return nil, fmt.Errorf("fail_to_parse_echo_function:%v", err)
			}
			exprs = append(exprs, NewEchoMacro("", expr))
		case strings.HasPrefix(l, WaitKeyword):
			secText := textBetween(l, "(", ")")
			sec, err := strconv.ParseInt(secText, 32, 64)
			if err != nil {
				return nil, fmt.Errorf("fail_to_parse_wait_function:%v", err)
			}
			exprs = append(exprs, NewWaitMacro(sec*1000))
		default:
			expr, err := syntax.NewSTDExpression(l)
			if err != nil {
				return nil, fmt.Errorf("fail_to_parse_expr:%v", err)
			}
			sExpr := NewExpression(expr)
			exprs = append(exprs, sExpr)
		}
	}
	return exprs, nil
}

// parseInnerBlocks creates script blocks from script
// body text.
func parseInnerBlocks(text string) ([]*ScriptBlock, error) {
	blocks := make([]*ScriptBlock, 0)
	for _, e := range strings.Split(text, "};") {
		if len(e) < 1 {
			continue
		}
		e += "}"
		b, err := parseBlock(e)
		if err != nil {
			return nil, fmt.Errorf("fail_to_parse_inner_block:%v", err)
		}
		blocks = append(blocks, b)
	}
	return blocks, nil
}

// parseCase creates script case from specified text.
func parseCase(text string) (*ScriptCase, error) {
	if len(text) < 1 || strings.HasPrefix(text, TrueKeyword) {
		c := NewCase(new(syntax.STDExpression), "", True)
		return c, nil
	}
	var compType ComparisonType
	switch {
	case strings.Contains(text, "<"):
		compType = Less
		exprs := strings.Split(text, "<")
		expr, err := parseCaseExpr(exprs[0])
		if err != nil {
			return nil, fmt.Errorf("fail_to_parse_case_expression:%v", err)
		}
		res := strings.TrimSpace(exprs[1])
		c := NewCase(expr, res, compType)
		return c, nil
	default:
		return nil, fmt.Errorf("unknown case expression:%s", text)
	}
}

// parseCaseExpr creates case expression from specified text.
func parseCaseExpr(text string) (burn.Expression, error) {
	switch {
	case strings.HasPrefix(text, RawdisKeyword):
		args := strings.Split(textBetween(text, "(", ")"), ",")
		if len(args) < 2 {
			return nil, fmt.Errorf("not enaught args for rawdis")
		}
		exprText := fmt.Sprintf("charman -o show -t %s -a range %s",
			strings.TrimSpace(args[0]), strings.TrimSpace(args[1]))
		expr, err := syntax.NewSTDExpression(exprText)
		if err != nil {
			return nil, fmt.Errorf("fail_to_create_rawdis_exression:%v", err)
		}
		return expr, nil
	default:
		expr, err := syntax.NewSTDExpression(text)
		if err != nil {
			return nil, fmt.Errorf("fail_to_create_std_expression:%v", err)
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
	return text[startID+1 : endID]
}
