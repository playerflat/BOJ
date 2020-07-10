a = input()
b = input()

for i in range(100000):
    if a[i] + b[i] == '11':
        print(1, end='')
    else:
        print(0, end='')
print()
for i in range(100000):
    if a[i] + b[i] != '00':
        print(1, end='')
    else:
        print(0, end='')
print()
for i in range(100000):
    if a[i] != b[i]:
        print(1, end='')
    else:
        print(0, end='')
print()
for i in range(100000):
    if a[i] == '0':
        print(1, end='')
    else:
        print(0, end='')
print()
for i in range(100000):
    if b[i] == '0':
        print(1, end='')
    else:
        print(0, end='')