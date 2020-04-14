## checkerr 取代讨厌的 if err!=nil。
你是否已经厌倦无止境的 if err!=nil{return err} ? 使用本包，解放你自己！

使用前，请仔细阅读说明书

!!!! 重要，唯一正确的用法：
```
func testError() (err error) { // 返回的error必须命名参数
	defer MarkPanic(&err) // 必须直接在defer 调用 MarkPanic，传入err地址
	err = dosth1()

	// 直接check这个err。如果err不为nil，则 testError 会返回err
	CheckError(err)
	// 也可使用额外的传参，替换需要返回的error
	// CheckError(err, replaceErr1, replaceErr2, string3)

	err = dosth2()
	// 也可以检查这个err并返回一个自定义的新err
	CheckErrorf(err, "check err %v", err) // 这里注意，不要传入err.Error()，因为err也许为nil
	return nil
}
```
