a, b = 0, 1
for _ in range(1, int(input())):
    a, b = b, a+b
print(a, b)