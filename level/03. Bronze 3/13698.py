a = [-1, 0, 0, 1]

for i in input():
    if i == "A":
        a[0], a[1] = a[1], a[0]
    elif i == "B":
        a[0], a[2] = a[2], a[0]
    elif i == "C":
        a[0], a[3] = a[3], a[0]
    elif i == "D":
        a[1], a[2] = a[2], a[1]
    elif i == "E":
        a[1], a[3] = a[3], a[1]
    else:
        a[2], a[3] = a[3], a[2]

print(a.index(-1)+1)
print(a.index(1)+1)
