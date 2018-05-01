# golang json decode 异常结构(http)兼容操作

> Keywords: golang; json; mutli-type field; Best practice;

## 背景

在 web 开发中，经常遇到第三方服务提供的接口返回的数据格式没有严格按照协议的要求。

## Exception Examples

1. json 里的数字，有时会返回字符串
2. json 里的数组为空时，可能会返回空字符串、数组、对象等，导致解析失败。

## 解码方式

```go
    structObj := targetStruct{}
    ajson.DecodeObject(data, &WrapStruct{}, &structObj)


```

## 解码要求

1. 正常定义数据结构，如 type S struct {}
2. 定义数据 Wrap，如 type W struct{}
3. 定义 Wrap 时, 对于不确定的字段类型使用 json.RawMessage 或者 json.Number 定义
4. 当使用了 json.RawMessage 类型时， Wrap 需要实现接口 WrapRawMessage
5. 当使用了 json.Number 类型时， Wrap 需要实现接口 WrapNumber
6. 在实现接口的 Decode 方法里使用 type switch + structs.Field.Name() 来判断字段与自定义逻辑
7. 检查是否完成。
