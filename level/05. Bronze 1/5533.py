d = []
N = int(input())
player = [0 for _ in range(N)]
for _ in range(N):
    d.append(list(map(int, input().split())))

for i in range(3):
    game = []
    for j in range(N):
        game.append(d[j][i])
    for j in range(N):
        if game.count(d[j][i]) < 2:
            player[j] += d[j][i]

for i in range(N):
    print(player[i])
