import collections

d = {}
for i in range(97, 123):
    d[chr(i)] = 0
m = 0

try:
    while 1:
        a = input().replace(' ', '')
        for i in a:
            d[i] += 1
            if d[i] > m:
                m = d[i]
except EOFError:
    for i, v in d.items():
        if v == m:
            print(i, end='')
