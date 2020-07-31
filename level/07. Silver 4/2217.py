import sys

rope = []
for i in range(int(sys.stdin.readline())):
    rope.append(int(sys.stdin.readline()))

rope.sort(reverse=True)

answer = 0
for i in range(len(rope)):
    if answer < rope[i] * (i + 1):
        answer = rope[i] * (i + 1)
print(answer)
