r = [1, 0, 0]
for i in input():
    if i == "A":
        r[0], r[1] = r[1], r[0]
    elif i == "B":
        r[1], r[2] = r[2], r[1]
    else:
        r[0], r[2] = r[2], r[0]
print(r.index(1)+1)
