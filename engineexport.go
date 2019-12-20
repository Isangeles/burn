/*
 * engineexport.go
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

package burn

import (
	"fmt"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/config"
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/module/character"
)

// engineexport handles enginesave command.
func engineexport(cmd Command) (int, string) {
	if len(cmd.OptionArgs()) < 1 {
		return 2, fmt.Sprintf("%s: no option args", EngineExport)
	}
	switch cmd.OptionArgs()[0] {
	case "game":
		return engineexportGame(cmd)
	case "character", "char":
		return engineexportCharacter(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", EngineExport,
			cmd.OptionArgs()[0])
	}
}

// engineexportGame handles game option for enginesave.
func engineexportGame(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: not enought args for: %s",
			EngineExport, cmd.OptionArgs()[0])
	}
	if flame.Mod() == nil {
		return 3, fmt.Sprintf("%s: no module loaded", EngineExport)
	}
	savePath := config.ModuleSavegamesPath()
	saveName := cmd.Args()[0]
	err := data.ExportGame(flame.Game(), savePath, saveName)
	if err != nil {
		return 3, fmt.Sprintf("%s: fail to export game: %v",
			EngineExport, err)
	}
	return 0, ""
}

// engineexportCharacter handles character option for engineexport.
func engineexportCharacter(cmd Command) (int, string) {
	if len(cmd.Args()) < 0 {
		return 3, fmt.Sprintf("%s: not enought args for: %s",
			EngineExport, cmd.OptionArgs()[0])
	}
	if flame.Game() == nil {
		return 3, fmt.Sprintf("%s: no game started", EngineExport)
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
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
		err := data.ExportCharacter(o, flame.Game().Module().Conf().CharactersPath())
		if err != nil {
			return 3, fmt.Sprintf("%s: %s#%s: fail to export: %v", EngineExport,
				o.ID(), o.Serial(), err)
		}
	}
	return 0, ""
}
