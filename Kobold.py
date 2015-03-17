__author__ = 'wendy'

class Kobold:
    'Class to define kobolds'

    def __init__(self, brawn, ego, extraneous, reflexes, skills, edges, bogies, armor, gear, weapon):
        self.brawn = brawn
        self.ego = ego
        self.extraneous = extraneous
        self.reflexes = reflexes
        self.skills = skills
        self.edges = edges
        self.bogies = bogies
        self.armor = armor
        self.gear = gear
        self.weapon = weapon


    def __getHandyNumber(stat):
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
