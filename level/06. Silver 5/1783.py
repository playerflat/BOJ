N, M = map(int, input().split())

if N == 1 or M == 1:
    print(1)
elif N == 2:
    print(min(4, (M+1) // 2))
else:
    if M <= 4:
        print(M)
    elif M <= 6:
        print(4)
    else:
        print(M - 2)
