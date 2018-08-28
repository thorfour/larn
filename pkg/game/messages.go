package game

var welcome = `
	    Welcome to the game of Larn.  At this moment, you face a great problem.
	Your daughter has contracted a strange disease, and none of your home remedies
	seem to have any effect.  You sense that she is in mortal danger, and you must
	try to save her.  Time ago you heard of a land of great danger and opportunity.
	Perhaps here is the solution you need.

	    It has been said that there once was a great magician who called himself
	Polinneaus.  Many years ago, after having many miraculous successes, Polinneaus
	retired to the caverns of Larn, where he devoted most of his time to the
	creation of magic.   Rumors have it that one day Polinneaus set out to dispel
	an attacking army in a forest some distance to the north.  It is believed that
	here he met his demise.

	    The caverns of Larn, it is thought, must be magnificent in design,
	and contain much magic and treasure.  One option you have is to undertake a
	journey into these caverns.


	    Good Luck!  You're going to need it!
		

Press enter to continue`

var help = []string{`
	              Help File for The Caverns of Larn

	h  move to the left    H  run left          .  stay here
	j  move down           J  run down          Z  teleport yourself
	k  move up             K  run up            c  cast a spell
	l  move to the right   L  run right         r  read a scroll
	y  move northwest      Y  run northwest     q  quaff a potion
	u  move northeast      U  run northeast     W  wear armor
	b  move southwest      B  run southwest     T  take off armor
	n  move southeast      N  run southeast     w  wield a weapon
	^  identify a trap     g  give present pack weight  P  give tax status
	d  drop an item        i  inventory your pockets    Q  quit the game
	v  print program version   S  save the game         D  list all items found
	?  this help screen        A  create diagnostic file    e  eat something
					(wizards only)
	larn ++   restore checkpointed game
	larn -s   list the scoreboard
	larn -i   list scores with inventories
	larn -n   suppress welcome message when beginning a game
	larn -h   print out all the command line options
	larn -<number>      specify difficulty of the game (may be used with -n)
	larn -o<optsfile>   specify the .larnopts file to be used
	larn -c           create new scoreboards -- prompts for a password

         --- Press enter to exit, space for more help ---  
`, `
                    Background Information for Larn

When dropping gold, if you type '*' as your amount, all your gold gets dropped.
In general, typing in '*' means all of what your interested in.  This is true
when visiting the bank, or when contributing at altars.

Larn may need a VT100 to operate.  A check is made of the environment variable
"TERM" and it must be equal to "vt100".  This only applies if
the game has been compiled with "VT100" defined in the Makefile.  If compiled
to use termcap, there are no terminal restrictions, save needing cm, ce, & cl
termcap entries.

When in the store, trading post, school, or home, an <escape> will get you out.

larn -l           print out the larn log file

When casting a spell, if you need a list of spells you can cast, type 'D' as
the first letter of your spell.  The available list of spells will be shown,
after which you may enter the spell code.  This only works on the 1st letter
of the spell you are casting.

The Author of Larn is Noah Morgan (1982-3), Copying for Profit is Prohibited
Copyright 1986 by Noah Morgan, All Rights Reserved.

         --- Press enter to exit, space for more help ---  
`, `
             Explanation of the Larn scoreboard facility

    Larn supports TWO scoreboards, one for winners, and one for deceased
characters.  Each player (by userid or playerid, see UIDSCORE in Makefile)
is allowed one slot on each scoreboard, if the score is in the top ten for
that scoreboard.  This design helps insure that frequent players of Larn
do not hog the scoreboard, and gives more players a chance for glory.  Level
of difficulty is also noted on the scoreboards, and this takes precedence
over score for determining what entry is on the scoreboard.  For example:
if "Yar, the Bug Slayer" has a score of 128003 on the scoreboard at diff 0,
then his game at diff 1 and a score of 4112 would replace his previous
entry on the scoreboard.  Note that when a player dies, his inventory is
stored in the scoreboard so that everyone can see what items the player had
at the time of his death.

         --- Press enter to exit, space for more help ---  
`}
