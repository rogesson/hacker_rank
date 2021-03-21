""" List Comprehensions """

def permutations(n, z):
    """ Get all possible forms """

        print [[a, b, c] for a in range(0, n[0] + 1) for b in range(0, n[1] +1) for c in range(0, n[2] + 1) if a + b + c != z]

if __name__ == '__main__':
    permutations([1, 1, 1], 2)

