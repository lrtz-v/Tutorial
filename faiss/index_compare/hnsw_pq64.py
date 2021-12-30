import faiss
from utils import timeit
from index import Index
from dataset import Dataset, Dataset_10M, Dataset_50M
from memory_profiler import profile
import gc


class HnswPQ64(Index):

    index = None

    def __init__(self) -> None:
        self.index = faiss.IndexHNSWPQ(self.d, 64, 64)
        self.index.hnsw.efConstruction = 80
        self.index.hnsw.efSearch = 64

    @timeit
    def train(self, vector_list: list):
        self.index.train(vector_list)

    @timeit
    def add(self, vector_list: list):
        self.index.add(vector_list)


def use(data_object: Dataset, index_file_name: str):
    index = HnswPQ64()
    index.train(data_object.get_train())

    for data in data_object.get_base_iterator():
        index.add(data)
        del data
        gc.collect()

    xq = data_object.get_query()
    nq, _ = xq.shape
    index.query(xq, int(nq), data_object.get_groundtruth())
    index.save(index_file_name)
    del index


if __name__ == "__main__":
    gc.disable()
    data = Dataset_10M()
    use(data, "hnsw_pq64_10M.index")
    del data
    gc.collect()
    data = Dataset_50M()
    use(data, "hnsw_pq64_50M.index")
    del data
    gc.collect()
