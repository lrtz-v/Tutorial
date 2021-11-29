import numpy as np
import time
from functools import wraps
import itertools


def ivecs_read(fname):
    a = np.fromfile(fname, dtype='int32')
    d = a[0]
    return a.reshape(-1, d + 1)[:, 1:].copy()


def fvecs_read(fname):
    return ivecs_read(fname).view('float32')


def mmap_fvecs(fname):
    x = np.memmap(fname, dtype='int32', mode='r')
    d = x[0]
    return x.view('float32').reshape(-1, d + 1)[:, 1:]


def ivecs_write(fname, m):
    n, d = m.shape
    m1 = np.empty((n, d + 1), dtype='int32')
    m1[:, 0] = d
    m1[:, 1:] = m
    m1.tofile(fname)


def fvecs_write(fname, m):
    # m = m.astype('float32')
    ivecs_write(fname, m.view('int32'))


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


def partition(l, size):
    for i in range(0, len(l), size):
        yield list(itertools.islice(l, i, i + size))
