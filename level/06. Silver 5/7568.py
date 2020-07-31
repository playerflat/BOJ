s = {}
N = int(input())
for i in range(N):
    x, y = map(int, input().split())
    s[i] = [x, y, 1]

for i in range(N):
    for j in range(i, N):
        if s[i][0] < s[j][0] and s[i][1] < s[j][1]:
            s[i][2] += 1
        elif s[i][0] > s[j][0] and s[i][1] > s[j][1]:
            s[j][2] += 1

for v in s.values():
    print(v[2], end=' ')
