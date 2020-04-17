d = {}
for i in range(97, 123):
    d[chr(i)] = -1

a = input()

for i in range(len(a) - 1, -1, -1):
    d[a[i]] = i

for i in d.keys():
    print(d[i], end='')
    if i != "z":
        print(" ", end='')
