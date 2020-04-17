a, b = map(int, input().split())
y = a % abs(b)
print((a - y) // b)
print(y)