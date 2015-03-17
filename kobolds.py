__author__ = 'whillegass'

import random


def makeKobold():

  brawn = rollDice(2)

  ego = rollDice(2)

  extraneous = rollDice(2)

  reflexes = rollDice(2)

  hits = brawn

  meat = getHandyNumber(brawn)

  cunning = getHandyNumber(ego)

  luck = getHandyNumber(extraneous)

  agility = getHandyNumber(reflexes)

  skills = []

  chooseSkills(skills, ego)

  edges = ['Bark like a Kobold', 'Kobold senses', getEdge()]

  bogies = ['Fearless', 'Tastes Like Chicken', getBogie()]

  armor = getArmor()

  gear = getGear()

  weapon = getWeapon()

  print 'Brawn ',  brawn , ' Meat ' , meat

  print 'Ego ', ego , ' cunning' , cunning

  print 'Extraneous ', extraneous , ' Luck ' , luck

  print 'Reflexes ', reflexes , ' Agility ' , agility

  print 'Skills: '

  print skills

  print "** If you do not have cook skill, start with 1 death check **"

  print 'edges: '

  print edges

  print 'Bogies:'

  print bogies

  print 'Armor ' + armor

  print 'Left Paw ' + gear

  print 'Right paw ' + weapon

  print 'Hits: ' , hits

  print ''

  print ''

  print ''

  print ''

  print ''



def getWeapon():

  roll = rollDice(2)

  gear = {2 :'Iron Skillet (3 dam, +cook)',

          3 :'Large Bone (3 dam, -big)',

          4 :'Chain (2 dam, +climb)',

          5 :'club (2 dam, +bash)',

          6 :'small sword (2 dam, looks wimpy)',

          7 :'Dagger (1 dam, looks cool)',

          8 :'Hammer (1 dam, +useful)',

          9 :'slingshot(1 dam, + stones)',

          11:'dead rat (0dam, +foul smelling',

          10:'Cooking Utensil (1 dam, + cook)',

          12:'Squat'}

  return gear[roll]





def getGear():

  roll = rollDice(2)

  gear = {2 :'Cup of milk elemental summoning',

          3 :'Bag of holding: chickens',

          4 :'Ring of human speaking',

          5 :'Codex of Tabriz the Arcane',

          6 :'Random Booze',

          7 :'Spice Sack',

          8 :'backpack',

          9 :"Adventurer's cast-offs",

          10:'25 feet of rope',

          11:'10 foot pole',

          12:'lint'}

  return gear[roll]



def getArmor():

  roll = rollDice(2)

  armor = {2 :'big shield (-Big, 10 hits, 2 dam)',

           3 :'Metal Pot Helm (10 hits, +cook)',

           4 :'Beer-Barrel (-Bulky 12 hits)',

           5 :'leather apron (+backpack 8 hits)',

           6 :'leather jacket (+Fonzie 8 hits)',

           7 :'small shield (-item, 1dam, 6 hits)',

           8 :'heavy t-shirt (4 hits)',

           9 :'kids clothes (2 hits)',

           10:'socks (1 hit)',

           11:'lintmail vest (1 hit)',

           12:'nekked'}

  return armor[roll]



def getBogie():

 roll = int(random.randint(1,6))

 skills = {1 : 'Angry Friends',

           2 : 'Animal Foe',

           3 : 'Foul Smelling',

           4 : 'Hungry',

           5 : 'In Heat',

           6 : 'Tastes Like Baby'}

 return skills[roll]



def getEdge():

 roll = int(random.randint(1,6))

 skills = {1 : 'Animal Chum',

           2 : 'Bouncy',

           3 : 'Extra Padding',

           4 : 'Troll Blood',

           5 : 'Winning Smile',

           6 : 'Zilch'}

 return skills[roll]



def rollDice(count):

  total = 0;

  for i in range(0, count):

    total += int(random.randint(1,6))

  return total



def getHandyNumber(stat):

  if stat <  5:

    return 1

  if stat < 9:

    return 2

  if stat < 13:

    return 3

  if stat < 17:

    return 4

  if stat < 21:

    return 5



def chooseSkills(skills, ego):

  numskills = 6

  if (ego < 6):

    numskills = ego

  skills.append(getBrawnSkill())

  skills.append(getEgoSkill())

  skills.append(getExtraneousSkill())

  skills.append(getReflexesSkill())

  if (numskills > 4):

    skills.append('cook')







def getBrawnSkill():

  roll = int(random.randint(1,6))

  skills = {1 : 'Athlete',

            2 : 'Bully',

            3 : 'Duel',

            4 : 'Lift',

            5 : 'Swim',

            6 : 'Wrassle'}

  return skills[roll]



def getEgoSkill():

  roll = int(random.randint(1,6))

  skills = {1 : 'Hide',

            2 : 'Lackey',

            3 : 'Sage',

            4 : 'Shoot',

            5 : 'Speak human',

            6 : 'trap'}

  return skills[roll]



def getExtraneousSkill():

  roll = int(random.randint(1,4))

  skills = {1 : 'Bard',

            2 : 'Dungeon',

            3 : 'Track',

            4 : 'Trade'}

  return skills[roll]



def getReflexesSkill():

  roll = int(random.randint(1,6))

  skills = {1 : 'Cower',

            2 : 'Fast',

            3 : 'Sneak',

            4 : 'Steal',

            5 : 'Ride',

            6 : 'Wiggle'}

  return skills[roll]



def main():

  for i in range (0, 100):

    makeKobold()



if __name__ == "__main__":

    main()