x, y = input().split()

a = int(x)
b = int(y)

if a < b:
    a, b = b, a
while b != 0:
    a, b = b, a % b
print(a)

a = int(int(x)/a)
b = int(y)
print(b*a)
