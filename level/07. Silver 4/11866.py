a, b = map(int, input().split())
c = list(range(1, a + 1))
i = -1
print("<", end='')

while True:
    i += b
    if i >= len(c):
        i %= len(c)
    print(c.pop(i), end='')
    i -= 1
    if len(c) == 0:
        print(">")
        break
    else:
        print(", ", end='')
