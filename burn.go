/*
 * burn.go
 *
 * Copyright 2018-2019 Dariusz Sikora <dev@isangeles.pl>
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

// Burn is Flame engine commands interpreter.
package burn

import (
	"fmt"
	"strings"
)

// Command interface for all commands interpreted by CI.
type Command interface {
	Tool() string
	String() string
	TargetArgs() []string
	OptionArgs() []string
	Args() []string
	AddArgs(args ...string)
	AddTargetArgs(args ...string)
}

// Interaface for all expressions interpreted by CI.
// Expressions are multiple commands connected by
// special character(e.g. pipe).
type Expression interface {
	Commands() []Command
	Type() ExpressionType
	String() string
}

// Type for expression type.
type ExpressionType int

const (
	// CI tools names.
	EngineShow   = "engineshow"
	EngineImport = "engineimport"
	EngineExport = "engineexport"
	EngineSet    = "engineset"
	ModuleAdd    = "moduleadd"
	ModuleShow   = "moduleshow"
	ChapterShow  = "chaptershow"
	GameAdd      = "gameadd"
	GameRemove   = "gameremove"
	ObjectAdd    = "objectadd"
	ObjectRemove = "objectremove"
	ObjectSet    = "objectset"
	ObjectShow   = "objectshow"
	ObjectHave   = "objecthave"
	ObjectUse    = "objectuse"
	// Syntax.
	IDSerialSep = "#"
	// Expr types.
	PipeArgExp ExpressionType = iota
	PipeTarExp
	SeqExp
	NoExp
)

var (
	tools map[string]func(cmd Command) (int, string)
)

// On init.
func init() {
	tools = make(map[string]func(cmd Command) (int, string), 0)
	tools[EngineShow] = engineshow
	tools[EngineImport] = engineimport
	tools[EngineExport] = engineexport
	tools[EngineSet] = engineset
	tools[ModuleAdd] = moduleadd
	tools[ModuleShow] = moduleshow
	tools[ChapterShow] = chaptershow
	tools[GameAdd] = gameadd
	tools[GameRemove] = gameremove
	tools[ObjectAdd] = objectadd
	tools[ObjectRemove] = objectremove
	tools[ObjectSet] = objectset
	tools[ObjectShow] = objectshow
	tools[ObjectHave] = objecthave
	tools[ObjectUse] = objectuse
}

// AddToolHandler adds specified command handling function as
// CI tool with specified name.
func AddToolHandler(name string, handler func(cmd Command) (int, string)) {
	tools[name] = handler
}

// HandleCommand handles specified command,
// returns response code and message.
func HandleCommand(cmd Command) (int, string) {
	for tool, handleFunc := range tools {
		if cmd.Tool() == tool {
			return handleFunc(cmd)
		}
	}
	return 1, fmt.Sprintf("no_such_ci_tool_found:'%s'",
		cmd.Tool())
}

// HandleExpression handles specified expression,
// returns response code and massage.
func HandleExpression(exp Expression) (int, string) {
	switch exp.Type() {
	case PipeArgExp:
		return HandleArgsPipe(exp.Commands()...)
	case PipeTarExp:
		return HandleTargetArgsPipe(exp.Commands()...)
	case NoExp:
		return HandleCommand(exp.Commands()[0])
	default:
		return 1, fmt.Sprintf("unknow expression type: %d", exp.Type())
	}
}

// HandleArgsPipe handles specified commands
// connected with pipe('|').
// Pipe pushes out from command on the left to
// command on right as arguments.
func HandleArgsPipe(cmds ...Command) (res int, out string) {
	for _, cmd := range cmds {
		res, out = pipeArgs(cmd, out)
		if res != 0 {
			return res, out
		}
	}
	return
}

// HandleTargetArgsPipe handles specified commands
// connected with pipe('|').
// Pipe pushes out from command on the left to
// command on right as target arguments.
func HandleTargetArgsPipe(cmds ...Command) (res int, out string) {
	for _, cmd := range cmds {
		res, out = pipeTargetArgs(cmd, out)
		if res != 0 {
			return res, out
		}
	}
	return
}

// pipeArgs pushes specified text(out from previous command)
// to specified command as arguments, and executes
// specified command.
func pipeArgs(cmd Command, out string) (int, string) {
	args := strings.Fields(strings.TrimSpace(out))
	cmd.AddArgs(args...)
	return HandleCommand(cmd)
}

// pipeTargetArgs pushes specified text(out from previous command)
// to specified command as target arguments, and executes
// specified command.
func pipeTargetArgs(cmd Command, out string) (int, string) {
	args := strings.Fields(strings.TrimSpace(out))
	cmd.AddTargetArgs(args...)
	return HandleCommand(cmd)
}

// argSerialID parses specified command arg to
// game object ID and serial value.
// Argument format: [id]IDSerialSep[serial].
// In case of error returns empty serial.
func argSerialID(arg string) (string, string) {
	serialid := strings.Split(arg, IDSerialSep)
	if len(serialid) < 2 {
		return arg, ""
	}
	return serialid[0], serialid[1]
}
