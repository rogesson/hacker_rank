""" find the runing-up """

def find_score(runners):
    """ find runner in a list """

    runners = list(set(runners))
    runners.sort()

    if len(runners) == 1:
        return runners[0]

    return runners[-2]

print find_score([2, 3, 6, 6, 5]) # 5
print find_score([1]) # 1
print find_score([57, 57, -57, 57])# -57
