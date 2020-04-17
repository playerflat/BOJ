for _ in range(int(input())):
    a = list(map(int, input().split()))
    hp = a[0] + a[4]
    mp = a[1] + a[5]
    atk = a[2] + a[6]
    if hp < 1:
        hp = 1
    if mp < 1:
        mp = 1
    if atk < 0:
        atk = 0
    print(hp + mp * 5 + atk * 2 + 2 * (a[3] + a[7]))
