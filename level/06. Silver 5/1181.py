a = set()
b = []
for i in range(int(input())):
    a.add(input())
for i in a:
    b.append((len(i), i))
b.sort()
for i in b:
    print(i[1])
