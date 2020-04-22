import sys

for _ in range(int(sys.stdin.readline())):
    s = 0
    for _ in range(int(sys.stdin.readline())):
        x, y = sys.stdin.readline().split()
        if x == y:
            pass
        elif (x == "R" and y == "S") or (x == "P" and y == "R") or (x == "S" and y == "P"):
            s += 1
        else:
            s -= 1
    if s == 0:
        print("TIE")
    elif s > 0:
        print("Player 1")
    else:
        print("Player 2")
