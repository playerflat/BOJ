import sys

for _ in range(int(sys.stdin.readline())):
    a = list(map(float, sys.stdin.readline().split()))
    print(f'{int(a[0])} {a[1]/(a[2]+a[3])*a[4]}')
