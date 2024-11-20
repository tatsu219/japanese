import os
import heapq

#平仮名文を漢字に変換するプログラムを作成する．ダイクストラ法を用いて．また，辞書はmozc/mozc-dictionaryディレクトリに存在しており，辞書のtxtファイルが複数存在している．それらの辞書を読み込んで

class Node:
    def __init__(self, cost, word, path):
        self.cost = cost
        self.word = word
        self.path = path

    def __lt__(self, other):
        return self.cost < other.cost

def load_dictionary(dictionary_dir):
    dictionary = {}
    for filename in os.listdir(dictionary_dir):
        if filename.endswith(".txt"):
            with open(os.path.join(dictionary_dir, filename), 'r', encoding='utf-8') as file:
                for line in file:
                    parts = line.strip().split('\t')
                    if len(parts) == 2:
                        hiragana, kanji = parts
                        if hiragana not in dictionary:
                            dictionary[hiragana] = []
                        dictionary[hiragana].append(kanji)
    return dictionary

def dijkstra(hiragana_sentence, dictionary):
    pq = []
    heapq.heappush(pq, Node(0, '', []))
    visited = set()

    while pq:
        current_node = heapq.heappop(pq)
        current_cost = current_node.cost
        current_word = current_node.word
        current_path = current_node.path

        if current_word in visited:
            continue
        visited.add(current_word)

        if current_word == hiragana_sentence:
            return ''.join(current_path)

        for i in range(1, len(hiragana_sentence) - len(current_word) + 1):
            next_word = hiragana_sentence[len(current_word):len(current_word) + i]
            if next_word in dictionary:
                for kanji in dictionary[next_word]:
                    new_cost = current_cost + 1
                    new_path = current_path + [kanji]
                    heapq.heappush(pq, Node(new_cost, current_word + next_word, new_path))

    return None

def hiragana_to_kanji(hiragana_sentence, dictionary_dir):
    dictionary = load_dictionary(dictionary_dir)
    return dijkstra(hiragana_sentence, dictionary)

if __name__ == "__main__":
    dictionary_dir = 'mozc/mozc-dictionary'
    hiragana_sentence = 'おおきくなってください'
    kanji_sentence = hiragana_to_kanji(hiragana_sentence, dictionary_dir)
    print(kanji_sentence)