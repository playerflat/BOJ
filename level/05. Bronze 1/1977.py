m = int(input())
n = int(input())
a = []
for i in range(int(m ** 0.5), int(n ** 0.5) + 1):
    if m <= i ** 2 <= n:
        a.append(i ** 2)
if len(a) >= 1:
    print(sum(a))
    print(min(a))
else:
    print(-1)
