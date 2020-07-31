'''
BOJ 코드 난이도 분류

개요
Solved.ac의 난이도 분류에 따라 BOJ 코드를 폴더에 이동시킵니다.

사용법
1. 문제번호.확장자 파일을 sourcecode 폴더에 추가
2. 본 코드 실행

흐름
1. Solved.ac의 난이도별 문제 목록을 크롤링하여 Level.csv로 저장
(csv 파일이 존재하지 않거나, 수정된지 10일이 지났으면 다시 크롤링)
2. 난이도별 폴더(level의 하위 폴더)를 생성
3. sourcecode 폴더의 문제번호와 일치하는 난이도의 폴더로 파일 이동

주의사항
1. Solved.ac 난이도 체계 바뀌면 다시 코드 짜야 함
2. 실행하기 전에 sourcecode 폴더는 미리 백업해놓기
'''

import pandas as pd
import os
import time
import requests
from bs4 import BeautifulSoup


def download_level():
    df = pd.DataFrame(data={'Level': [], 'Number': []}, columns=['Level', 'Number'])
    df_index = 0
    for i in range(31):  # 난이도 0(unrated) ~ 30(Ruby I)
        c = 0
        for j in range(1, 100):  # 특정 난이도 내 페이지 개수
            res = requests.get(f'https://solved.ac/problems/level/{i}?page={j}')
            soup = BeautifulSoup(res.content, 'html.parser')
            r = soup.find_all('span')
            c += 1
            if c > 1:  # 페이지에 항목 없으면 다음 난이도로 이동
                break
            for tag in r:
                if tag.text.isnumeric():
                    c = 0
                    df.loc[df_index] = [i, tag.text.strip()]
                    df_index += 1
            time.sleep(0.1)
        print(f'complete {i}')
    df.to_csv('./Level.csv', index=False)


# 난이도별 문제 번호가 적힌 csv 파일 불러오기
# csv 파일이 없거나 수정된 지 10일 이상 지난 파일이면 solved.ac 크롤링
level = './Level.csv'
if not os.path.isfile(level) or (time.time() - os.stat(level).st_mtime) / 86400 > 10:
    download_level()
level = pd.read_csv(level).set_index('Number')

# 난이도 딕셔너리 생성
level_dict = {0: '00. Unrated'}
i = 1
for j in ['Bronze ', 'Silver ', 'Gold ', 'Platinum ', 'Diamond ', 'Ruby ']:
    for k in range(5, 0, -1):
        level_dict[i] = f'{i:02d}. {j}{k}'
        i += 1

# 난이도 별 폴더 생성
if not os.path.isdir('./level'):
    os.makedirs('./level')
for i in level_dict:
    if not os.path.isdir(f'./level/{level_dict[i]}'):
        os.makedirs(f'./level/{level_dict[i]}')

# 난이도 별 파일 이동
# csv 파일의 Level 열에 기준하여 파일 이동
sourcecode_list = os.listdir('./sourcecode/')
for i in range(len(sourcecode_list)):
    lv = int(level.at[int(sourcecode_list[i].split(".")[0]), 'Level'])
    os.rename(f'./sourcecode/{sourcecode_list[i]}', f'./level/{level_dict[lv]}/{sourcecode_list[i]}')

# 난이도 변경 시 폴더에서 이동
# level 폴더의 코드파일이 속한 난이도와 csv의 Level 열과 비교 후 이동
level_list = os.listdir('./level/')
for i in range(len(level_list)):
    level_code_list = os.listdir(f'./level/{level_list[i]}')
    for j in range(len(level_code_list)):
        lv = int(level.at[int(level_code_list[j].split(".")[0]), 'Level'])
        if lv != i:
            os.rename(f'./level/{level_list[i]}/{level_code_list[j]}', f'./level/{level_list[lv]}/{level_code_list[j]}')
