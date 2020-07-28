N, K = map(int, input().split())
medal = {}
rank = 1
for _ in range(N):
    a, b, c, d = map(int, input().split())
    medal[a] = [b, c, d]

target = medal[K]
for _, v in medal.items():
    if v[0] > target[0]:
        rank += 1
    elif v[0] == target[0]:
        if v[1] > target[1]:
            rank += 1
        elif v[1] == target[1]:
            if v[2] > target[2]:
                rank += 1
print(rank)
