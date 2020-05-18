p = int(input())
c = [500, 100, 50, 10, 5, 1]
a = 1000 - p
r = 0
for i in c:
    r += a // i
    a -= (a // i) * i

print(r)