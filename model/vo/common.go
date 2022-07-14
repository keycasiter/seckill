package vo

const (
	//成功
	SuccessCode = "0000"
	SuccessMsg  = "成功"
	//失败
	FailCode = "0001"
	FailMsg  = "失败"
	//请求参数错误
	ParamErrCode = "0002"
	ParamErrMsg  = "请求参数错误"
)

type CommonResult struct {
	//返回码
	Code string
	//描述
	Msg string
}

func BuildCommonResult(code string, msg string) *CommonResult {
	return &CommonResult{
		Code: code,
		Msg:  msg,
	}
}

func BuildSuccessResult() *CommonResult {
	return &CommonResult{
		Code: SuccessCode,
		Msg:  SuccessMsg,
	}
}

func BuildFailResult() *CommonResult {
	return &CommonResult{
		Code: FailCode,
		Msg:  FailMsg,
	}
}

func BuildParamErrResult() *CommonResult {
	return &CommonResult{
		Code: ParamErrCode,
		Msg:  ParamErrMsg,
	}
}
