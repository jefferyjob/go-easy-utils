# Version update record

## v1.0.0
- Toolkit development, supporting common function development

## v1.0.1
- Optimized the parameter judgment of pointer type in `JsonToStruct` method
- Optimized the judgment of pointer type in `anyUti`l package
- Optimized the documentation of `anyUtil` and `jsonUtil`
- Fix: Removed debug code from the `AnyToInt` method (#5)
- Fix: Problems in ID number verification (#7)

## v1.0.2
- Optimized the validation of the result type in the `jsonUtil` package (#9)
- Fixed the validation error of passing an empty string for the data type in `JsonToStruct` (#9)

## v1.1.0
- `sliceUtil` More methods have been added to the package (#11)
- Some formal parameters in the project are more standardized in naming (#10)

## v1.1.1
- Fixed the data type inconsistency parsing error in `slice` in `JsonToStruct` (#36)
- `JsonToStruct` supports parsing of `map` data type (#36)

## v1.2.0
- `JsonToStruct` supports `interface{}` defined in basic data types (#40)
- `JsonToStruct` supports `any` defined in primitive data types (#40)
- Optimize `string` escape compatibility of `parsePrimitiveValue` (#42)
- Optimize `parsePrimitiveValue`'s `bool` escape compatibility (#42)

## v2.0.0
- The project fully supports generics and any (#19)
- Added `mathUtil` package (#22)
- Removed `intUtil` package (#27)
- `mathUtil` package generic optimization (#30)

## v2.0.1
- Fixed the data type inconsistency parsing error in `slice` in `JsonToStruct` (#37)
- `JsonToStruct` supports parsing of `map` data type (#37)

## v2.1.0
- `JsonToStruct` supports `interface{}` defined in basic data types (#41)
- `JsonToStruct` supports `any` defined in primitive data types  (#41)
- Optimize `string` escape compatibility of `parsePrimitiveValue`  (#43)
- Optimize `parsePrimitiveValue`'s `bool` escape compatibility  (#43)

## v2.1.1
- Optimize the display of makefile help information (#53)
- Add unit test code (#52)

## v2.1.2
- `cryptoUtil` package supports RAS encryption method (#62)
- `mathUtil` package unit test optimization (#63)