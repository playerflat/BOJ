import sys
import math


def init(size):
    for i in range(size - 1, 0, -1):
        mintree[i] = min(mintree[i * 2], mintree[i * 2 + 1])
        maxtree[i] = max(maxtree[i * 2], maxtree[i * 2 + 1])


def mini(left, right, nleft, nright, nnum):
    if left > nright or right < nleft:
        return 10 ** 9
    if left <= nleft and right >= nright:
        return mintree[nnum]
    m = (nleft + nright) // 2
    return min(mini(left, right, nleft, m, nnum * 2), mini(left, right, m + 1, nright, nnum * 2 + 1))


def maxi(left, right, nleft, nright, nnum):
    if left > nright or right < nleft:
        return 1
    if left <= nleft and right >= nright:
        return maxtree[nnum]
    m = (nleft + nright) // 2
    return max(maxi(left, right, nleft, m, nnum * 2), maxi(left, right, m + 1, nright, nnum * 2 + 1))


n, m = map(int, sys.stdin.readline().split())
size = 2 ** math.ceil(math.log(n, 2))
msize = size * 2
maxtree = [10 ** 9] * msize
mintree = [10 ** 9] * msize
for i in range(n):
    num = int(sys.stdin.readline())
    maxtree[size + i] = num
    mintree[size + i] = num
init(size)
for _ in range(m):
    x, y = map(int, sys.stdin.readline().split())
    print(mini(x - 1, y - 1, 0, size - 1, 1), maxi(x - 1, y - 1, 0, size - 1, 1))
