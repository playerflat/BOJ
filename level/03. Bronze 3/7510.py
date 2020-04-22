for i in range(1, int(input()) + 1):
    a = list(map(int, input().split()))
    print(f'Scenario #{i}:')
    if max(a) ** 2 == min(a) ** 2 + (sum(a)-min(a)-max(a)) ** 2:
        print("yes")
    else:
        print("no")
    print("")
