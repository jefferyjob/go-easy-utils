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