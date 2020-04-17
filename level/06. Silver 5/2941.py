a = input()
cro = ["c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="]

for i in cro:
    a = a.replace(i, "0")

print(len(a))
