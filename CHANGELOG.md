# Version update record

## v1.0.0
- Toolkit development, supporting common function development

## v1.0.1
- Optimized the parameter judgment of pointer type in `JsonToStruct` method
- Optimized the judgment of pointer type in `anyUti`l package
- Optimized the documentation of `anyUtil` and `jsonUtil`
- Fix: Removed debug code from the `AnyToInt` method #5
- Fix: Problems in ID number verification #7

## v1.0.2
- Optimized the validation of the result type in the `jsonUtil` package
- Fixed the validation error of passing an empty string for the data type in `JsonToStruct`

## v1.1.0
- `sliceUtil` More methods have been added to the package
- Some formal parameters in the project are more standardized in naming

## v1.1.1
- Fixed the data type inconsistency parsing error in `slice` in `JsonToStruct` (#36)
- `JsonToStruct` supports parsing of `map` data type (#36)

## v1.2.0
- `JsonToStruct` supports `interface{}` defined in basic data types
- `JsonToStruct` supports `any` defined in primitive data types