## Introduction
  Burn is command interpreter for the Flame engine.

## Syntax
  Standard Burn command syntax:
```
  [tool name] -o [option] -t [targets...] -a [arguments...]
```
  Beside many arguments Burn handles also many targets, e.g. one charman tool command can be executed on many game characters, which results with a combined output.

  Example command:
```
  >objectshow -o position -t player_test_0 player_test_1
```
  Example output:
```
  0.0x0.0 420.0x100.0
```
  Shows positions of game objects with serial IDs 'player_test_0' and 'player_test_1'.

  Commands can also be joined into expressions.

  Target pipe expression:
```
  [command] |t [command]
```
  Executes the first command and uses the output as target arguments to execute next command.

  Example expression:
```
  >moduleshow -o area-chars -t area1_test |t objectshow -o position
```
  Example output:
```
  0.0x0.0 12.0x131.0 130.0x201.0
```
  Shows positions of all game characters in the area with ID 'area1_test'.

## Ash
[Ash](https://github.com/Isangeles/burn/tree/master/ash) is scripting language that allows running Burn commands under conditional loop.

## Commands
Set target:
```
  $charman -o set -t [ID]_[serial] -a target [ID]_[serial]
```
Description: sets object with specified serial ID(-a) as target of character with specified serial ID(-t).

Export game character:
```
  $charman -o export -t [character ID]
```
Description: exports game character with specified ID to XML file in
data/modules/[module]/characters directory.

Load module:
```
  $engineload -t module -a [module name] [module path](optional)
```
Description: loads module with the specified name(module directory name) and with a specified path,
if no path provided, the engine will search default modules directory(data/modules).

Save game:
```
  $enginesave -t game -a [save file name]
```
Description: saves current game to 'savegames/[module]' directory.

Add item:
```
  $charadd -o item -a item [item ID] -t [character ID]#[character serial]
```
Description: adds item with specified ID to inventory of game character with specified serial ID.

Equip item:
```
  $charman -o equip -t [character serial ID] -a [slot ID] [item serial ID]
```
Description: equips item with specified ID for game character with specified serial ID.

Add effect:
```
  $charman -o add -t [character serial ID] -a effect [effect ID]
```
Description: puts effect with specified ID on game character with specified serial ID

Spawn NPC:
```
  $moduleadd -t character -a [character ID] [scenario ID] [areaID] [posX](optional) [posY](optional)
```
Description: spawns new chapter NPC with specified ID in specified scenario area at given position(0, 0 if not specified).

## Contributing
You are welcome to contribute to project development.

If you looking for things to do, then check TODO file.

When you finish, open pull request to merge your changes with main branch.

## Contact
* Isangeles <<dev@isangeles.pl>>

## License
Copyright 2018-2019 Dariusz Sikora <<dev@isangeles.pl>>

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
MA 02110-1301, USA.
