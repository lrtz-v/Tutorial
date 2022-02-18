import paddle
import numpy as np
import matplotlib
import matplotlib.pyplot as plt
import pandas as pd
import seaborn as sns
import warnings
warnings.filterwarnings("ignore")

print(paddle.__version__)

"""
数据集
wget https://archive.ics.uci.edu/ml/machine-learning-databases/housing/housing.data -O housing.data

信息详情
https://www.paddlepaddle.org.cn/documentation/docs/zh/practices/quick_start/linear_regression.html
"""

# 从文件导入数据
datafile = './housing.data'
housing_data = np.fromfile(datafile, sep=' ')
feature_names = ['CRIM', 'ZN', 'INDUS', 'CHAS', 'NOX', 'RM', 'AGE', 'DIS', 'RAD', 'TAX', 'PTRATIO', 'B', 'LSTAT', 'MEDV']
feature_num = len(feature_names)
# 将原始数据进行Reshape，变成[N, 14]这样的形状
housing_data = housing_data.reshape([housing_data.shape[0] // feature_num, feature_num])


def check():
    # 画图看特征间的关系,主要是变量两两之间的关系（线性或非线性，有无明显较为相关关系）
    features_np = np.array([x[:13] for x in housing_data], np.float32)
    labels_np = np.array([x[-1] for x in housing_data], np.float32)
    data_np = np.c_[features_np, labels_np]
    df = pd.DataFrame(housing_data, columns=feature_names)
    matplotlib.use('TkAgg')
    sns.pairplot(df.dropna(), y_vars=feature_names[-1], x_vars=feature_names[::-1], diag_kind='kde')
    # plt.show()

    # 相关性分析
    fig, ax = plt.subplots(figsize=(15, 1))
    corr_data = df.corr().iloc[-1]
    corr_data = np.asarray(corr_data).reshape(1, 14)
    ax = sns.heatmap(corr_data, cbar=True, annot=True)
    plt.show()


"""
归一化
1.过大或过小的数值范围会导致计算时的浮点上溢或下溢。
2.不同的数值范围会导致不同属性对模型的重要性不同（至少在训练的初始阶段如此），而这个隐含的假设常常是不合理的。这会对优化的过程造成困难，使训练时间大大的加长
"""
features_max = housing_data.max(axis=0)
features_min = housing_data.min(axis=0)
features_avg = housing_data.sum(axis=0) / housing_data.shape[0]

BATCH_SIZE = 20


def feature_norm(input):
    f_size = input.shape
    output_features = np.zeros(f_size, np.float32)
    for batch_id in range(f_size[0]):
        for index in range(13):
            output_features[batch_id][index] = (input[batch_id][index] - features_avg[index]) / (features_max[index] - features_min[index])
    return output_features


# 只对属性进行归一化
housing_features = feature_norm(housing_data[:, :13])
# print(feature_trian.shape)
housing_data = np.c_[housing_features, housing_data[:, -1]].astype(np.float32)
# print(training_data[0])

# 将训练数据集和测试数据集按照8:2的比例分开
ratio = 0.8
offset = int(housing_data.shape[0] * ratio)
train_data = housing_data[:offset]
test_data = housing_data[offset:]

"""
模型组网
"""


class Regressor(paddle.nn.Layer):
    def __init__(self):
        super(Regressor, self).__init__()
        self.fc = paddle.nn.Linear(13, 1,)

    def forward(self, inputs):
        pred = self.fc(inputs)
        return pred


train_nums = []
train_costs = []


def draw_train_process(iters, train_costs):
    plt.title("training cost", fontsize=24)
    plt.xlabel("iter", fontsize=14)
    plt.ylabel("cost", fontsize=14)
    plt.plot(iters, train_costs, color='red', label='training cost')
    plt.show()


"""
模型训练&预测
"""
paddle.set_default_dtype("float64")
# step1:用高层API定义数据集，无需进行数据处理等，高层API为你一条龙搞定
train_dataset = paddle.text.datasets.UCIHousing(mode='train')
eval_dataset = paddle.text.datasets.UCIHousing(mode='test')

# step2:定义模型
class UCIHousing(paddle.nn.Layer):
    def __init__(self):
        super(UCIHousing, self).__init__()
        self.fc = paddle.nn.Linear(13, 1, None)

    def forward(self, input):
        pred = self.fc(input)
        return pred

# step3:训练模型
model = paddle.Model(UCIHousing())
model.prepare(paddle.optimizer.Adam(parameters=model.parameters()),
              paddle.nn.MSELoss())
model.fit(train_dataset, eval_dataset, epochs=5, batch_size=8, verbose=1)
