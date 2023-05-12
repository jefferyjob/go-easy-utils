# 版本更新记录

## v1.0.0
- 工具包开发，支持常用功能开发

## v1.0.1
- 优化了 `JsonToStruct` 方法中指针类型的参数判断
- 优化了 `anyUtil` 包中关于指针类型的判断
- 优化了 `anyUtil` 和 `jsonUtil` 文档
- Fix：删除了 `AnyToInt` 方法中的debug代码 #5
- Fix：修复了身份证号验证存在的问题 #7

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
