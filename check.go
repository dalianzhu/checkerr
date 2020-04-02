package checkerr

import (
    "errors"
    "fmt"
    "strings"
)

/*
checkerr 取代讨厌的 if err!=nil。
你是否已经厌倦无止境的 if err!=nil{return err} ? 使用本包，解放你自己！
使用前，请仔细阅读说明书
!!!! 重要，唯一正确的用法：

func testError() (err error) { // 返回的error必须命名参数
	defer MarkPanic(&err) // 必须直接在defer 调用 MarkPanic，传入err地址

	err = dosth1()
	// 直接check这个err。如果err不为nil，则 testError 会返回err
	CheckError(err)

	err = dosth2()
	// 也可以检查这个err并返回一个自定义的新err
	CheckErrorf(err, "check err %v", err) // 这里注意，不要传入err.Error()，因为err也许为nil

	return nil
}
*/

type InnerError struct {
    InputError error
}

func (e InnerError) Error() string {
    return e.InputError.Error()
}

func newInnerErr(err error) *InnerError {
    innerErr := new(InnerError)
    innerErr.InputError = err
    return innerErr
}

// MarkPanic 处理由check err产生的panic，使用前仔细阅读包注释
func MarkPanic(e *error) {
    if err := recover(); err != nil {
        switch v := err.(type) {
        case InnerError:
            *e = v.InputError
        case *InnerError:
            *e = v.InputError
        default:
            // 其它的panic，继续传递
            panic(err)
        }
    }
}

// CheckError 检查error，必须和MarkPanic配合使用，使用前仔细阅读包注释
func CheckError(err error, replaceErr ...error) {
    if err != nil {
        if len(replaceErr) > 0 {
            var errInfo strings.Builder
            errInfo.Grow(3)
            for _, e := range replaceErr {
                errInfo.WriteString(e.Error())
            }
            panic(newInnerErr(errors.New(errInfo.String())))
        }
        panic(newInnerErr(err))
    }
}

func CheckErrorf(err error, f string, args ...interface{}) {
    if err != nil {
        newErr := fmt.Errorf(f, args...)
        panic(newInnerErr(newErr))
    }
}
