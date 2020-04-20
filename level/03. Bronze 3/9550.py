for i in range(int(input())):
    a = list(map(int, input().split()))
    b = list(map(int, input().split()))
    print(sum(j//a[1] for j in b))
