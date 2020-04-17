from collections import deque
import sys

deq = deque()

for i in range(int(input())):
    b = sys.stdin.readline().split()

    if b[0] == "push_front":
        deq.appendleft(b[1])
    elif b[0] == "push_back":
        deq.append(b[1])
    elif b[0] == "pop_front":
        if len(deq) == 0:
            print(-1)
        else:
            print(deq.popleft())
    elif b[0] == "pop_back":
        if len(deq) == 0:
            print(-1)
        else:
            print(deq.pop())
    elif b[0] == "size":
        print(len(deq))
    elif b[0] == "empty":
        if len(deq) == 0:
            print(1)
        else:
            print(0)
    elif b[0] == "front":
        if len(deq) == 0:
            print(-1)
        else:
            print(deq[0])
    else:
        if len(deq) == 0:
            print(-1)
        else:
            print(deq[-1])