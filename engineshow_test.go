/*
 * engineshow_test.go
 *
 * Copyright 2025 Dariusz Sikora <ds@isangeles.dev>
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
	"fmt"
	"testing"
)

// TestEngineShowNoOptions tests handling of the command without any options/args.
func TestEngineShowNoOptions(t *testing.T) {
	// Create command
	cmd := testCommand{
		tool: EngineShow,
	}
	// Test
	res, out := engineshow(cmd)
	if res != 2 {
		t.Errorf("Command result invalid: %d != 2", res)
	}
	expOut := fmt.Sprintf("%s: no option args", EngineShow)
	if out != expOut {
		t.Errorf("Command output invalid: '%s' != '%s'", out, expOut)
	}
}

// TestEngineShowVersion tests handling of the version option.
func TestEngineShowVersion(t *testing.T) {
	// Create command
	cmd := testCommand{
		tool:       EngineShow,
		optionArgs: []string{"version"},
	}
	// Test
	res, out := engineshow(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if len(out) < 1 {
		t.Errorf("Command output empty")
	}
}

// TestEngineShowEcho tests handling of the echo option.
func TestEngineShowEcho(t *testing.T) {
	msg := "test"
	// Create command
	cmd := testCommand{
		tool:       EngineShow,
		optionArgs: []string{"echo"},
		args:       []string{msg},
	}
	// Test
	res, out := engineshow(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if out != msg {
		t.Errorf("Command output invalid: '%s' != '%s'", out, msg)
	}
}
