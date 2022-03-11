import paddle
import pickle


# 加载模型参数
def load_params(params_file_path):
    return paddle.load(params_file_path)


def pickle_dump(data_dict, filepath):
    pickle.dump(data_dict, open(filepath, 'wb'))


def pickle_load(filepath):
    return pickle.load(open(filepath, 'rb'))
