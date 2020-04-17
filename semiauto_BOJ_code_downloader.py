'''
BOJ 코드 반자동 다운로더 (Windows)

개요
BOJ에서 내가 맞은 코드를 파일로 저장합니다.

사용법
1. C드라이브의 최상위 경로에 ChromeTemp 폴더 생성하기
2. 크롬 실행파일이 있는 디렉토리 찾기(대부분 C:/Program Files (x86)/Google/Chrome/Application)
3. 크롬 실행파일이 있는 디렉토리에 들어가면 크롬 버전이 써진 폴더가 있는데, 이 버전에 맞춰 크롬드라이버 다운로드
    크롬 버전이 80.0.3987.163 이면 ChromeDriver 80.x.x.x 다운로드
    크롬드라이버 다운로드 : https://chromedriver.chromium.org/downloads
4. 다운받은 크롬드라이버 압축파일에서 exe만 빼내서 C드라이브의 최상위 경로에 복사
5. cmd - 크롬 실행파일이 있는 디렉토리로 이동 (cd C:/Program Files (x86)/Google/Chrome/Application)
6. cmd에서 chrome.exe --remote-debugging-port=9222 --user-data-dir="C:/ChromeTemp" 입력 - 디버깅모드 크롬 켜짐
7. 디버깅모드 크롬으로 BOJ 접속 - 로그인 - 우측 상단의 아이디 클릭 - 좌측의 맞은 문제 숫자 클릭
8. 본 코드 실행
9. exit 0 떴으면 8번부터 반복

흐름
코드 실행 - 소스코드 다운로드 - 파일 제목을 '제출번호'에서 '문제번호'로 변경하여 sourcecode 폴더에 저장 - 다음 페이지 클릭

문제점
1. 반자동이라 제출물이 많으면 귀찮음
2. 모든 정답을 받기 때문에, 한 문제에 대한 여러 정답은 수동으로 선별해야 함
3. py3, go, java, c, txt 형식으로 다운로드 되는 파일만 이동함
4. 혹시 모르니 다운로드 폴더를 미리 백업해놓고 사용
'''

from selenium import webdriver
import os
import time

if not os.path.isdir('./sourcecode'):
    os.makedirs('./sourcecode')

chrome_options = webdriver.ChromeOptions()
chrome_options.add_experimental_option("debuggerAddress", "127.0.0.1:9222")
chrome_driver = "C:/chromedriver.exe"
driver = webdriver.Chrome(chrome_driver, options=chrome_options)

lang_dict = {}
problem_number = {}
download_lists = []


def filename_extension(lang):
    extensions = driver.find_elements_by_partial_link_text(lang)
    for extension in extensions:
        if len(lang_dict) == 20:
            break
        lang_dict[extension.get_attribute("href").split("/")[4]] = lang.lower()


def download_files():
    submits = driver.find_elements_by_link_text("수정")
    for submit in submits:
        split = submit.get_attribute("href").split("/")
        problem_number[split[5]] = split[4]
        download_lists.append(split[5])
        url = f"https://www.acmicpc.net/source/download/{split[5]}"
        driver.get(url)


def many_answer(chrome_download_path, destination, download, old, new):
    c = 1
    if not os.path.isfile(f"{destination}/{problem_number[download]}.{new}"):
        os.replace(f"{chrome_download_path}/{download}.{old}", f"{destination}/{problem_number[download]}.{new}")
    else:
        while True:
            if not os.path.isfile(f"{destination}/{problem_number[download]}({c}).{new}"):
                os.replace(f"{chrome_download_path}/{download}.{old}",
                           f"{destination}/{problem_number[download]}({c}).{new}")
                break
            else:
                c += 1


def move_files(chrome_download_path, destination):
    for download in download_lists:
        if lang_dict[download] == "py":
            many_answer(chrome_download_path, destination, download, "py3", "py")
        elif lang_dict[download] == "go":
            many_answer(chrome_download_path, destination, download, "go", "go")
        elif lang_dict[download] == "java":
            many_answer(chrome_download_path, destination, download, "java", "java")
        elif lang_dict[download] == "text":
            many_answer(chrome_download_path, destination, download, "txt", "txt")
        elif lang_dict[download] == "c":
            many_answer(chrome_download_path, destination, download, "c", "c")


def next_page():
    try:
        driver.find_element_by_link_text("다음 페이지").click()
    except Exception:
        print("마지막 페이지")


for i in ["Py", "Go", "Java", "Text", "C"]:
    filename_extension(i)
download_files()
time.sleep(3)
move_files("C:/Users/user/Downloads", "./sourcecode")  # 다운로드된 파일 경로는 적절하게 수정
next_page()
