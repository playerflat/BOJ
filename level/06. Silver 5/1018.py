startw = ["WBWBWBWB", "BWBWBWBW", "WBWBWBWB", "BWBWBWBW", "WBWBWBWB", "BWBWBWBW", "WBWBWBWB", "BWBWBWBW"]
startb = ["BWBWBWBW", "WBWBWBWB", "BWBWBWBW", "WBWBWBWB", "BWBWBWBW", "WBWBWBWB", "BWBWBWBW", "WBWBWBWB"]

a, b = map(int, input().split())
c = []
m = 64

for i in range(a):
    c += input().split("\n")

for i in range(a - 7):
    for j in range(b - 7):
        z = 0
        if c[i][j] == "B":
            d = [c[i][j:j + 8], c[i + 1][j:j + 8], c[i + 2][j:j + 8], c[i + 3][j:j + 8], c[i + 4][j:j + 8],
                 c[i + 5][j:j + 8], c[i + 6][j:j + 8], c[i + 7][j:j + 8]]
            for k in range(8):
                for l in range(8):
                    if d[k][l] != startb[k][l]:
                        z += 1
            if z > 32:
                z = 64 - z
            if z < m:
                m = z
        if c[i][j] == "W":
            d = [c[i][j:j + 8], c[i + 1][j:j + 8], c[i + 2][j:j + 8], c[i + 3][j:j + 8], c[i + 4][j:j + 8],
                 c[i + 5][j:j + 8], c[i + 6][j:j + 8], c[i + 7][j:j + 8]]
            for k in range(8):
                for l in range(8):
                    if d[k][l] != startw[k][l]:
                        z += 1
            if z > 32:
                z = 64 - z
            if z < m:
                m = z
print(m)
