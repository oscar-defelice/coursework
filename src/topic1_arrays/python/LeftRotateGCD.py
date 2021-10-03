#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'rotLeft' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts following parameters:
#  1. INTEGER_ARRAY a
#  2. INTEGER d
#


def rotLeft(a, d):
    for i in range(gcd(d, len(a))):
        tmp = a[i]
        j = i
        while True:
            k = j+d
            if k >= len(a):
                k -= n
            if k == i:
                break
            a[i] = a[k]
            j = k
        a[j] = tmp


def gcd(a, b):
    if b == 0:
        return a
    else:
        gcd(b, a % b)


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    d = int(first_multiple_input[1])

    a = list(map(int, input().rstrip().split()))

    result = rotLeft(a, d)

    fptr.write(' '.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
