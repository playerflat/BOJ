d = {}
s = []
for i in range(1, 9):
    t = int(input())
    s.append(t)
    d[t] = i

s.sort(reverse=True)

print(sum(s[:5]))

r = []
for i in s[:5]:
    r.append(d[i])

r.sort()
print(' '.join(map(str, r)))
