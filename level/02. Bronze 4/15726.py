import math

a, b, c = map(int, input().split())
if a / b * c > a * b / c:
    print(math.floor(a / b * c))
else:
    print(math.floor(a * b / c))