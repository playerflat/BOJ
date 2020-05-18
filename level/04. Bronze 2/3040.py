a = []
for _ in range(9):
    a.append(int(input()))

b = sum(a) - 100
f = False
for i in range(9):
    for j in range(i + 1, 9):
        if a[i] + a[j] == b:
            a.remove(a[j])
            a.remove(a[i])
            f = True
            break
    if f:
        break

for i in a:
    print(i)
