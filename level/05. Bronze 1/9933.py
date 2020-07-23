def calc():
    for i in range(N):
        for j in range(i, N):
            if a[i][::-1] == a[j]:
                print(len(a[i]), a[i][len(a[i]) // 2])
                return


N = int(input())
a = []
for i in range(N):
    a.append(input())

calc()
