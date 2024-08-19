# 版本更新记录

## v2.2.4
- 单元测试覆盖率提升 [#81](https://github.com/jefferyjob/go-easy-utils/pull/81)
- Dependabot 计划间隔每周 [#82](https://github.com/jefferyjob/go-easy-utils/pull/82)
- 功能：更新问题模板 [#83](https://github.com/jefferyjob/go-easy-utils/pull/83)
- 更新问题表单 [#84](https://github.com/jefferyjob/go-easy-utils/pull/84)

## v2.2.3
- `jsonUtil` 优化 var 定义 [#80](https://github.com/jefferyjob/go-easy-utils/pull/80)

## v2.2.2
- 增加集合方法：差集、对称差集、交集 [#77](https://github.com/jefferyjob/go-easy-utils/pull/77)
- 补充文档 [#78](https://github.com/jefferyjob/go-easy-utils/pull/78)
- 优化 var 定义 [#79](https://github.com/jefferyjob/go-easy-utils/pull/79)

## v2.2.1
- 优化文档中已知问题 [#76](https://github.com/jefferyjob/go-easy-utils/pull/76)

## v2.2.0
- 新增方法 `ExtractKeys` 用于切片字段提取 [#74](https://github.com/jefferyjob/go-easy-utils/pull/74)
- 新增方法 `SliceToMap` 将切片转换为 map [#74](https://github.com/jefferyjob/go-easy-utils/pull/74)

## v2.1.3
- `error code` 迁移到各自的包下 [#67](https://github.com/jefferyjob/go-easy-utils/pull/67)
- 增加 `cryptoUtil` 包的单元测试方法 [#68](https://github.com/jefferyjob/go-easy-utils/pull/68)

## v2.1.2
- `cryptoUtil` 包支持 RAS 加密方法 [#62](https://github.com/jefferyjob/go-easy-utils/pull/62)
- `mathUtil` 包单元测试优化 [#63](https://github.com/jefferyjob/go-easy-utils/pull/63)

## v2.1.1
- 优化 Makefile 帮助信息展示 [#53](https://github.com/jefferyjob/go-easy-utils/pull/53)
- 增加单元测试代码 [#52](https://github.com/jefferyjob/go-easy-utils/pull/52)

## v2.1.0
- `JsonToStruct` 支持基本数据类型中定义 `interface{}` [#41](https://github.com/jefferyjob/go-easy-utils/pull/41)
- `JsonToStruct` 支持基本数据类型中定义 `any` [#41](https://github.com/jefferyjob/go-easy-utils/pull/41)
- 优化 `parsePrimitiveValue` 的 `string` 和 `bool` 转义兼容 [#43](https://github.com/jefferyjob/go-easy-utils/pull/43)

## v2.0.1
- 修复 `JsonToStruct` 中 `slice` 数据类型不一致解析错误 [#37](https://github.com/jefferyjob/go-easy-utils/pull/37)
- `JsonToStruct` 支持解析 `map` 数据类型 [#37](https://github.com/jefferyjob/go-easy-utils/pull/37)

## v2.0.0
- 项目全面支持泛型和 `any` [#19](https://github.com/jefferyjob/go-easy-utils/pull/19)
- 新增 `mathUtil` 包 [#22](https://github.com/jefferyjob/go-easy-utils/pull/22)
- 删除 `intUtil` 包 [#27](https://github.com/jefferyjob/go-easy-utils/pull/27)
- 优化 `mathUtil` 包泛型 [#30](https://github.com/jefferyjob/go-easy-utils/pull/30)

## v1.2.0
- `JsonToStruct` 支持 `interface{}` 和 `any` 在基本数据类型中的定义 [#40](https://github.com/jefferyjob/go-easy-utils/pull/40)
- 优化 `parsePrimitiveValue` 中 `string` 和 `bool` 的转义兼容 [#42](https://github.com/jefferyjob/go-easy-utils/pull/42)

## v1.1.1
- 修复 `JsonToStruct` 中 `slice` 数据类型不一致解析错误 [#36](https://github.com/jefferyjob/go-easy-utils/pull/36)
- `JsonToStruct` 支持 `map` 数据类型的解析 [#36](https://github.com/jefferyjob/go-easy-utils/pull/36)

## v1.1.0
- `sliceUtil` 包增加更多方法 [#11](https://github.com/jefferyjob/go-easy-utils/pull/11)
- 形参命名更为规范 [#10](https://github.com/jefferyjob/go-easy-utils/pull/10)

## v1.0.2
- 优化 `jsonUtil` 包中 result 类型的验证 [#9](https://github.com/jefferyjob/go-easy-utils/pull/9)
- 修复 `JsonToStruct` 数据类型为空字符串时的校验错误 [#9](https://github.com/jefferyjob/go-easy-utils/pull/9)

## v1.0.1
- 优化 `JsonToStruct` 方法中指针类型参数判断
- 优化 `anyUtil` 包中指针类型的判断
- 优化 `anyUtil` 和 `jsonUtil` 的文档
- 修复 `AnyToInt` 方法中的 debug 代码 [#5](https://github.com/jefferyjob/go-easy-utils/pull/5)
- 修复身份证号验证问题 [#7](https://github.com/jefferyjob/go-easy-utils/pull/7)

## v1.0.0
- 工具包开发，支持常用功能开发