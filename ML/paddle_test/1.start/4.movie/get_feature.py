import paddle
import numpy as np
from utils import pickle_dump
from model import Model
from utils import load_params


# 定义将list中每个元素转成tensor的函数
def list2tensor(inputs, shape):
    inputs = np.reshape(np.array(inputs).astype(np.int64), shape)
    return paddle.to_tensor(inputs)


# 定义特征保存函数
def get_usr_mov_features(model: Model, params_file_path, poster_path):
    usr_pkl = {}
    mov_pkl = {}

    # 加载模型参数到模型中，设置为验证模式eval（）
    model_state_dict = load_params(params_file_path)
    model.load_dict(model_state_dict)
    model.eval()
    # 获得整个数据集的数据
    dataset = model.Dataset.dataset

    for i in range(len(dataset)):
        # 获得用户数据，电影数据，评分数据
        # 本案例只转换所有在样本中出现过的user和movie，实际中可以使用业务系统中的全量数据
        usr_info, mov_info, score = dataset[i]['usr_info'], dataset[i]['mov_info'], dataset[i]['scores']
        usrid = str(usr_info['usr_id'])
        movid = str(mov_info['mov_id'])

        # 获得用户数据，计算得到用户特征，保存在usr_pkl字典中
        if usrid not in usr_pkl.keys():
            usr_id_v = list2tensor(usr_info['usr_id'], [1])
            usr_age_v = list2tensor(usr_info['age'], [1])
            usr_gender_v = list2tensor(usr_info['gender'], [1])
            usr_job_v = list2tensor(usr_info['job'], [1])

            usr_in = [usr_id_v, usr_gender_v, usr_age_v, usr_job_v]
            usr_feat = model.get_usr_feat(usr_in)
            usr_pkl[usrid] = usr_feat.numpy()

        # 获得电影数据，计算得到电影特征，保存在mov_pkl字典中
        if movid not in mov_pkl.keys():
            mov_id_v = list2tensor(mov_info['mov_id'], [1])
            mov_tit_v = list2tensor(mov_info['title'], [1, 1, 15])
            mov_cat_v = list2tensor(mov_info['category'], [1, 6])
            mov_in = [mov_id_v, mov_cat_v, mov_tit_v, None]
            mov_feat = model.get_mov_feat(mov_in)
            mov_pkl[movid] = mov_feat.numpy()

    # 保存特征到本地
    pickle_dump(usr_pkl, './usr_feat.pkl')
    pickle_dump(mov_pkl, './mov_feat.pkl')
    print("usr & mov features saved!!!")


if __name__ == "__main__":
    paddle.set_device('cpu')
    fc_sizes = [128, 64, 32]
    use_poster, use_mov_title, use_mov_cat, use_age_job = False, True, True, True
    model = Model(use_poster, use_mov_title, use_mov_cat, use_age_job, fc_sizes)

    param_path = "./checkpoint/epoch9.pdparams"
    poster_path = "./data/posters/"
    get_usr_mov_features(model, param_path, poster_path)
