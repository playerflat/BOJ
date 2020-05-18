r = [1 for i in range(100)]

for i in range(2, 100):
    for j in range(i, 100):
        if j % i == 0:
            if r[j-1] == 1:
                r[j-1] = 0
            else:
                r[j-1] = 1

for _ in range(int(input())):
    print(sum(r[0:int(input())]))
