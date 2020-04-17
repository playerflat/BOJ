import math

a = int(input())
r = a
for i in range(2, int(math.sqrt(a)) + 1):
    for j in range(i, a):
        if i * j <= a:
            r += 1
        else:
            break
print(r)
