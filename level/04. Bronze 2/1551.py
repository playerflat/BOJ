import copy

n, k = map(int, input().split())
a = list(map(int, input().replace(',', ' ').split()))
for i in range(k):
    b = []
    for j in range(len(a) - 1):
        b.append(a[j + 1] - a[j])
    a = copy.deepcopy(b)

print(','.join(map(str, a)))