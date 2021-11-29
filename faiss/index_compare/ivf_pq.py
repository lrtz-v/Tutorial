import faiss
from index import Index
from utils import timeit


class IvfPQ(Index):

    index = None
    index_file_name = "ivf_pq.index"
    datatype = "float32"
    nlist = 10000
    d = 128
    nprobe = 100
    top_k = 10

    def __init__(self, d: int, nlist: int):
        quantizer = faiss.IndexFlatL2(d)
        self.index = faiss.IndexIVFPQ(quantizer, d, nlist, 16, 8)
        self.index.nprobe = 100

    @timeit
    def train(self, xt: list):
        self.index.train(xt)
        assert self.index.is_trained

    @timeit
    def add(self, vector_list: list):
        self.index.add(vector_list)
