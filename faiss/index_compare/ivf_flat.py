import faiss
import time
from utils import timeit
from index import Index


class IvfFlat(Index):

    index = None
    index_file_name = "ivf_flat.index"
    datatype = "float32"
    nlist = 10000
    d = 128
    nprobe = 100
    top_k = 10

    def __init__(self, d: int, nlist: int) -> None:
        self.d = d
        self.nlist = nlist
        quantizer = faiss.IndexFlatL2(d)
        self.index = faiss.IndexIVFFlat(quantizer, d, nlist, faiss.METRIC_L2)
        self.index.nprobe = self.nprobe

    @timeit
    def train(self, xt: list):
        self.index.train(xt)
        assert self.index.is_trained

    @timeit
    def add_with_ids(self, vector_list: list, id_list: list):
        self.index.add_with_ids(vector_list, id_list)
