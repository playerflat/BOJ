a = int(input())
b = list(map(int, input().split()))
c = []
for i in range(1, a + 1):
    c.append(b[i - 1] * i - sum(c[:i]))

for i in c:
    print(i, end=' ')