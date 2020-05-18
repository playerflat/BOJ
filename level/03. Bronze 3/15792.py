a, b = map(int, input().split())
c = a % b
print(f'{a // b}.', end='')
for _ in range(1001):
    c *= 10
    d = c // b
    print(d, end='')
    c %= b
