import faiss
from utils import timeit
from index import Index
from dataset import Dataset, Dataset_10M, Dataset_50M
from memory_profiler import profile
import gc


class Flat(Index):

    index = None

    def __init__(self) -> None:
        self.index = faiss.IndexFlat(self.d, faiss.METRIC_L2)
        assert self.index.is_trained

    @timeit
    def add(self, vector_list: list):
        self.index.add(vector_list)


def use(data_object: Dataset, index_file_name: str):
    index = Flat()
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
    use(data, "flat_10M.index")
    del data
    gc.collect()
    data = Dataset_50M()
    use(data, "flat_50M.index")
    del data
    gc.collect()


"""
Total execution time add: 3.8920249938964844 s

  4.200 ms per query, R@1 1.0000, missing rate 0.0000
Total execution time query: 41.998547077178955 s


Total execution time add: 20.221054315567017 s

 20.561 ms per query, R@1 1.0000, missing rate 0.0000
Total execution time query: 205.6123538017273 s
"""
