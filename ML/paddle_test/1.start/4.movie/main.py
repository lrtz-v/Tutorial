import pandas as pd
import numpy as np
import paddle
import paddle.nn as nn
from utils import pickle_load


def get_user_info():
    usr_file = "./data/users.dat"
    usr_info = {}
    # 打开文件，读取所有行到data中
    with open(usr_file, 'r') as f:
        data = f.readlines()
        for item in data:
            item = item.strip().split("::")
            usr_info[str(item[0])] = item
    return usr_info


def get_movie_info():
    # 电影特征的路径
    movie_data_path = "./data/movies.dat"
    mov_info = {}
    # 打开电影数据文件，根据电影ID索引到电影信息
    with open(movie_data_path, 'r', encoding="ISO-8859-1") as f:
        data = f.readlines()
        for item in data:
            item = item.strip().split("::")
            mov_info[str(item[0])] = item

    return mov_info


def get_user_rating_info(user_id):
    rating_path = "./data/ratings.dat"
    with open(rating_path, 'r') as f:
        ratings_data = f.readlines()

    usr_rating_info = {}
    for item in ratings_data:
        item = item.strip().split("::")
        # 处理每行数据，分别得到用户ID，电影ID，和评分
        usr_id,movie_id,score = item[0],item[1],item[2]
        if usr_id == str(user_id):
            usr_rating_info[movie_id] = float(score)
    return usr_rating_info


def get_user_feature():
    return pickle_load('./usr_feat.pkl')


def get_movie_feature():
    return pickle_load('./mov_feat.pkl')


user_info = get_user_info()
mov_info = get_movie_info()

usr_feats = get_user_feature()
mov_feats = get_movie_feature()


def recommand(usr_ID):
    global user_info, mov_info, usr_feats, mov_feats

    usr_ID_feat = usr_feats[str(usr_ID)]
    # 记录计算的相似度
    cos_sims = []
    # 记录下与用户特征计算相似的电影顺序

    # with dygraph.guard():
    paddle.disable_static()
    # 索引电影特征，计算和输入用户ID的特征的相似度
    for idx, key in enumerate(mov_feats.keys()):
        mov_feat = mov_feats[key]
        usr_feat = paddle.to_tensor(usr_ID_feat)
        mov_feat = paddle.to_tensor(mov_feat)

        # 计算余弦相似度
        sim = paddle.nn.functional.common.cosine_similarity(usr_feat, mov_feat)
        # 打印特征和相似度的形状
        if idx == 0:
            print("电影特征形状：{}, 用户特征形状：{}, 相似度结果形状：{}，相似度结果：{}".format(mov_feat.shape, usr_feat.shape, sim.numpy().shape, sim.numpy()))
        # 从形状为（1，1）的相似度sim中获得相似度值sim.numpy()[0]，并添加到相似度列表cos_sims中
        cos_sims.append(sim.numpy()[0])

    # 对相似度排序，获得最大相似度在cos_sims中的位置
    index = np.argsort(cos_sims)
    # 打印相似度最大的前topk个位置
    topk = 10
    print("相似度最大的前{}个索引是{}\n对应的相似度是：{}\n".format(topk, index[-topk:], [cos_sims[k] for k in index[-topk:]]))

    for i in index[-topk:]:
        print("对应的电影分别是：movie:{}".format(mov_info[list(mov_feats.keys())[i]]))


def get_user_rating(usr_a):
    usr_rating_info = get_user_rating_info(usr_a)
    movie_ids = list(usr_rating_info.keys())
    print("ID为 {} 的用户，评分过的电影数量是: ".format(usr_a), len(movie_ids))

    #####################################
    ## 选出ID为usr_a评分最高的前topk个电影 ##
    #####################################
    topk = 10
    ratings_topk = sorted(usr_rating_info.items(), key=lambda item:item[1])[-topk:]
    for k, score in ratings_topk:
        print("电影ID: {}，评分是: {}, 电影信息: {}".format(k, score, mov_info[k]))


if __name__ == "__main__":
    usr_id = 2
    recommand(usr_id)
    get_user_rating(usr_id)
