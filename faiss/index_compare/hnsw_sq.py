import faiss
import time
from utils import timeit
from index import Index


class HnswSQ(Index):

    index = None
    index_file_name = "hnswsq.index"
    datatype = "float32"
    d = 128
    top_k = 10
    efSearch = 64

    def __init__(self, d: int) -> None:
        self.d = d
        self.index = faiss.IndexHNSWSQ(d, faiss.ScalarQuantizer.QT_8bit, 64)
        self.index.hnsw.efConstruction = 40
        self.index.hnsw.efSearch = self.efSearch

    @timeit
    def train(self, vector_list: list):
        self.index.train(vector_list)

    @timeit
    def add(self, vector_list: list):
        self.index.add(vector_list)
