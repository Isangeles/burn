/*
 * engineload.go
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
  "path/filepath"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/data"
	flamecfg "github.com/isangeles/flame/config"
)

// engineload handles engineload command.
func engineload(cmd Command) (int, string) {
	if len(cmd.OptionArgs()) < 1 {
		return 2, fmt.Sprintf("%s: no_option_args", EngineLoad)
	}
	switch cmd.OptionArgs()[0] {
  case "module":
    return engineloadModule(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", EngineLoad,
			cmd.OptionArgs()[0])
	}
}

// engineloadModule handles module option for engineload.
func engineloadModule(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			EngineLoad, cmd.OptionArgs()[0])
	}
	modPath := filepath.FromSlash("data/modules/" + cmd.Args()[0])
	m, err := data.Module(modPath, flamecfg.LangID())
	if err != nil {
		return 3, fmt.Sprintf("%s: module load fail: %s",
			EngineLoad, err)
	}
	flame.SetModule(m)
	return 0, ""
}
