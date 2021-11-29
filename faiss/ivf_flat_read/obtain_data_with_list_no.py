import faiss
import numpy as np
import sys
from utils import fvecs_write, timeit


datatype = 'float32'
d = 128


@timeit
def load_index(index_filename: str) -> faiss.IndexIVFFlat:
    index = faiss.read_index(index_filename)
    index.set_direct_map_type(faiss.DirectMap.Hashtable)
    return index


@timeit
def get_ids(invlists, list_no: int):
    invlists = faiss.downcast_InvertedLists(invlists)
    ls = invlists.list_size(list_no)
    list_ids = np.zeros(ls, dtype='int64')
    ids = None
    origin_ids = None

    try:
        ids = invlists.get_ids(list_no)
        origin_ids = np.array(faiss.rev_swig_ptr(ids, ls))
    except Exception as e:
        print("get_ids failed ", e)
    finally:
        if ids is not None:
            invlists.release_ids(list_no, ids)
    return origin_ids


@timeit
def get_origin_ids(index: faiss.IndexIVFFlat, list_no: int):
    ls = index.invlists.list_size(list_no)
    if ls == 0:
        return []

    ids = get_ids(index.invlists, list_no)
    origin_ids = ids.tolist()
    return origin_ids


@timeit
def obtain_origin_data(index: faiss.IndexIVFFlat, origin_ids: list, list_no: int):
    codes = []
    for origin_id in origin_ids:
        codes.append(index.reconstruct(origin_id))

    filename = "fvecs/{}.fvecs".format(list_no)
    fvecs_write(filename, np.array(codes))
    print("obtain codes: {}, saved: {}".format(len(codes), filename))


if __name__ == "__main__":
    usage = """Usage:
    python main.py index_file_path 1-10000\n"""

    argv = sys.argv
    if len(argv) < 3:
        print(usage)
        exit(1)
    index_file_name = argv[1]
    list_no_range = int(argv[2])
    if list_no_range < 1 or list_no_range > 10000:
        print(usage)
        exit(1)

    index = load_index(index_file_name)

    for list_no in range(list_no_range):
        print("{}/{}...".format(list_no, list_no_range-1))
        origin_ids = get_origin_ids(index, list_no)
        if len(origin_ids) == 0:
            continue
        obtain_origin_data(index, origin_ids, list_no)
    print("[obtain finished]")
