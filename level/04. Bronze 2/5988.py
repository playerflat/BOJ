import sys

for _ in range(int(sys.stdin.readline())):
    a = sys.stdin.readline().strip()
    if a[-1] in ['0', '2', '4', '6', '8']:
        print('even')
    else:
        print('odd')
