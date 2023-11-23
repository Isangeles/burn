/*
 * objectset_test.go
 *
 * Copyright 2023 Dariusz Sikora <ds@isangeles.dev>
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

package burn

import (
	"testing"

	"github.com/isangeles/flame/character"
)

// TestObjectSetChapter tests chapter option for objectset.
func TestObjectSetChapter(t *testing.T) {
	// Create objects
	ob1 := character.New(charData)
	ob2 := character.New(charData)
	// Create command
	cmd := testCommand{
		optionArgs: []string{"chapter"},
		targetArgs: []string{
			ob1.ID() + IDSerialSep + ob1.Serial(),
			ob2.ID() + IDSerialSep + ob2.Serial(),
		},
		args: []string{"testChapter"},
	}
	// Test
	res, out := objectset(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if len(out) > 0 {
		t.Errorf("Command output not empty: '%s'", out)
	}
	if ob1.ChapterID() != cmd.args[0] {
		t.Errorf("Chapter ID was not set for object 1: '%s' != '%s'",
			ob1.ChapterID(), cmd.args[0])
	}
	if ob2.ChapterID() != cmd.args[0] {
		t.Errorf("Chapter ID was not set for object 2: '%s' != '%s'",
			ob2.ChapterID(), cmd.args[0])
	}
}
