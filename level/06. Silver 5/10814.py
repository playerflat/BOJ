b = []
for i in range(int(input())):
    a = input().split()
    a = int(a[0]), a[1]
    b.append((a[0], i, a[1]))

b.sort()

for i in b:
    print(i[0], i[2])
