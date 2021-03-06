/*
 * script.go
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

// Struct for Ash script.
type Script struct {
	name      string
	args      map[int]string
	text      string
	blocks    []*ScriptBlock
	stopped   bool
}

const (
	CommentPrefix = "#"
	BodyExprSep   = ";"
	VarPrefix     = "@"
	// Keywords.
	TrueKeyword   = "true"
	EchoKeyword   = "echo"
	WaitKeyword   = "wait"
	RawdisKeyword = "rawdis"
	OutKeyword    = "out"
	EndKeyword    = "end"
)

// NewScript creates new Ash script from specified
// text, returns error in case of syntax error.
func NewScript(name, text string, args ...string) (*Script, error) {
	s := new(Script)
	s.name = name
	s.args = make(map[int]string)
	for i, a := range args {
		s.args[i] = a
	}
	// Remove comment lines.
	for _, l := range strings.Split(text, "\n") {
		if strings.HasPrefix(l, CommentPrefix) {
			continue
		}
		if strings.HasPrefix(l, VarPrefix) {
			continue
		}
		s.text += l
	}
	// Parse args declaration.
	for _, l := range strings.Split(text, "\n") {
		if strings.HasPrefix(l, CommentPrefix) {
			continue
		}
		if !strings.HasPrefix(l, VarPrefix) {
			continue
		}
		argID, err := strconv.Atoi(l[1:2])
		if err != nil {
			return nil, fmt.Errorf("fail to parse var declaration: %v", err)
		}
		val := l[strings.Index(l, "=")+1:]
		val = strings.TrimSpace(val)
		if strings.HasPrefix(val, "\"") || !strings.HasPrefix(val, OutKeyword) {
			s.args[argID] = val
			continue
		}
		val = textBetween(val, OutKeyword+"(", ")")
		expr, err := syntax.NewSTDExpression(val)
		if err != nil {
			return nil, fmt.Errorf("fail to build var expr: %v", err)
		}
		r, o := burn.HandleExpression(expr)
		if r != 0 {
			return nil, fmt.Errorf("fail to run var expr: '%s': [%d]%s",
				expr, r, o)
		}
		s.args[argID] = o
	}
	// Insert args.
	for i := 1; i <= len(s.args); i ++ {
		macro := fmt.Sprintf("@%d", i)
		s.text = strings.ReplaceAll(s.text, macro, s.args[i])
	}
	// Expression blocks.
	blocks, err := parseBlocks(s.text)
	if err != nil {
		return nil, fmt.Errorf("fail to parse script blocks: %v", err)
	}
	s.blocks = blocks
	return s, nil
}

// Name returns script name.
func (s *Script) Name() string {
	return s.name
}

// String returns script text body.
func (s *Script) String() string {
	return s.text
}

// Blocks returns script blocks with expressions.
func (s *Script) Blocks() []*ScriptBlock {
	return s.blocks
}

// Stop toggles script stop flag.
func (s *Script) Stop(stop bool) {
	s.stopped = stop
}

// Stopped checks if script was stopped.
func (s *Script) Stopped() bool {
	return s.stopped
}
