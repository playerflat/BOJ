a, b = input().split()
r = [['.'] * len(a) for i in range(len(b))]
f = 0

for i in range(len(a)):
    for j in range(len(b)):
        if a[i] == b[j]:
            for l in range(len(b)):
                r[l][i] = b[l]
            for k in range(len(a)):
                r[j][k] = a[k]
            f = 1
            break
    if f:
        break
for i in r:
    print(''.join(i))
