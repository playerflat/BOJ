import sys

r = [0, 0, 0, 0, 0]
for _ in range(int(sys.stdin.readline())):
    a, b = map(int, sys.stdin.readline().split())
    if a > 0 and b > 0:
        r[0] += 1
    elif a < 0 < b:
        r[1] += 1
    elif a < 0 and b < 0:
        r[2] += 1
    elif a > 0 > b:
        r[3] += 1
    elif a == 0 or b == 0:
        r[4] += 1

print(f'Q1: {r[0]}')
print(f'Q2: {r[1]}')
print(f'Q3: {r[2]}')
print(f'Q4: {r[3]}')
print(f'AXIS: {r[4]}')
