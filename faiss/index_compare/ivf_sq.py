import faiss
from index import Index
from utils import timeit
from dataset import Dataset, Dataset_10M, Dataset_50M
from memory_profiler import profile
import gc


class IvfSQ(Index):

    index = None
    nlist = 10000
    nprobe = 100

    def __init__(self):
        quantizer = faiss.IndexFlatL2(self.d)
        self.index = faiss.IndexIVFScalarQuantizer(quantizer, self.d, self.nlist, faiss.ScalarQuantizer.QT_8bit)
        self.index.nprobe = self.nprobe

    @timeit
    def train(self, xt: list):
        self.index.train(xt)
        assert self.index.is_trained

    @timeit
    def add(self, vector_list: list):
        self.index.add(vector_list)


def use(data_object: Dataset, index_file_name: str):
    index = IvfSQ()
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
    use(data, "ivf_sq_10M.index")
    del data
    gc.collect()
    data = Dataset_50M()
    use(data, "ivf_sq_50M.index")
    del data
    gc.collect()
