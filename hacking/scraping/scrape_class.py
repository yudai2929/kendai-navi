from pprint import pprint
import requests
from bs4 import BeautifulSoup

from ng_words import ng_words
from ok_words import ok_words
import re


def get_scraped_urls() -> list[str]:
    """
    スクレイピング対象のURLを取得する
    """
    url = "https://www.aichi-pu.ac.jp/disclosure/credit"
    paths = [
        "foreign_studies.html",
        "japanese_studies.html",
        "education_and_welfare.html",
        "nursing_and_health.html",
        "information_science_and_technology.html",
        "graduate_information_science_and_technology.html",
    ]
    return [url + "/" + path for path in paths]


#
def scrape_class_names(url: str) -> list[str]:
    """
    urlをスクレイピングしてクラス名を取得する
    """

    def to_th_list_from_tbody_list(tbody_list: list):
        table_header_list = []
        for tbody in tbody_list:
            table_header_list.extend(tbody.find_all("th"))
        return table_header_list

    # Requestsを使ってURLからHTMLを取得する
    response = requests.get(url)
    if response.status_code != 200:
        raise Exception(
            "Failed to access the site. Status code: " + str(response.status_code)
        )

    # BeautifulSoupを使ってHTMLを解析する
    soup = BeautifulSoup(response.content, "html.parser")

    tbody_list = soup.find_all("tbody")

    # tbodyタグに含まれるthタグを取得する
    table_header_list = to_th_list_from_tbody_list(tbody_list)

    class_names = [headline.text.strip() for headline in table_header_list]

    return class_names


def filter_class_names(class_names: list[str], ng_words: list[str]) -> list[str]:
    """
    講義名ではないものを取り除く
    """

    def valid_class_name(class_name: str) -> bool:
        # "計（〇〇）"のパターンをマッチする正規表現
        pattern = re.compile(r"計（.*?）")

        # クラス名がパターンを含まない場合にTrueを返す
        return pattern.search(class_name) is None

    def is_ng_word(class_name: str) -> bool:
        if not valid_class_name(class_name):
            return False

        return not any(nw == class_name for nw in ng_words)

    return list(filter(is_ng_word, class_names))


def write_class_names_to_csv(class_names: list):
    """
    クラス名をcsvファイルに書き込む
    """

    def gen_filename():
        import datetime

        return (
            "class_names_" + datetime.datetime.now().strftime("%Y%m%d_%H%M%S") + ".csv"
        )

    file_name = gen_filename()
    file_path = "scraping/.dist/" + file_name
    with open(file_path, "w") as f:
        for class_name in class_names:
            f.write(class_name + "\n")


def main():
    print("Scraping started...")

    # スクレイピング対象のURLを取得
    urls = get_scraped_urls()

    class_names = []
    for url in urls:
        try:
            scraped_class_names = scrape_class_names(url)
            # いらない単語を取り除く
            filtered = filter_class_names(scraped_class_names, ng_words)
            class_names = class_names + filtered
        except Exception as e:
            print("Failed to scrape the site: " + url)
            print(e)
            break
    # スクレイピングで取得し切れな授業名を追加
    class_names = class_names + ok_words
    # 重複している可能性があるので取り除く
    unique_class_names = list(set(class_names))

    # csvファイルに書き込む
    write_class_names_to_csv(unique_class_names)

    print("Scraping finished.")
    return


if __name__ == "__main__":
    main()
