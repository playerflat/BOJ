a = [4, 7]

for i in a:
    t = i * 10 + 4
    a.append(t)
    if t > 1000000:
        break
    t = i * 10 + 7
    a.append(t)

N = int(input())
for i in range(len(a)):
    if N == a[i]:
        print(a[i])
        break
    elif N < a[i]:
        print(a[i - 1])
        break
