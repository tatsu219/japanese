import os
import heapq


def read_mozc_dictionary(mozc_path):
    files = os.listdir(mozc_path)
    del files[-1]
    dic = {}
    
    for file in files:
        with open(mozc_path + '/' + file, 'r', encoding='utf-8') as f:
            for line in f:
                parts = line.strip().split('\t')
                hiragana = parts[0]
                kanji = parts[-1]
                cost = parts[-2]
                if hiragana not in dic:
                    dic[hiragana] = []
                dic[hiragana].append((kanji, cost))
    return dic

def convert_text(text,dic):
    #ダイクストラ法
    t_num = len(text)
    dp = [float('inf')] * (t_num + 1) #table
    dp[0] = 0
    prev_node = [-1] * (t_num + 1)
    best_word = [''] * (t_num + 1)
    pq = [(0, 0)]
    
    max_len = max(len(key) for key in dic)
    
    while pq:
        current_cost, index = heapq.heappop(pq)
        if current_cost > dp[index]:
            continue
        
        for length in range(1,min(max_len,t_num - index) + 1):
            word = text[index:index + length]
            if word in dic:
                for kanji, cost in dic[word]:
                    next_index = index + length
                    new_cost = dp[index] + int(cost)
                    if new_cost < dp[next_index]:
                        dp[next_index] = new_cost
                        prev_node[next_index] = index
                        best_word[next_index] = (word,kanji)
                        heapq.heappush(pq, (new_cost, next_index))
                             
    result = []
    idx = t_num
    while idx > 0:
        if best_word[idx]:
            result.append(best_word[idx][1])
            idx = prev_node[idx]
        else:
            result.append(text[idx - 1])
            idx -= 1
    return ''.join(result[::-1])


if __name__ == '__main__':
    
    mozc_path = 'mozc/mozc-dictionary'
    input = 'あしたはあめがふるかもしれない'
    dic = read_mozc_dictionary(mozc_path)
    
    text = convert_text(input,dic)
    print(text)