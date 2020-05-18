a = int(input())
print(sum(i for i in range(1, a + 1) if a % i == 0) * 5 - 24)
