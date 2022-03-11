# 数据

- ml-1m 是 GroupLens Research 从 MovieLens 网站上收集并提供的电影评分数据集。包含了 6000 多位用户对近 3900 个电影的共 100 万条评分数据，评分均为 1 ～ 5 的整数，其中每个电影的评分数据至少有 20 条。该数据集包含三个数据文件，分别是：

## users.dat，存储用户属性信息的文本格式文件

- 每一行表示一个用户的数据，以::隔开，第一列到最后一列分别表示 UserID、Gender(F 或 M)、Age、Occupation、Zip-code
  - UserID 每个用户的数字代号 1、2、3 等序号
  - Gender F 表示女性，M 表示男性 F 或 M
  - Age 用数字表示各个年龄段, 1: “Under 18”, 18: “18-24”, 25: “25-34”, 35: “35-44”, 45: “45-49”, 50: “50-55”, 56: “56+”
  - Occupation 用数字表示不同职业, 0: “other” or not specified, 1: “academic/educator”, 2: “artist”, 3: “clerical/admin”, 4: “college/grad student”, 5: “customer service”, 6: “doctor/health care”, 7: “executive/managerial”, 8: “farmer”, 9: “homemaker”, 10: “K-12 student”, 11: “lawyer”, 12: “programmer”, 13: “retired”, 14: “sales/marketing”, 15: “scientist”, 16: “self-employed”, 17: “technician/engineer”, 18: “tradesman/craftsman”, 19: “unemployed”, 20: “writer”
  - zip-code 邮政编码，与用户所处的地理位置有关。在本次实验中，不使用这个数据。 48067

## movies.dat，存储电影属性信息的文本格式文件

- 数据格式为：MovieID::Title::Genres，保存的格式与用户数据相同，每一行表示一条电影数据信息
  - MovieID 每个电影的数字代号 1、2、3 等序号
  - Title 每个电影的名字和首映时间 比如：Toy Story (1995)
  - Genres 电影的种类，每个电影不止一个类别，不同类别以　|　隔开 比如：Animation| Children’s|Comedy, 包含的类别有：【Action，Adventure，Animation，Children’s，Comedy，Crime，Documentary，Drama，Fantasy，Film-Noir，Horror，Musical，Mystery，Romance，Sci-Fi，Thriller，War，Western】

## ratings.dat, 存储电影评分信息的文本格式文件

- 评分数据格式为 UserID::MovieID::Rating::Timestamp
  - UserID用户id
  - MovieID 电影id
  - Rating 评分
  - Timestamp 时间戳（忽略）

## new_rating.txt, 包含海报图像的评分数据文件（从原始评分数据中过滤得到）

## posters, 包含电影海报图像。

- 电影海报图像在 posters 文件夹下，海报图像的名字以"mov_id" + 电影 ID + ".png"的方式命名。由于这里的电影海报图像有缺失，我们整理了一个新的评分数据文件，新的文件中包含的电影均是有海报数据的，因此，本次实验使用的数据集在 ml-1m 基础上增加了两份数据：
