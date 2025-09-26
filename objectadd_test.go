/*
 * objectadd_test.go
 *
 * Copyright 2023-2025 Dariusz Sikora <ds@isangeles.dev>
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

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/item"
	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/data/res"
)

// TestObjectAddNoOptions tests handling objectadd command
// with no options/args.
func TestObjectAddNoOptions(t *testing.T) {
	// Create module.
	Module = flame.NewModule(res.ModuleData{})
	// Create command.
	cmd := testCommand{
		tool: ObjectAdd,
	}
	// Test.
	res, out := objectadd(cmd)
	if res != 2 {
		t.Errorf("Command result invalid: %d != 2", res)
	}
	expOut := fmt.Sprintf("%s: no option args", ObjectAdd)
	if out != expOut {
		t.Errorf("Command output invalid: '%s' != '%s'", out, expOut)
	}
}

// TestObjectAddEquipment tests handling objectadd equipement option.
func TestObjectAddEquipment(t *testing.T) {
	// Create module.
	Module = flame.NewModule(res.ModuleData{})
	weaponData := res.WeaponData{ID: "weapon", Slots: []res.ItemSlotData{res.ItemSlotData{"hand"}}}
	it := item.NewWeapon(weaponData)
	char := character.New(charData)
	char.Inventory().AddItem(it)
	// Create command.
	cmd := testCommand{
		tool: ObjectAdd,
		optionArgs: []string{"eq"},
		targetArgs: []string{fmt.Sprintf("%s#%s", char.ID(), char.Serial())},
		args: []string{"hand", fmt.Sprintf("%s#%s", it.ID(), it.Serial())},
	}
	// Test.
	res, out := objectadd(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if len(out) > 0 {
		t.Errorf("Command output invalid: '%s' != ''", out)
	}
	if !char.Equipment().Equiped(it) {
		t.Errorf("Item was not equiped")
	}
	equiped := false
	for _, s := range char.Equipment().Slots() {
		if s.Type() == item.Hand && s.Item() == it {
			equiped = true
		}
	}
	if !equiped {
		t.Errorf("Item was not equiped on proper slot")
	}
}

// TestObjectAddItem tests handling objectadd item option.
func TestObjectAddItem(t *testing.T) {
	// Create module.
	data := res.ModuleData{}
	weaponData := res.WeaponData{ID: "weapon"}
	data.Resources.Weapons = append(data.Resources.Weapons, weaponData)
	Module = flame.NewModule(data)
	char := character.New(charData)
	// Create command.
	cmd := testCommand{
		tool: ObjectAdd,
		optionArgs: []string{"item"},
		targetArgs: []string{fmt.Sprintf("%s#%s", char.ID(), char.Serial())},
		args: []string{"weapon", "2"},
	}
	// Test.
	res, out := objectadd(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if len(out) > 0 {
		t.Errorf("Command output invalid: %s != ''", out)
	}
	itemsCount := 0
	for _, it := range char.Inventory().Items() {
		if it.ID() == weaponData.ID {
			itemsCount++
		}
	}
	if itemsCount != 2 {
		t.Errorf("Items not found in the command target inventory")
	}
}

// TestObjectAddQuest tests handling objectadd quest option.
func TestObjectAddQuest(t *testing.T) {
	// Create module.
	data := res.ModuleData{}
	questData := res.QuestData{ID: "quest"}
	data.Resources.Quests = append(data.Resources.Quests, questData)
	Module = flame.NewModule(data)
	char := character.New(charData)
	// Create command.
	cmd := testCommand{
		tool: ObjectAdd,
		optionArgs: []string{"quest"},
		targetArgs: []string{fmt.Sprintf("%s#%s", char.ID(), char.Serial())},
		args: []string{"quest"},
	}
	// Test.
	res, out := objectadd(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if len(out) > 0 {
		t.Errorf("Command output invalid: %s != ''", out)
	}
	found := false
	for _, q := range char.Journal().Quests() {
		if q.ID() == questData.ID {
			found = true
		}
	}
	if !found {
		t.Errorf("Quest not added to the command target quest log")
	}
}

// TestObjectAddEffect tests handling of objectadd effect option.
func TestObjectAddEffect(t *testing.T) {
	// Create module.
	data := res.ModuleData{}
	effectData := res.EffectData{ID: "testEffect"}
	data.Resources.Effects = append(data.Resources.Effects, effectData)
	Module = flame.NewModule(data)
	char := character.New(charData)
	// Create command.
	cmd := testCommand{
		tool: ObjectAdd,
		optionArgs: []string{"effect"},
		targetArgs: []string{fmt.Sprintf("%s#%s", char.ID(), char.Serial())},
		args: []string{"testEffect"},
	}
	// Test.
	res, out := objectadd(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if len(out) > 0 {
		t.Errorf("Command output invalid: %s != ''", out)
	}
	found := false
	for _, e := range char.Effects() {
		if e.ID() == effectData.ID {
			found = true
		}
	}
	if !found {
		t.Errorf("Effect not applied on the command target")
	}
}
