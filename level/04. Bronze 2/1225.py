a, b = input().split()
r1 = r2 = 0
for i in a:
    r1 += int(i)
for i in b:
    r2 += int(i)
print(r1 * r2)
