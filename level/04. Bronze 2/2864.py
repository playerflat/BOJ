a = input().split()

a[0] = a[0].replace('6', '5')
a[1] = a[1].replace('6', '5')

print(int(a[0]) + int(a[1]), end=' ')

a[0] = a[0].replace('5', '6')
a[1] = a[1].replace('5', '6')

print(int(a[0]) + int(a[1]))
