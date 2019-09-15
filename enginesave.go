/*
 * enginesave.go
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
)

// enginesave handles enginesave command.
func enginesave(cmd Command) (int, string) {
	if len(cmd.OptionArgs()) < 1 {
		return 2, fmt.Sprintf("%s: no_option_args", EngineSave)
	}
	switch cmd.OptionArgs()[0] {
	case "game":
		return enginesaveGame(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", EngineSave,
			cmd.OptionArgs()[0])
	}
}

// enginesaveGame handles game option for enginesave.
func enginesaveGame(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: not enought args for: %s",
			EngineSave, cmd.OptionArgs()[0])
	}
	if flame.Mod() == nil {
		return 3, fmt.Sprintf("%s: no module loaded", EngineSave)
	}
	savePath := config.ModuleSavegamesPath()
	saveName := cmd.Args()[0]
	err := data.SaveGame(flame.Game(), savePath, saveName)
	if err != nil {
		return 3, fmt.Sprintf("%s: fail to save game: %v",
			EngineSave, err)
	}
	return 0, ""
}
