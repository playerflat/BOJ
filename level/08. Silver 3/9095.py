n = int(input())

for i in range(n):
    a, b, c = 1, 0, 0
    x = int(input())
    for j in range(x):
        a, b, c = a + b + c, a, b
    print(a)