x, y = 1, 1
for i in range(1, int(input())):
    if i % 2 != 0:
        y += x
    else:
        x += y
print(x * 2 + y * 2)
