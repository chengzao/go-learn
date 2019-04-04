# 切片Slice

- 切片是围绕动态数组的概念构建的，可以按需自动增长和缩小。
切片的动态增长是通过内置函数 append 来实现的。
这个函数可以快速且高效地增长切片。还可以通过对切片再次切片来缩小一个切片的大小。
- 切片有 3 个字段分别是指向底层数组的指针、切片访问的元素的个数（即长度）和切片允许增长到的元素个数（即容量）

## append

- 当切片容量（而非数组长度，默认切片容量等于数组长度，也可以显示指定）用完了，再追加需要扩容，
此处会新建数组，长度为原数组的2倍，然后将旧数组元素拷贝到新数组，newSlice 的底层数组是新数组，
newSlice 容量为8；而 slice 的底层数组是旧数组，二者互不影响
- slice 扩容机制：在切片的容量小于 1000 个元素时，总是会成倍地增加容量。
一旦元素个数超过 1000，容量的增长因子会设为 1.25，也就是会每次增加 25% 的容量

## 显示设置容量

- 在没有显示指定容量的情况下，切片容量就是其底层数组的长度
- 如果在创建切片时设置切片的容量和长度一样，就可以强制让新切片的第一个 append 操作创建新的底层数组，
与原有的底层数组分离。新切片与原有的底层数组分离后，可以安全地进行后续修改

## 函数间传递切片

- 在函数间传递切片就是要在函数间以值的方式传递切片。由于切片的尺寸很小，
在函数间复制和传递切片成本也很低；而在函数间传递数组是需要拷贝整个数组的，所以内存和性能上都不好
- 调用函数，传递一个切片副本，实际上内部还是传递了对数组的指针，所以 foo 内部的操作会影响 main 中的 slice