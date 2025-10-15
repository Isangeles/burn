/*
 * engineexport.go
 *
 * Copyright 2019-2025 Dariusz Sikora <ds@isangeles.dev>
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
	"path/filepath"

	"github.com/isangeles/flame/data"
	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/serial"
)

// engineexport handles engineexport command.
func engineexport(cmd Command) (int, string) {
	if cmd.OptionArgs() == nil || len(cmd.OptionArgs()) < 1 {
		return 2, fmt.Sprintf("%s: no option args", EngineExport)
	}
	switch cmd.OptionArgs()[0] {
	case "module", "mod":
		return engineexportModule(cmd)
	case "character", "char":
		return engineexportCharacter(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", EngineExport,
			cmd.OptionArgs()[0])
	}
}

// engineexportModule handles module option for engineexport.
func engineexportModule(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: not enough args for: %s",
			EngineExport, cmd.OptionArgs()[0])
	}
	if Module == nil {
		return 3, fmt.Sprintf("%s: no game set", EngineExport)
	}
	path := filepath.Join("data", "modules", cmd.Args()[0])
	err := data.ExportModuleDir(path, Module.Data())
	if err != nil {
		return 3, fmt.Sprintf("%s: unable to export module: %v",
			EngineExport, err)
	}
	return 0, ""
}

// engineexportCharacter handles character option for engineexport.
func engineexportCharacter(cmd Command) (int, string) {
	if len(cmd.Args()) < 0 {
		return 3, fmt.Sprintf("%s: not enough args for: %s",
			EngineExport, cmd.OptionArgs()[0])
	}
	if Module == nil {
		return 3, fmt.Sprintf("%s: no module set", EngineExport)
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not charatcer",
				ObjectSet, ob.ID(), ob.Serial())
		}
		objects = append(objects, char)
	}
	for _, o := range objects {
		path := filepath.Join(Module.Conf().CharactersPath(), o.ID()+o.Serial())
		err := data.ExportCharacters(path, o.Data())
		if err != nil {
			return 3, fmt.Sprintf("%s: %s#%s: unable to export: %v", EngineExport,
				o.ID(), o.Serial(), err)
		}
	}
	return 0, ""
}
