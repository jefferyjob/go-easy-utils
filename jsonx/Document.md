# JsonToStruct 开发文档

## 处理指针类型

```go
v := reflect.ValueOf(i)
if v.Kind() == reflect.Ptr {
    if v.IsNil() {
        return 0, nil // 检查解引用后的值是否为 nil
    }
    v = v.Elem() // 获取指针所指向的值
}
```

在`reflect`包中，指针类型和其指向的值的类型是不同的。例如，一个指向`int`类型的指针，其类型是`*int`，而其指向的值的类型是`int`。因此，在使用`reflect`包时，我们需要注意区分指针类型和其指向的值的类型，否则可能会出现类型错误等问题。

因此，对于`value`为指针类型的情况，我们需要通过Elem方法获取其指向的值，然后再进行类型转换。而在使用`Elem`方法之前，我们需要先检查v是否为指针类型，因此使用循环语句进行判断。

**举个例子**

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := &Person{"Tom", 18}
    v := reflect.ValueOf(p).Elem()
    name := v.FieldByName("Name")
    age := v.FieldByName("Age")
    
    fmt.Println(name.String()) // 输出：Tom
    fmt.Println(age.Int())     // 输出：18
    
    name.SetString("Jerry")    // 修改 Name 字段的值
    age.SetInt(20)             // 修改 Age 字段的值
    
    fmt.Println(p.Name)        // 输出：Jerry
    fmt.Println(p.Age)         // 输出：20
}
```

在上面的例子中，我们首先将结构体指针 `p` 转换为 `reflect.Value` 类型的值 `v`，然后使用 `v.Elem()` 获取指针所指向的值，接着使用 `FieldByName()` 方法获取结构体字段的 `reflect.Value` 类型的值，最后使用 `SetString()` 和 `SetInt()` 方法修改结构体字段的值。

## 处理Slice数据类型

```go
case reflect.Slice:
    if subData, ok := value.([]interface{}); ok {
        if err := parseSlice(fieldValue, subData); err != nil {
            return err
        }
    } else {
        return fmt.Errorf("unexpected value type for slice: %T", value)
    }
```

将每个子元素都转换为结构体定义的数据类型再赋值给对应的字段。如果子元素是一个切片类型，需要进行递归处理。

首先获取到切片的元素类型 subType，如果 subType 是一个指针类型，需要将其转换为实际的元素类型。然后创建一个与原始切片相同长度的新切片 subSliceValue，并循环遍历原始切片中的每个子元素。在每个子元素中，需要将其转换为结构体定义的数据类型，可以通过递归调用 parseStruct 或 parseSlice 来实现。最后，将新的子元素设置到新切片 subSliceValue 中，并将其赋值给对应的字段。注意，如果元素类型是一个结构体，需要使用 Set 方法设置到新切片中，如果是基本类型，则需要判定类型然后添加到值中。

## 处理Map数据类型

**第一种更加简洁的调用方法**

```go
case reflect.Map:
    if subData, ok := value.(map[string]interface{}); ok {
        subResult := reflect.New(fieldType.Type).Elem()
        if err := parseMap(subResult, subData); err != nil {
            return err
        }
        fieldValue.Set(subResult)
    }
```

其中 subResult 使用 reflect.New 创建一个新的结构体对象，然后使用 Elem() 方法获取其指针所指向的值，最后传递给 parseMap 方法进行解析。这样可以避免对 reflect.MakeMap 的使用，更加简洁。


**第二种实在原有的代码进行更改。**
```go
case reflect.Map:
    if subData, ok := value.(map[string]interface{}); ok {
        subResult := reflect.MakeMap(fieldValue.Type())
        for subKey, subValue := range subData {
            subElem := reflect.New(fieldType.Type.Elem())
            err := parseValue(subElem.Elem(), subValue)
            if err != nil {
                return err
            }
            subResult.SetMapIndex(reflect.ValueOf(subKey), subElem.Elem())
        }
        fieldValue.Set(subResult)
    }
```

其中，subData 是一个 map[string]interface{} 类型的变量，表示一个嵌套的 JSON 对象；subResult 是一个 reflect.Value 类型的变量，表示一个 map 类型的变量的反射值；subKey 和 subValue 分别表示 subData 中的键和值，可以通过 range 循环遍历 subData。

在调用 parseValue 函数时，需要传入 subElem.Elem() 作为第一个参数，这是因为 subElem 表示一个指向 fieldType.Type.Elem() 类型变量的指针，而 parseValue 函数需要传入一个值类型的变量。所以需要通过 subElem.Elem() 获取到 fieldType.Type.Elem() 类型的变量。