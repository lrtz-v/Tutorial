import faiss
from utils import timeit
from index import Index


class Flat(Index):

    index = None
    index_file_name = "flat.index"

    def __init__(self, d: int) -> None:
        self.d = d
        self.index = faiss.IndexFlat(d, faiss.METRIC_L2)
        assert self.index.is_trained

    @timeit
    def add(self, vector_list: list):
        self.index.add(vector_list)
