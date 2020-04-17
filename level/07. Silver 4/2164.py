a = int(input())
b = {}
c = 1
d = a + 1
for i in range(1, a + 1):
    b[i] = i

while len(b) != 1:
    del b[c]
    c += 1
    b[d] = b.pop(c)
    c += 1
    d += 1

print(b.popitem()[1])
