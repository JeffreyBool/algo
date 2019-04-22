# Golang Array 泛型数组

>基于 golang 实现的泛型数组,支持动态扩容等特性。

[![Build Status](https://www.travis-ci.org/JeffreyBool/array.svg?branch=master)](https://www.travis-ci.org/JeffreyBool/array)

## 项目结构
``` 
.
├── README.md
├── array.go  //实现
├── array_test.go  //测试用例
└── example_array_test.go  //使用例子
```

## 功能
- [x] `GetCapacity`获取数组容量
- [x] `GetSize`获取数组长度
- [x] `IsEmpty`判断数组是否为空
- [x] `AddFirst`向数组头插入元素
- [x] `AddLast`向数组尾插入元素
- [x] `Add`在索引位置插入元素
- [x] `Get`获取索引元素
- [x] `Set`修改索引位置元素
- [x] `Contains`查找数组中是否有元素
- [x] `Find`通过索引查找数组，索引范围[0,n-1](未找到，返回 -1)
- [x] `Remove`删除 index 位置的元素，并返回
- [x] `RemoveFirst`删除数组首个元素
- [x] `RemoveLast`删除末尾元素
- [x] `RemoveElement`从数组中删除指定元素
- [x] `Clear`清空数组
- [x] `PrintIn`打印格式化
