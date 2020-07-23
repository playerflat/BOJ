ans = 0
for i in range(10):
    mush = int(input())
    if ans + mush >= 100:
        if 100 - ans < ans + mush - 100:
            print(ans)
            break
        else:
            print(ans + mush)
            break
    ans += mush
    if i == 9:
        print(ans)