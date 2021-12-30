from utils import bvecs_mmap, fvecs_write, fvecs_mmap, ivecs_read, sanitize
import datetime

"""
data source: http://corpus-texmex.irisa.fr/
"""


class Dataset:
    num = 0
    basedir = "./data/bigann_base_{}.fvecs"
    traindir = "./data/bigann_learn.bvecs"
    querydir = "./data/bigann_query.bvecs"
    groundtruth_idx_dir = ""

    def get_base_iterator(self):
        for i in range(self.num):
            filename = self.basedir.format(i)
            yield sanitize(fvecs_mmap(filename))

    def get_groundtruth(self):
        return ivecs_read(self.groundtruth_idx_dir)

    def get_train(self):
        return sanitize(bvecs_mmap(self.traindir))

    def get_query(self):
        return sanitize(bvecs_mmap(self.querydir))


class Dataset_10M(Dataset):
    groundtruth_idx_dir = "./data/idx_10M.ivecs"

    def __init__(self) -> None:
        super().__init__()
        self.num = 1


class Dataset_50M(Dataset):
    groundtruth_idx_dir = "./data/idx_50M.ivecs"

    def __init__(self) -> None:
        super().__init__()
        self.num = 5


def cut_data():
    # bigann_base.bvecs contains 1B vertors, after split operation, we got five files, each contain 10M vectors
    one_B_base_dir = "~/bigann_base.bvecs"
    x = bvecs_mmap(one_B_base_dir)

    file_num = 0
    for i in range(0, 50000000, 10000000):
        data = x[i:i + 10000000].copy()
        fvecs_write("./data/bigann_base_{}.fvecs".format(file_num), data)
        file_num += 1


if __name__ == "__main__":
    cut_data()
