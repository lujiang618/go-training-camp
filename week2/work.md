
# 题目
```
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
```


dao层从数据库内获取数据，一般是被其他层依赖，如果遇到ErrNoRows则dao层不能自己处理err， 需要将error抛给上层。
通过pkg/error包进行wrap，记录error的栈堆信息返回更多的上下文信息（如查询的参数等）， 在上层调用errors.Cause()来获取Sentinel error判断是否是 sql.ErrNoRows，然后作响应的处理。

底层代码不处理error，通过pkg/error包装error抛给上层，由顶层统一处理， 顶层通过`%+v`谓词打印调用栈，通过errors.Cause() 获取原始的error类型。






