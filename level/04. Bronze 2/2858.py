r, b = map(int, input().split())
for i in range(3, (r + b) // 2 + 1):
    j = (r + b) // i
    if i * j == r + b and r == 2 * i + 2 * j - 4:
        print(max(i, j), min(i, j))
        break
