把文本行顺序随机打乱，用 crypto/rand 做的洗牌，比 math/rand 更难被预测。

用法：
  cat 名单.txt | go-shuffle