"""
知道乘客每次乘坐出租车的公里数, 也知道乘客每次下车的时候支付给出租车司机的总费用。但是并不知道乘车的起步价, 以及每公里行驶费用是多少。希望让机器从这些数据当中学习出来计算总费用的规则

引入paddle
"""
import paddle
print("paddle: " + paddle.__version__)


"""
准备数据

在这个机器学习任务中, 已经知道了乘客的 行驶里程distance_travelled, 和对应的, 这些乘客的 总费用total_fee。
通常情况下, 在机器学习任务中, 像distance_travelled这样的输入值, 一般被称为x (或者特征feature), 像total_fee这样的输出值, 一般被称为y (或者标签label)。
可以用paddle.to_tensor把示例数据转换为paddle的Tensor数据。
"""

x_data = paddle.to_tensor([[1.], [3.0], [5.0], [9.0], [10.0], [20.0]])
y_data = paddle.to_tensor([[12.], [16.0], [20.0], [28.0], [30.0], [50.0]])


"""
用飞桨定义模型的计算

使用飞桨定义模型的计算的过程，本质上，是用python，通过飞桨提供的API，来告诉飞桨计算规则的过程。回顾一下，想要通过飞桨用机器学习方法，从数据当中学习出来如下公式当中的w和b。这样在未来，给定x时就可以估算出来y值（估算出来的y记为y_predict）

    y_predict = w * x + b

将会用飞桨的线性变换层：paddle.nn.Linear来实现这个计算过程，这个公式里的变量x, y, w, b, y_predict，对应着飞桨里面的Tensor概念。
"""

linear = paddle.nn.Linear(in_features=1, out_features=1)


"""
准备好运行飞桨

机器（计算机）在一开始的时候会随便猜w和b，先看看机器猜的怎么样。你应该可以看到，这时候的w是一个随机值，b是0.0，这是飞桨的初始化策略，也是这个领域常用的初始化策略
"""

w_before_opt = linear.weight.numpy().item()
b_before_opt = linear.bias.numpy().item()

print("w before optimize: {}".format(w_before_opt))
print("b before optimize: {}".format(b_before_opt))

"""
告诉飞桨怎么样学习

告诉飞桨，怎么样去学习，从而能得到参数w和b
在机器学习/深度学习当中，机器（计算机）在最开始的时候，得到参数w和b的方式是随便猜一下，用这种随便猜测得到的参数值，去进行计算（预测）的时候，得到的y_predict，跟实际的y值一定是有差距的
接下来，机器会根据这个差距来调整w和b，随着这样的逐步的调整，w和b会越来越正确，y_predict跟y之间的差距也会越来越小，从而最终能得到好用的w和b。这个过程就是机器学习的过程

用更加技术的语言来说，衡量差距的函数（一个公式）就是损失函数，用来调整参数的方法就是优化算法。
在本示例当中，用最简单的均方误差(mean square error)作为损失函数(paddle.nn.MSELoss)；和最常见的优化算法SGD（stocastic gradient descent)作为优化算法（传给paddle.optimizer.SGD的参数learning_rate，你可以理解为控制每次调整的步子大小的参数）。
"""
mse_loss = paddle.nn.MSELoss()
sgd_optimizer = paddle.optimizer.SGD(learning_rate=0.001, parameters = linear.parameters())

"""
运行优化算法

让飞桨运行一下这个优化算法，这会是一个前面介绍过的逐步调整参数的过程，你应该可以看到loss值（衡量y和y_predict的差距的loss)在不断的降低。
"""


total_epoch = 10000
for i in range(total_epoch):
    y_predict = linear(x_data)
    loss = mse_loss(y_predict, y_data)
    loss.backward()
    sgd_optimizer.step()
    sgd_optimizer.clear_grad()

    if i%1000 == 0:
        print("epoch {} loss {}".format(i, loss.numpy()))

print("finished training,  loss {}".format(loss.numpy()))


"""
经过了这样的对参数w和b的调整（学习)，再通过下面的程序，来看看现在的参数变成了多少。
会发现w变成了很接近2.0的一个值，b变成了接近10.0的一个值。
虽然并不是正好的2和10，但却是从数据当中学习出来的还不错的模型的参数，可以在未来的时候，用从这批数据当中学习到的参数来预估了。
（如果你愿意，也可以通过让机器多学习一段时间，从而得到更加接近2.0和10.0的参数值。)
"""

w_after_opt = linear.weight.numpy().item()
b_after_opt = linear.bias.numpy().item()

print("w after optimize: {}".format(w_after_opt))
print("b after optimize: {}".format(b_after_opt))


