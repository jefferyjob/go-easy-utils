# 版本更新记录

## v1.0.0
- 工具包开发，支持常用功能开发

## v1.0.1
- 优化了 `JsonToStruct` 方法中指针类型的参数判断
- 优化了 `anyUtil` 包中关于指针类型的判断
- 优化了 `anyUtil` 和 `jsonUtil` 文档
- Fix：删除了 `AnyToInt` 方法中的debug代码 [#5](https://github.com/jefferyjob/go-easy-utils/pull/5)
- Fix：修复了身份证号验证存在的问题 [#7](https://github.com/jefferyjob/go-easy-utils/pull/7)

## v1.0.2
- 优化了 `jsonUtil` 包中result类型的验证 [#9](https://github.com/jefferyjob/go-easy-utils/pull/9)
- 修复了`JsonToStruct`中数据类型传入空字符串的校验错误 [#9](https://github.com/jefferyjob/go-easy-utils/pull/9)

## v1.1.0
- `sliceUtil` 包中增加了更多的方法 [#11](https://github.com/jefferyjob/go-easy-utils/pull/11)
- 项目中的某些形参更加规范命名 [#10](https://github.com/jefferyjob/go-easy-utils/pull/10)

## v1.1.1
- 修复了`JsonToStruct`中`slice`中数据类型不一致解析错误 [#36](https://github.com/jefferyjob/go-easy-utils/pull/36)
- `JsonToStruct`支持`map`数据类型的解析 [#36](https://github.com/jefferyjob/go-easy-utils/pull/36)

## v1.2.0
- `JsonToStruct` 支持基本数据类型中定义 `interface{}` [#40](https://github.com/jefferyjob/go-easy-utils/pull/40)
- `JsonToStruct` 支持基本数据类型中定义 `any` [#40](https://github.com/jefferyjob/go-easy-utils/pull/40)
- 优化 `parsePrimitiveValue` 的 `string` 转义兼容 [#42](https://github.com/jefferyjob/go-easy-utils/pull/42)
- 优化 `parsePrimitiveValue` 的 `bool` 转义兼容 [#42](https://github.com/jefferyjob/go-easy-utils/pull/42)

## v2.0.0
- 项目全面支持泛型和any [#19](https://github.com/jefferyjob/go-easy-utils/pull/19)
- 新增 `mathUtil` 包 [#22](https://github.com/jefferyjob/go-easy-utils/pull/22)
- 删除了 `intUtil` 包 [#27](https://github.com/jefferyjob/go-easy-utils/pull/27)
- `mathUtil` 包泛型优化 [#30](https://github.com/jefferyjob/go-easy-utils/pull/30)

## v2.0.1
- 修复了`JsonToStruct`中`slice`中数据类型不一致解析错误 [#37](https://github.com/jefferyjob/go-easy-utils/pull/37)
- `JsonToStruct`支持`map`数据类型的解析 [#37](https://github.com/jefferyjob/go-easy-utils/pull/37)

## v2.1.0
- `JsonToStruct` 支持基本数据类型中定义 `interface{}` [#41](https://github.com/jefferyjob/go-easy-utils/pull/41)
- `JsonToStruct` 支持基本数据类型中定义 `any` [#41](https://github.com/jefferyjob/go-easy-utils/pull/41)
- 优化 `parsePrimitiveValue` 的 `string` 转义兼容 [#43](https://github.com/jefferyjob/go-easy-utils/pull/43)
- 优化 `parsePrimitiveValue` 的 `bool` 转义兼容 [#43](https://github.com/jefferyjob/go-easy-utils/pull/43)

## v2.1.1
- 优化 makefile 文件帮助信息的展示 [#53](https://github.com/jefferyjob/go-easy-utils/pull/53)
- 增加单元测试代码 [#52](https://github.com/jefferyjob/go-easy-utils/pull/52)

## v2.1.2
- `cryptoUtil` 包支持RAS加密方法 [#62](https://github.com/jefferyjob/go-easy-utils/pull/62)
- `mathUtil` 包单元测试优化 [#63](https://github.com/jefferyjob/go-easy-utils/pull/63)

## v2.1.3
- error code 迁移到各自的包下 [#67](https://github.com/jefferyjob/go-easy-utils/pull/67)
- 增加 cryptoUtil 包的单元测试方法 [#68](https://github.com/jefferyjob/go-easy-utils/pull/68)

## v2.2.0
- 开发新方法 ExtractKeys 切片字段提取 [#74](https://github.com/jefferyjob/go-easy-utils/pull/74)
- 开发新方法 SliceToMap 切片转map [#74](https://github.com/jefferyjob/go-easy-utils/pull/74)

## v2.2.1
- 优化文档中存在的已知问题 [#76](https://github.com/jefferyjob/go-easy-utils/pull/76)

## v2.2.2
- feat: 增加了集合方法 差集、对称差集、交集 [#77](https://github.com/jefferyjob/go-easy-utils/pull/77)
- feat: 文档补充  [#78](https://github.com/jefferyjob/go-easy-utils/pull/78)
- 优化var定义 in [#79](https://github.com/jefferyjob/go-easy-utils/pull/79)

## v2.2.3
- jsonUtil 优化var定义 [#80](https://github.com/jefferyjob/go-easy-utils/pull/80)

## v2.2.4
- 单元测试覆盖率提升 #81](https://github.com/jefferyjob/go-easy-utils/pull/81)
- Dependabot 计划间隔每周 [#82](https://github.com/jefferyjob/go-easy-utils/pull/82)