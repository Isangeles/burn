## Introduction
Ash is a scripting language that executes Burn commands in a conditional loop.

This allows creating cutscenes and special events using the game world objects.

## Examples
Infinite loop script:
```
# Every 5 seconds, sends text from first argument on chat
# channel of object with serial ID from second argument
true {
    objectset -o chat -a @1 -t @2;
    wait(5);
};
```

Script with inner block:
```
# Script that sets position of character with serial ID from arg 1
# to 0x0 if range between him and character with serial ID from arg 2
# is less than 50.
true {
    rawdis(@1, @2) < 50 {
      objectset -o position -a 0 0 -t @1;
      wait(5);
    };
};
```

Declaring arguments inside script:
```
# Script that sets position of character with serial ID from arg 1
# to 0x0 if range between him and character with serial ID from arg 2
# is less than 50.
@1 = char#1
@2 = char#3
true {
    rawdis(@1, @2) < 50 {
      objectset -o position -a 0x0 -t @1;
      wait(5);
    };
};
```

For loop:
```
# Script that sends arg2 text on chat channel of character with
# arg 1 serial ID if raw distance between him and any character from
# area with ID 'area1_main' is less than 50, after that script halts 
# for 5 secs.
@1 = testchar#0
@2 = "hay you!"
true {
	for(@3 = out(moduleshow -o area-chars -t area1_main)) {
     		rawdis(@1, @3) < 50 {
        		objectset -o char -a @2 -t @1;
        		wait(5);
	     	};
	};
};
```

Output comparison:
```
# Script that sends arg2 text on chat channel of character with
# arg1 serial ID if raw distance between him and any character from
# area with ID 'area1_main'(expect char with arg1 serial ID) is less 
# than 50, after that script halts for 5 secs.
@1 = testchar#0
@2 = "hay you!"
true {
	for(@3 = out(moduleshow -o area-chars -t area1_main)) {
		@1 != @3 {
	     		rawdis(@1, @3) < 50 {
				objectset -o chat -a @2 -t @1;
				wait(5);
		     	};
		};
	};
};
```

End macro:
```
# Spawns character with arg1 ID in scenario with arg2 ID and area
# with arg3 ID then ends script.
@1 = testchar
@2 = scenario
@3 = scenario_area
{
	moduleadd -o char -a @1 @2 @3
	end();
}
```
