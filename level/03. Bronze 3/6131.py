c = 0
x = int(input())
for a in range(1, 501):
    for b in range(1, 501):
        if a ** 2 == b ** 2 + x:
            c += 1
print(c)
