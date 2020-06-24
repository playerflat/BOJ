a = int(input())
b = ''
r = 0
s = 0
d = a // 100000

if d > 0:
    r += d * 150000
    if (d - 1) // 3 > 0:
        r += (d - 1) // 3 * 100000

for i in range(d*100000, a + 1):
    b += str(i)

for i in range(len(b)):
    if b[i] in ['3', '6', '9']:
        r += 1
print(r)