import math

a = int(input())
b = {}
while True:
    c = a % 10
    if c == 9:
        c = 6
    if b.get(c) is None:
        b[c] = 1
    else:
        b[c] += 1
    a = int(a / 10)
    if a == 0:
        break

if b.get(6) is not None:
    b[6] = math.ceil(b[6] / 2)
print(max(b.values()))