for _ in range(int(input())):
    a, b = map(int, input().split())
    print(int(a ** 2 / b ** 2 + (1 if a % b else 0)))
