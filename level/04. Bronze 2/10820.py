import sys

while 1:
    lo = u = n = b = 0
    a = sys.stdin.readline()
    if not a:
        break
    for i in a:
        if i.islower():
            lo += 1
        elif i.isupper():
            u += 1
        elif i.isnumeric():
            n += 1
        elif i == " ":
            b += 1
    print(lo, u, n, b)
