/*
 * moduleshow.go
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
	"strings"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/data/res"
	"github.com/isangeles/flame/core/data/text/lang"
	"github.com/isangeles/flame/core/module/scenario"
)

// moduleshow handles moduleshow command.
func moduleshow(cmd Command) (int, string) {
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ModuleShow)
	}
	switch cmd.OptionArgs()[0] {
	case "id":
		return 0, flame.Mod().Conf().ID
	case "res-objects":
		out := ""
		for _, obd := range res.Objects() {
			out = fmt.Sprintf("%s%s ", out, obd.BasicData.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "res-dialogs":
		out := ""
		for _, dl := range res.Dialogs() {
			out = fmt.Sprintf("%s%s ", out, dl.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "res-quests":
		out := ""
		for _, qr := range res.Quests() {
			out = fmt.Sprintf("%s%s ", out, qr.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "res-miscs":
		out := ""
		for _, md := range res.MiscItems() {
			out = fmt.Sprintf("%s%s ", out, md.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "res-recipes":
		out := ""
		for _, rd := range res.Recipes() {
			out = fmt.Sprintf("%s%s ", out, rd.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "area-chars":
		return moduleshowAreaChars(cmd)
	case "lang":
		return moduleshowLang(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", ModuleShow,
			cmd.OptionArgs()[0])
	}
}

// moduleshowAreaChars handles area-chars option for moduleshow.
func moduleshowAreaChars(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 3, fmt.Sprintf("%s: no game loaded", ModuleShow)
	}
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ModuleShow, cmd.OptionArgs()[0])
	}
	areaID := cmd.TargetArgs()[0]
	var area *scenario.Area
	for _, s := range flame.Mod().Chapter().Scenarios() {
		for _, a := range s.Areas() {
			if a.ID() == areaID {
				area = a
			}
		}
	}
	if area == nil {
		return 3, fmt.Sprintf("%s: area not found: %s",
			ModuleShow, areaID)
	}
	out := ""
	for _, c := range area.Characters() {
		out = fmt.Sprintf("%s%s%s%s ", out, c.ID(),
			IDSerialSep, c.Serial())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// moduleshowAreaObjects handles area-objects
// option for moduleshow.
func moduleshowAreaObjects(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 3, fmt.Sprintf("%s: no game loaded", ModuleShow)
	}
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ModuleShow, cmd.Args()[0])
	}
	areaID := cmd.TargetArgs()[0]
	var area *scenario.Area
	for _, s := range flame.Mod().Chapter().Scenarios() {
		for _, a := range s.Areas() {
			if a.ID() == areaID {
				area = a
			}
		}
	}
	if area == nil {
		return 3, fmt.Sprintf("%s: area not found: %s",
			ModuleShow, areaID)
	}
	out := ""
	for _, o := range area.Objects() {
		out = fmt.Sprintf("%s%s%s%s ", out, o.ID(),
			IDSerialSep, o.Serial())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// moduleshowLang handles lang optiond for moduleshow.
func moduleshowLang(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no args", ModuleShow)
	}
	id := cmd.Args()[0]
	return 0, lang.TextDir(flame.Mod().Conf().LangPath(), id)
}
