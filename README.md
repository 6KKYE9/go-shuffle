把文本行顺序随机打乱，用 crypto/rand 做的洗牌，比 math/rand 更难被预测。

用法：
  cat 名单.txt | go-shuffle
  cat 名单.txt | go-shuffle -n 3            # 只输出洗牌后的前 3 行
  cat 名单.txt | go-shuffle -seed 123        # 固定种子，结果可复现
  cat 名单.txt | go-shuffle -reverse         # 洗牌后再倒序
  echo "a,b,c" | go-shuffle -delim ,        # 按逗号分隔来洗牌
  cat 名单.txt | go-shuffle -o out.txt      # 写到文件

参数：
  -n N       只输出前 N 行（0 表示全部）
  -seed 数字 用固定种子洗牌，同一数字出同一顺序
  -reverse   洗牌后把顺序反过来
  -delim 串  输入按什么分隔，默认按行；可设空格、逗号等
  -o 文件    结果写到文件，不给就打到标准输出

测试：
  go test
