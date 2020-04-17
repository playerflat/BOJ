a, b = input().split()

a = int(a)
b = int(b)
# a시 b분 - 45분
if b < 45:
    a -= 1
    b = 60-45+b
else:
    b -= 45

if a < 0:
    a = 24 + a
print(a, b)