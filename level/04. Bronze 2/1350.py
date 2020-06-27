import math

input()
a = list(map(int, input().split()))
b = int(input())
r = 0
for i in a:
    if i == 0:
        continue
    else:
        r += math.ceil(i / b) * b
print(r)
