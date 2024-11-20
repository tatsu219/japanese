import heapq
import os

class DijkstraConverter:
    def __init__(self, dictionary_path):
        self.graph = self.load_dictionary(dictionary_path)

    def load_dictionary(self, mozc_path):
        
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

    def convert(self, hiragana_string):
        queue = [(0, '', hiragana_string)]
        visited = set()

        while queue:
            cost, kanji_string, remaining_hiragana = heapq.heappop(queue)

            if not remaining_hiragana:
                return kanji_string

            if (kanji_string, remaining_hiragana) in visited:
                continue

            visited.add((kanji_string, remaining_hiragana))

            for i in range(1, len(remaining_hiragana) + 1):
                hiragana_substring = remaining_hiragana[:i]
                if hiragana_substring in self.graph:
                    for kanji in self.graph[hiragana_substring]:
                        new_cost = cost + 1
                        new_kanji_string = kanji_string + kanji
                        new_remaining_hiragana = remaining_hiragana[i:]
                        heapq.heappush(queue, (new_cost, new_kanji_string, new_remaining_hiragana))

        return None

if __name__ == "__main__":
    dictionary_path = 'mozc/mozc-dictionary/'
    converter = DijkstraConverter(dictionary_path)
    hiragana_string = 'あしたはてんきになるです'
    kanji_string = converter.convert(hiragana_string)
    print(f'Converted: {kanji_string}')