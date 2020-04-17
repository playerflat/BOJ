a = int(input())
b = list(range(1, a + 1))

while len(b) != 1:
    print(b.pop(0), end=' ')
    b.append(b.pop(0))

print(b[0])
