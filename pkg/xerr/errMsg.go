package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	message[REQISTER_USER_EXIST_ERROR] = "该手机号已注册"
	message[REGISTER_USER_ERROR] = "账号注册失败"
	message[LOGIN_USER_NOT_FOUND_ERROR] = "账号不存在"
	message[LOGIN_PASSWORD_ERROR] = "密码错误"
	message[LOGIN_ERROR] = "账号登录失败"
	message[COMMENT_CREATE_FAIL] = "民宿评价失败"
	message[ORDER_LIVE_TIME_ERROR] = "订单入住时间错误"
	message[ORDER_LIVE_PEOPLE_NUM_ERROR] = "订单入住人数错误"
	message[HOMESTAY_DETAIL_FAIL] = "民宿信息不存在"
	message[ORDER_CREATE_FAIL] = "订单创建失败"
	message[USER_NOT_FOUND_ERROR] = "用户不存在或未登录"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
