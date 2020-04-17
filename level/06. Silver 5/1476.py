E, S, M = map(int, input().split())
i = 1
while True:
    ec = i % 15
    sc = i % 28
    mc = i % 19
    if ec == 0:
        ec = 15
    if sc == 0:
        sc = 28
    if mc == 0:
        mc = 19
    if ec == E and sc == S and mc == M:
        print(i)
        break
    else:
        i += 1