/*
 * case.go
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

// Struct for script case.
type ScriptCase struct {
	text     string
	expr     burn.Expression
	expRes   string
	compType ComparisonType
	argID    int
}

// Type for script case types.
type ComparisonType int

const (
	Greater ComparisonType = iota
	Equal
	Less
	Dif
	For
	True
)

// newCase creates new script case.
func newCase(text string) (*ScriptCase, error) {
	c := new(ScriptCase)
	c.text = text
	if len(text) < 1 || strings.HasPrefix(text, TrueKeyword) {
		c.expr = new(syntax.STDExpression)
		c.compType = True
		return c, nil
	}
	switch {
	case strings.Contains(text, "<"):
		c.compType = Less
		exprs := strings.Split(text, "<")
		expr, err := parseCaseExpr(exprs[0])
		if err != nil {
			return nil, fmt.Errorf("fail to parse case expression: %v", err)
		}
		c.expr = expr
		res := strings.TrimSpace(exprs[1])
		c.expRes = res
		return c, nil
	case strings.Contains(text, "!="):
		c.compType = Dif
		exprs := strings.Split(text, "!=")
		expr, err := parseCaseExpr(exprs[0])
		if err != nil {
			return nil, fmt.Errorf("fail to parse case expression: %v", err)
		}
		c.expr = expr
		res := strings.TrimSpace(exprs[1])
		c.expRes = res
		return c, nil
	case strings.HasPrefix(text, "for"):	
		c.compType = For
		exprText := textBetween(text, "for(", ")")
		if strings.HasPrefix(exprText, "@") {
			argID, err := strconv.Atoi(exprText[1:2])
			if err != nil {
				return nil, fmt.Errorf("fail to parse case arg: %v", err)
			}
			c.argID = argID
			exprStart := strings.Index(exprText, "=")
			exprText = exprText[exprStart+1:]
			exprText = strings.TrimSpace(exprText)
		}
		expr, err := parseCaseExpr(exprText)
		if err != nil {
			return nil, fmt.Errorf("fail to parse case expression: %v", err)
		}
		c.expr = expr
		return c, nil
	default:
		return nil, fmt.Errorf("unknown case expression:%s", text)
	}
}

// Expression returns case expression.
func (c *ScriptCase) Expression() burn.Expression {
	return c.expr
}

// CorrectRes checks if specified result value is
// correct.
func (c *ScriptCase) CorrectRes(r string) (bool, error) {
	switch c.compType {
	case Greater:
		n, err := strconv.ParseFloat(r, 64)
		if err != nil {
			return false, fmt.Errorf("fail to parse result: %v", err)
		}
		exp, err := strconv.ParseFloat(c.expRes, 64)
		if err != nil {
			return false, fmt.Errorf("fail to parse expected result: %v", err)
		}
		return n > exp, nil
	case Less:
		n, err := strconv.ParseFloat(r, 64)
		if err != nil {
			return false, fmt.Errorf("fail to parse result: %v", err)
		}
		exp, err := strconv.ParseFloat(c.expRes, 64)
		if err != nil {
			return false, fmt.Errorf("fail to parse expected result: %v", err)
		}
		return n < exp, nil
	case Equal:
		return r == c.expRes, nil
	case Dif:
		return r != c.expRes, nil
	case For:
		return len(r) > 0, nil
	case True:
		return true, nil
	default:
		return false, nil
	}
}

// String returns condition text.
func (c *ScriptCase) String() string {
	return c.text
}
