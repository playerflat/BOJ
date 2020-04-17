a = int(input())
b = map(int, input().split())
l = []
i = 1

for keys in b:
    if keys==0:
        l.append(i)
    else:
        l.insert(-keys, i)
    i += 1

for i in range(len(l)):
    print(l[i], end=' ')
