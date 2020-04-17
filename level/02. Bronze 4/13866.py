a = list(map(int, input().split()))
t = max(a) + min(a)
print(abs(t - (sum(a)-t)))
