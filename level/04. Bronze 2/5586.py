a = input()
joi = ioi = 0
for i in range(0, len(a)):
    if a[i:i + 3] == 'JOI':
        joi += 1
    elif a[i:i + 3] == 'IOI':
        ioi += 1
print(joi)
print(ioi)