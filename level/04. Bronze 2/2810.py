b = int(input())
a = input().replace('S', '*S*').replace('LL', '*LL*').replace('**', '*')
print(b if a.count('*') > b else a.count('*'))
