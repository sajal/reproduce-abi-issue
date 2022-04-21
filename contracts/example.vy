# @version 0.3.2

struct Numbers:
    a: uint256
    b: uint256

NUM_SIZE: constant(uint256) = 5

@external
def __init__():
    pass

@pure
@external
def doMath(items: Numbers[NUM_SIZE]) -> uint256:
    res: uint256 = 0
    for i in range(NUM_SIZE):
        res = res + (items[i].a * items[i].b)
    return res
