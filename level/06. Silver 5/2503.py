answer = {}
for i in range(123, 988):
    s = {i % 10, i // 10 % 10, i // 100}
    if len(s) != 3 or 0 in s:
        continue
    answer[i] = -1

for _ in range(int(input())):
    number, strike, ball = map(int, input().split())
    for i, v in answer.items():
        if answer[i] == 0:
            continue
        ans_strike = (i % 10 == number % 10) + (i // 10 % 10 == number // 10 % 10) + \
                     (i // 100 == number // 100)
        ans_ball = (i % 10 == number // 10 % 10 or i % 10 == number // 100) + (
                i // 10 % 10 == number % 10 or i // 10 % 10 == number // 100) + (
                           i // 100 == number % 10 or i // 100 == number // 10 % 10)
        if ans_strike == strike and ans_ball == ball:
            answer[i] = 1
        else:
            answer[i] = 0

count = 0
for i, v in answer.items():
    if v == 1:
        count += 1

print(count)
