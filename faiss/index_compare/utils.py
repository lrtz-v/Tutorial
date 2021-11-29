import numpy as np
import time
import random
from functools import wraps
from faiss.contrib.datasets import DatasetSIFT1M


ds = DatasetSIFT1M()

xq = ds.get_queries()
xb = ds.get_database()
gt = ds.get_groundtruth()
xt = ds.get_train()

nq, d = xq.shape


def generate_data(d: int, size: int, mu: int, sigma: float, datatype: str):
    np.random.seed(0)
    data = []
    for i in range(size):
        data.append(np.random.normal(mu, sigma, d))
    return np.array(data).astype(datatype)


def generate_ids(size):
    ids = []
    for i in list(range(size)):
        ids.append(i * (random.randint(1, size) // 100))
    return np.array(ids)


def generate_query(d: int, size: int, mu: int, sigma: float):
    np.random.seed(12)
    query = []
    for i in range(size):
        query.append(np.random.normal(mu, sigma, d))
    return np.array(query).astype('float32')


def timeit(func):
    @wraps(func)
    def _time_it(*args, **kwargs):
        start = time.time()
        try:
            return func(*args, **kwargs)
        finally:
            end_ = time.time() - start
            print(
                f"Total execution time {func.__name__}: {end_ if end_ > 0 else 0} s\n"
            )

    return _time_it
