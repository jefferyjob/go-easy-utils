# 版本更新记录

## v1.0.0
- 工具包开发，支持常用功能开发

## v1.0.1
- 优化了 `JsonToStruct` 方法中指针类型的参数判断
- 优化了 `anyUtil` 包中关于指针类型的判断
- 优化了 `anyUtil` 和 `jsonUtil` 文档
- Fix：删除了 `AnyToInt` 方法中的debug代码
- Fix：修复了身份证号验证存在的问题

## v1.0.2
- 优化了 `jsonUtil` 包中result类型的验证
- 修复了`JsonToStruct`中数据类型传入空字符串的校验错误

## v1.1.0
- `sliceUtil` 包中增加了更多的方法
- 项目中的某些形参更加规范命名

## v1.1.1
- 修复了`JsonToStruct`中`slice`中数据类型不一致解析错误
- `JsonToStruct`支持`map`数据类型的解析

## v1.2.0
- `JsonToStruct` 支持基本数据类型中定义 `interface{}`
- `JsonToStruct` 支持基本数据类型中定义 `any`
- 优化 `parsePrimitiveValue` 的 `string` 转义兼容
- 优化 `parsePrimitiveValue` 的 `bool` 转义兼容

## v2.0.0
- 项目全面支持泛型和any
- 新增 `mathUtil` 包
- 删除了 `intUtil` 包
- `mathUtil` 包泛型优化

## v2.0.1
- 修复了`JsonToStruct`中`slice`中数据类型不一致解析错误
- `JsonToStruct`支持`map`数据类型的解析

## v2.1.0
- `JsonToStruct` 支持基本数据类型中定义 `interface{}`
- `JsonToStruct` 支持基本数据类型中定义 `any`
- 优化 `parsePrimitiveValue` 的 `string` 转义兼容
- 优化 `parsePrimitiveValue` 的 `bool` 转义兼容

## v2.1.1
- 优化 makefile 文件帮助信息的展示
- 增加单元测试代码

## v2.1.2
- `cryptoUtil` 包支持RAS加密方法
- `mathUtil` 包单元测试优化

## v2.1.3
- error code 迁移到各自的包下
- 增加 cryptoUtil 包的单元测试方法

## v2.2.0
- 开发新方法 ExtractKeys 切片字段提取
- 开发新方法 SliceToMap 切片转map

## v2.2.1
- 优化文档中存在的已知问题