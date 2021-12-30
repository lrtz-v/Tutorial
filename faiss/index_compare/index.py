import faiss
import time
from utils import timeit


class Index:
    d = 128
    index = None
    datatype = "float32"
    top_k = 10

    def train(self, xt: list):
        pass

    @timeit
    def query(self, query: list, nq, gt):
        t0 = time.time()
        D, I = self.index.search(query, self.top_k)
        t1 = time.time()

        missing_rate = (I == -1).sum() / float(self.top_k * nq)
        recall_at_1 = (I == gt[:, :1]).sum() / float(nq)
        print("%7.3f ms per query, R@1 %.4f, missing rate %.4f" %
              ((t1 - t0) * 1000.0 / nq, recall_at_1, missing_rate))

    def save(self, filename: str):
        faiss.write_index(self.index, "./index/" + filename)
