import os
import heapq
from itertools import combinations


def read_mozc_dictionary(mozc_path):
    files = os.listdir(mozc_path)
    del files[-1]
    dic = {}
    
    for file in files:
        with open(mozc_path + '/' + file, 'r', encoding='utf-8') as f:
            for line in f:
                parts = line.strip().split('\t')
                hiragana = parts[0]
                kanji = parts[4]
                cost = int(parts[3])
                if hiragana not in dic:
                    dic[hiragana] = []
                dic[hiragana].append((kanji, cost))
    # print(dic)
    return dic

# def all_splits(string):
    # インデックス位置の全通りを生成
    splits = []
    for i in range(1, len(string)):  # 分割点の数
        for comb in combinations(range(1, len(string)), i):  # 分割点を選ぶ
            parts = []
            prev = 0
            for point in comb:
                parts.append(string[prev:point])
                prev = point
            parts.append(string[prev:])  # 最後の部分
            splits.append(parts)
    print(splits)
    return splits



def convert_text(text,conversion_dict):
    n = len(text)
    dp = [float('inf')] * (n + 1)  # コストを保存
    dp[0] = 0
    prev = [-1] * (n + 1)  # 復元用のインデックス
    best_candidate = [None] * (n + 1)  # 最適な単語を保存
    pq = [(0, 0)]  # 優先度キュー (コスト, 現在のインデックス)

    max_len = max(len(key) for key in conversion_dict)  # 辞書内の最長の単語長

    while pq:
        current_cost, index = heapq.heappop(pq)
        if current_cost > dp[index]:
            continue

        for length in range(1, min(max_len, n - index) + 1):
            candidate = text[index:index + length]
            if candidate in conversion_dict:
                for kanji, cost in conversion_dict[candidate]:
                    next_index = index + length
                    new_cost = dp[index] + cost
                    if new_cost < dp[next_index]:
                        dp[next_index] = new_cost
                        prev[next_index] = index
                        best_candidate[next_index] = (candidate, kanji)
                        heapq.heappush(pq, (new_cost, next_index))

    # 漢字列を復元
    result = []
    idx = n
    while idx > 0:
        if best_candidate[idx]:
            result.append(best_candidate[idx][1])  # 最適な漢字を取得
            idx = prev[idx]
        else:
            result.append(text[idx - 1])  # 変換できない場合、元の文字を保持
            idx -= 1
    return ''.join(result[::-1])  # 逆順なのでひっくり返す


if __name__ == '__main__':
    
    mozc_path = 'mozc/mozc-dictionary'
    input = 'あしたはあめがふるかのうせいがある'
    dic = read_mozc_dictionary(mozc_path)
    # print(all_splits(input))
    text = convert_text(input, dic)
    print(text)