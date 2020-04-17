a = int(input())
b = int(input())
c = int(input())
d = int(input())
p = int(input())

ap = a * p
if c > p:
    bp = b
else:
    bp = b + (p - c) * d

if ap>bp:
    print(bp)
else:
    print(ap)