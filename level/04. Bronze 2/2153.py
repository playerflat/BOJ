d = {}
r = 0
a = 'It is a prime word.'
for i in range(1, 27):
    d[chr(i + 96)] = i
    d[chr(i + 64)] = i + 26

for i in input():
    r += d[i]

for i in range(2, r):
    if r % i == 0:
        a = 'It is not a prime word.'
        break

print(a)
