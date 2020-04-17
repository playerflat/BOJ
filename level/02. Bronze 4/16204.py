a, o, b, = map(int, input().split())
x = a - o

print(min(o, b) + min(a - b, a - o))
