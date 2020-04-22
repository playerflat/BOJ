import sys

a = []
b = float(sys.stdin.readline())
while 1:
    c = float(sys.stdin.readline())
    if c == 999:
        break
    else:
        a.append(c)

for i in range(len(a)):
    print(f'{a[i] - b:.2f}')
    b = a[i]
