# PyPy3

for _ in range(int(input())):
    n, m = map(int, input().split())
    r = 0
    for b in range(1, n):
        for a in range(1, b):
            c = (a ** 2 + b ** 2 + m) / (a * b)
            if int(c) == c:
                r += 1
    print(r)