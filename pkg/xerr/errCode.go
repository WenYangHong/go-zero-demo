package xerr

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006

// 用户模块
const REQISTER_USER_EXIST_ERROR uint32 = 100101  // 注册用户存在
const REGISTER_USER_ERROR uint32 = 100102        // 注册用户错误
const LOGIN_USER_NOT_FOUND_ERROR uint32 = 100103 // 登录用户未找到
const LOGIN_PASSWORD_ERROR uint32 = 100104       // 登录账号密码不匹配
const LOGIN_ERROR uint32 = 100105                // 登录异常
const USER_NOT_FOUND_ERROR = 100106              // 用户信息不存在

// 民宿模块
const COMMENT_CREATE_FAIL = 200001  // 民宿评价失败
const HOMESTAY_DETAIL_FAIL = 200002 // 无民宿详情
