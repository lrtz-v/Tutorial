import faiss
import time
from utils import timeit
from index import Index


class Hnsw64(Index):

    index = None
    index_file_name = "hnsw64.index"
    datatype = "float32"
    d = 128
    top_k = 10
    efSearch = 64

    def __init__(self, d: int) -> None:
        self.d = d
        self.index = faiss.IndexHNSWFlat(d, 64, faiss.METRIC_L2)
        self.index.hnsw.efConstruction = 100
        self.index.hnsw.efSearch = self.efSearch
        self.index = faiss.IndexIDMap2(self.index)

    def train():
        pass

    @timeit
    def add(self, vector_list: list):
        self.index.add(vector_list)

    @timeit
    def add_with_ids(self, vector_list: list, id_list: list):
        faiss.normalize_L2(vector_list)
        print(vector_list[0])
        self.index.add_with_ids(vector_list, id_list)
