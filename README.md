# golang json decode 异常结构(http)兼容操作

> Keywords: golang; json; mutli-type field; Best practice;

## 解码入口

```go

    1. collect, err := biz.DecodeCollect(line)
    2. collect := biz.Collect{}
    2.1 err := collect.Decode(line)
    2.2 err := ajson.DecodeObject(line, &biz.WrapCollect{}, &collect)

    // 或者直接使用方法:
    ajson.DecodeObject(data, &Wrap{}, &container)

    // 调用方可以有2种方式进行解密, 目前才去方法2

    // 1, 传入待解码的数据， 获取返回对应类型
    // 优势， 调用时非常简单， 一行代码加异常处理就可以
    // 劣势， 要么只能获取指针，要么只能获取变量
    func DecodeCollect(data []byte) (Collect, error) {
        container := Collect{}
        err := ajson.DecodeObject(data, &WrapCollect{}, &container)
        return container, err
    }

    // 2, 调用方先自行初始化类型，再传入待解码的数据，不需要获取返回类型
    // 优势， 获取指针与变量都可以拿到
    // 劣势， 调用时需要2个步骤，初始化+调用解码方法
    func (container *Collect) Decode(data []byte) error {
        return ajson.DecodeObject(data, &WrapCollect{}, container)
    }

```

## 解码要求

```go
    // 可能需要实现2个接口,2个方法
    DecodeRawMessage
    DecodeNumber

    // 有异常的结构体则需要写一个对应的异常的 Wrap
```

## Examples

```go
    // Decode Collect
    func (container *Collect) Decode(data []byte) error {
        return ajson.DecodeObject(data, &WrapCollect{}, container)
    }

    // Decode Ques
    func (container *Ques) Decode(data []byte) error {
        return ajson.DecodeObject(data, &WrapQues{}, container)
    }
```
