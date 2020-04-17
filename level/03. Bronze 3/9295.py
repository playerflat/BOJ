import sys

for i in range(1, int(sys.stdin.readline())+1):
    print(f'Case {i}:', sum(map(int, sys.stdin.readline().split())))
