/*
 * burn_test.go
 *
 * Copyright 2022-2023 Dariusz Sikora <ds@isangeles.dev>
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

	"github.com/isangeles/flame/data/res"
)

var (
	weaponData = res.WeaponData{ID: "weapon", Slots: []res.ItemSlotData{res.ItemSlotData{"hand"}}}
)

// TestAddToolHandler tests adding new tool handler.
func TestAddToolHandler(t *testing.T) {
	// Create and add new handler.
	handler := func(cmd Command) (int, string) {
		return 0, "ok"
	}
	AddToolHandler("handler", handler)
	// Test.
	cmd := testCommand{tool: "handler"}
	res, out := HandleCommand(cmd)
	if res != 0 {
		t.Errorf("Invalid test result: %d != 0", res)
	}
	if out != "ok" {
		t.Errorf("Invalid test output: '%s' != 'ok'", out)
	}
}
