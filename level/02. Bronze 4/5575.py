t1 = list(map(int, input().split()))
t2 = list(map(int, input().split()))
t3 = list(map(int, input().split()))

for i in t1, t2, t3:
    diff = (i[3] - i[0]) * 3600 + (i[4] - i[1]) * 60 + i[5] - i[2]

    print(int(diff / 3600), int(diff % 3600 / 60), int(diff % 3600 % 60))
