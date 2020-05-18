import math
import sys

while 1:
    c = 0
    a = int(sys.stdin.readline())
    if not a:
        break
    for i in range(len(str(a))):
        c += a % 10 * math.factorial(i + 1)
        a //= 10
    print(c)
