min = 0
x, y = map(int, input().split())
min = 1000 / y * x
for _ in range(int(input())):
    x, y = map(int, input().split())
    if min > 1000 / y * x:
        min = 1000 / y * x

print(f'{min:.2f}')
