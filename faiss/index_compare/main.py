import sys
from flat import Flat
from ivf_flat import IvfFlat
from ivf_pq import IvfPQ
from hnsw import Hnsw64
from hnsw_sq import HnswSQ
from hnsw_pq import HnswPQ
from utils import generate_ids
from utils import xb, nq, xq, d, xt, gt


def flat():
    flat = Flat(d)
    flat.add(xb)
    flat.query(xq, nq, gt)
    flat.save()


def ivf_flat():
    id_list = generate_ids(len(xb))
    ivf_flat = IvfFlat(d, 10000)
    ivf_flat.train(xb)
    ivf_flat.add_with_ids(xb, id_list)
    ivf_flat.query(xq, nq, gt)
    ivf_flat.save()


def ivf_pq():
    pq = IvfPQ(d, 10000)
    pq.train(xb)
    pq.add(xb)
    pq.query(xq, nq, gt)
    pq.save()


def hnsw64():
    id_list = generate_ids(len(xb))
    hnsw_flat = Hnsw64(d)
    hnsw_flat.add_with_ids(xb, id_list)
    hnsw_flat.query(xq, nq, gt)
    hnsw_flat.save()


def hnsw_sq():
    hnsw_sq = HnswSQ(d)
    hnsw_sq.train(xb)
    hnsw_sq.add(xb)
    hnsw_sq.query(xq, nq, gt)
    hnsw_sq.save()


def hnsw_pq():
    hnsw_sq = HnswPQ(d)
    hnsw_sq.train(xb)
    hnsw_sq.add(xb)
    hnsw_sq.query(xq, nq, gt)
    hnsw_sq.save()


if __name__ == "__main__":
    usage = """usage:
    python main.py FLAT/IVFFlat/IVFPQ/HNSW64/HNSWSQ/HNSWPQ
    """
    argv = sys.argv

    index_type = argv[1]
    if index_type == "FLAT":
        flat()
    elif index_type == "IVFFlat":
        ivf_flat()
    elif index_type == "IVFPQ":
        ivf_pq()
    elif index_type == "HNSW64":
        hnsw64()
    elif index_type == "HNSWSQ":
        hnsw_sq()
    elif index_type == "HNSWPQ":
        hnsw_pq()
    else:
        print("unknown index")
