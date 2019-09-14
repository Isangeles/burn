/*
 * stdexpression.go
 *
 * Copyright 2018 Dariusz Sikora <dev@isangeles.pl>
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

package syntax

import (
	"fmt"
	"strings"

	"github.com/isangeles/burn"
)

const (
	StdArgPipe = " |a "
	StdTarPipe = " |t "
)

// Type for standard expressions.
// Syntax:
// Arg pipe delimiter: ' |a '.
// Target arg pipe delimiter: ' |t ',
// e.g. 'moduleman -o show -t areachars [area ID] |t charman -o show -t position' - shows
// positions of all characters in area with specified ID.
type STDExpression struct {
	text     string
	commands []burn.Command
	etype    burn.ExpressionType
}

// NewStdExpression creates new standard expression from
// specified text.
func NewSTDExpression(text string) (*STDExpression, error) {
	exp := new(STDExpression)
	exp.text = strings.TrimSpace(text)
	exp.etype = stdExpressionType(text)
	if exp.Type() == burn.NoExp {
		cmd, err := NewSTDCommand(exp.text)
		if err != nil {
			return exp, fmt.Errorf("fail to build expression command: %v", err)
		}
		exp.commands = append(exp.commands, cmd)
		return exp, nil
	}
	cmdsText := splitExprCommands(exp.text, exp.Type())
	for _, cmdText := range cmdsText {
		cmd, err := NewSTDCommand(strings.TrimSpace(cmdText))
		if err != nil {
			return nil, fmt.Errorf("fail to build expression command: %v", err)
		}
		exp.commands = append(exp.commands, cmd)
	}
	return exp, nil
}

// Commands returns all expression commands.
func (exp *STDExpression) Commands() []burn.Command {
	return exp.commands
}

// Type returns expression type.
func (exp *STDExpression) Type() burn.ExpressionType {
	return exp.etype
}

// Retruns expression text.
func (exp *STDExpression) String() string {
	return exp.text
}

// stdExpressionType returns expression type for specified
// text.
func stdExpressionType(t string) burn.ExpressionType {
	switch {
	case strings.Contains(t, StdTarPipe):
		return burn.PipeTarExp
	case strings.Contains(t, StdArgPipe):
		return burn.PipeArgExp
	default:
		return burn.NoExp	
	}
}

// splitExprCommands splits specified expression text into
// separate commands text.
func splitExprCommands(c string, t burn.ExpressionType) []string {
	switch t {
	case burn.PipeTarExp:
		return strings.Split(c, StdTarPipe)
	case burn.PipeArgExp:
		return strings.Split(c, StdArgPipe)
	default:
		return strings.Fields(c)
	}
}
