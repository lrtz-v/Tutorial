# import faiss
import numpy as np
from sklearn.metrics.pairwise import cosine_similarity
import struct
import base64


def feature_encode(float_vector_list: list):
    buf = bytes()
    for val in float_vector_list:
        buf += struct.pack('f', val)
    return base64.standard_b64encode(buf)


def feature_decode(feature: str):
    data = base64.standard_b64decode(feature)
    count = len(data) // 4

    return struct.unpack('<{0}f'.format(count), data)


def simility(arr1, arr2):
    return cosine_similarity(np.array([arr1]), np.array([arr2]))
