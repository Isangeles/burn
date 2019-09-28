/*
 * engineimport.go
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
	flameconf "github.com/isangeles/flame/config"
	"github.com/isangeles/flame/core/data"
)

// engineimport handles engineload command.
func engineimport(cmd Command) (int, string) {
	if len(cmd.OptionArgs()) < 1 {
		return 2, fmt.Sprintf("%s: no option args", EngineImport)
	}
	switch cmd.OptionArgs()[0] {
	case "module":
		return engineimportModule(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", EngineImport,
			cmd.OptionArgs()[0])
	}
}

// engineimportModule handles module option for engineload.
func engineimportModule(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			EngineImport, cmd.OptionArgs()[0])
	}
	modPath := filepath.FromSlash("data/modules/" + cmd.Args()[0])
	m, err := data.Module(modPath, flameconf.LangID())
	if err != nil {
		return 3, fmt.Sprintf("%s: module load fail: %s",
			EngineImport, err)
	}
	flame.SetModule(m)
	return 0, ""
}
