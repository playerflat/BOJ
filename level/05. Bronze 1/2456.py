import collections

d = collections.defaultdict(int)
for _ in range(int(input())):
    score = list(map(int, input().split()))
    for i in range(3):
        d[i + 1] += score[i]
        d[int(f'{i + 1}{score[i]}')] += 1

score = [d[1], d[2], d[3]]
three = [d[13], d[23], d[33]]
two = [d[12], d[22], d[32]]

if (sum(score) - min(score)) / 2 == max(score):
    if (sum(three) - min(three)) / 2 == max(three):
        two[three.index(min(three))] = 0
        if (sum(two) - min(two)) / 2 == max(two):
            print(0, max(score))
        else:
            print(two.index(max(two)) + 1, max(score))
    else:
        print(three.index(max(three)) + 1, max(score))
else:
    print(score.index(max(score)) + 1, max(score))
