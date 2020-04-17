import math

a, b, c = map(int, input().split())
print(math.floor((a+1)*(b+1)/(c+1)-1))