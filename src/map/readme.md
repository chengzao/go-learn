# 映射Map

- Map 是一个存储键值对的无序集合，每次迭代映射的时候顺序也可能不一样
- Slice、Map、function 以及包含切片的结构类型不能作为 Map 的 key
- map在函数间传递，不会拷贝一份map，相当于是"引用传递"，所以remove函数对传入的map的操作是会影响到main函数中的map的
