from math import sqrt
import paddle
import numpy as np
from model import Model


def evaluation(model: Model, params_file_path):
    model_state_dict = paddle.load(params_file_path)
    model.load_dict(model_state_dict)
    model.eval()

    acc_set = []
    avg_loss_set = []
    squaredError = []
    for idx, data in enumerate(model.valid_loader()):
        usr, mov, score_label = data
        usr_v = [paddle.to_tensor(var) for var in usr]
        mov_v = [paddle.to_tensor(var) for var in mov]

        _, _, scores_predict = model(usr_v, mov_v)

        pred_scores = scores_predict.numpy()

        avg_loss_set.append(np.mean(np.abs(pred_scores - score_label)))
        squaredError.extend(np.abs(pred_scores - score_label)**2)

        diff = np.abs(pred_scores - score_label)
        diff[diff > 0.5] = 1
        acc = 1 - np.mean(diff)
        acc_set.append(acc)
    RMSE = sqrt(np.sum(squaredError) / len(squaredError))
    # print("RMSE = ", sqrt(np.sum(squaredError) / len(squaredError)))#均方根误差RMSE
    return np.mean(acc_set), np.mean(avg_loss_set), RMSE


if __name__ == "__main__":
    param_path = "./checkpoint/epoch"

    fc_sizes=[128, 64, 32]
    use_poster, use_mov_title, use_mov_cat, use_age_job = False, True, True, True
    model = Model(use_poster, use_mov_title, use_mov_cat, use_age_job,fc_sizes)
    for i in range(10):
        acc, mae,RMSE = evaluation(model, param_path+str(i)+'.pdparams')
        print("ACC:", acc, "MAE:", mae,'RMSE:',RMSE)
