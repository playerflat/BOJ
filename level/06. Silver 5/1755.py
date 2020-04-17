import math

word = ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
a, b = map(int, input().split())
dic = {}
for i in range(a, b + 1):
    if i < 10:
        dic[f"{word[i]}"] = i
    else:
        dic[f"{word[math.floor(i / 10)] + word[i % 10]}"] = i

dic = sorted(dic.items())

for i in range(1, len(dic)+1):
    print(dic[i-1][1], end=' ')
    if i % 10 == 0:
        print()