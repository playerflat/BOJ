def convert(a):
    z = "012345678"
    x, y = divmod(a, 9)
    if x == 0:
        return z[y]
    else:
        return convert(x) + z[y]


a = int(input())
print(convert(a))
