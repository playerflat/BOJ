e, f, c = map(int, input().split())
r = (e + f) // c
a = (e + f) % c + r
while 1:
    if a < c:
        break
    r += a // c
    a = (a // c) + (a % c)

print(r)
