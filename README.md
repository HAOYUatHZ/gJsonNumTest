# json 精度损失调研报告

网上搜不到 关于struct 定义 `json.....` 还能再加数据类型的描述

搜索 JSON 数值精度损失的 只能找到 [这种](https://ethancai.github.io/2016/06/23/bad-parts-about-json-serialization-in-Golang/#2%E5%8F%B7%E5%9D%91%EF%BC%9AJSON%E5%8F%8D%E5%BA%8F%E5%88%97%E5%8C%96%E6%88%90interface-%E5%AF%B9Number%E7%9A%84%E5%A4%84%E7%90%86)：

```
const jsonStream = `{"name":"ethancai", "fansCount": 9223372036854775807}`
json.Unmarshal([]byte(jsonStream), &user)
```

如果没有定义 结构体 直接 `unmarshal()` 会变 float ，定义了 结构体 则不会
```
type User struct {
    Name      string
    FansCount int64
}
```

保险起见还是写了这个服务穷举 `uint64(0)` ~ `^uint64(0)` 测试
把一个数包进 json 发过去，再返回来，解析完看是否不变
