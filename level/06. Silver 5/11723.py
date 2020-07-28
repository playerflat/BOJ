import collections, sys


def add(x):
    S[x] = 1


def remove(x):
    S[x] = 0


def check(x):
    print(S[x])


def toggle(x):
    if S[x] == 1:
        S[x] = 0
    else:
        S[x] = 1


def all():
    for i in range(1, 21):
        S[i] = 1


def empty():
    for i in range(1, 21):
        S[i] = 0


S = collections.defaultdict(int)

for _ in range(int(sys.stdin.readline())):
    a = sys.stdin.readline().rstrip()
    b = ''
    if ' ' in a:
        a, b = a.split(' ')
        b = int(b)
    if a == 'add':
        add(b)
    elif a == 'remove':
        remove(b)
    elif a == 'check':
        check(b)
    elif a == 'toggle':
        toggle(b)
    elif a == 'all':
        all()
    else:
        empty()
