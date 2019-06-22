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
	"strings"
)

// Struct for Ash script.
type Script struct {
	name      string
	args      []string
	text      string
	mainBlock *ScriptBlock
}

const (
	Comment_prefix = "#"
	Body_expr_sep  = ";"
	// Keywords.
	True_keyword   = "true"
	Echo_keyword   = "echo"
	Wait_keyword   = "wait"
	Rawdis_keyword = "rawdis"
)

// NewScript creates new Ash script from specified
// text, returns error in case of syntax error.
func NewScript(text string, args ...string) (*Script, error) {
	s := new(Script)
	s.args = args
	// Remove comment lines.
	for _, l := range strings.Split(text, "\n") {
		if strings.HasPrefix(l, Comment_prefix) {
			continue
		}
		s.text += l
	}
	// Insert args.
	for i := 1; i < len(s.args); i ++ {
		macro := fmt.Sprintf("@%d", i)
		s.text = strings.ReplaceAll(s.text, macro, s.args[i])
	}
	// Main block.
	mainBlock, err := parseBlock(s.text)
	if err != nil {
		return nil, fmt.Errorf("fail_to_parse_main_block:%v", err)
	}
	s.mainBlock = mainBlock
	return s, nil
}

// String returns script text body.
func (s *Script) String() string {
	return s.text
}

// MainBlock returns script main block.
func (s *Script) MainBlock() *ScriptBlock {
	return s.mainBlock
}
