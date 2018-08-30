package errno

var (
	// 常见错误
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "内部服务器错误"}

	// 用户错误类
	ErrEncrypt             = &Errno{Code: 20101, Message: "在加密用户密码时发生错误"}
	Err_User_Not_Found     = &Errno{Code: 20102, Message: "没有找到用户"}
	Err_Password_Incorrect = &Errno{Code: 20103, Message: "你输入的密码不正确，请重新输入"}

	//token错误类
	Err_AUTH_CHECK_TOKEN_FAIL    = &Errno{Code: 20104, Message: "Token鉴权失败"}
	Err_AUTH_CHECK_TOKEN_TIMEOUT = &Errno{Code: 20105, Message: "Token已超时"}
	Err_AUTH_TOKEN               = &Errno{Code: 20106, Message: "Token生成失败"}
	Err_AUTH                     = &Errno{Code: 20107, Message: "Token错误"}

	//标签错误类
	Err_EXIST_TAG       = &Errno{Code: 20108, Message: "已存在该标签名称"}
	Err_EXIST_TAG_FAIL  = &Errno{Code: 20109, Message: "获取已存在标签失败"}
	Err_NOT_EXIST_TAG   = &Errno{Code: 20110, Message: "该标签不存在"}
	Err_GET_TAGS_FAIL   = &Errno{Code: 20111, Message: "获取所有标签失败"}
	Err_COUNT_TAG_FAIL  = &Errno{Code: 20112, Message: "统计标签失败"}
	Err_ADD_TAG_FAIL    = &Errno{Code: 20113, Message: "新增标签失败"}
	Err_EDIT_TAG_FAIL   = &Errno{Code: 20114, Message: "修改标签失败"}
	Err_DELETE_TAG_FAIL = &Errno{Code: 20115, Message: "删除标签失败"}
	Err_EXPORT_TAG_FAIL = &Errno{Code: 20116, Message: "导出标签失败"}
	Err_IMPORT_TAG_FAIL = &Errno{Code: 20117, Message: "导入标签失败"}

	//文章错误类

	//图片错误类

	//

)

//ERROR_EXIST_TAG:                 "已存在该标签名称",
//ERROR_EXIST_TAG_FAIL:            "获取已存在标签失败",
//ERROR_NOT_EXIST_TAG:             "该标签不存在",
//ERROR_GET_TAGS_FAIL:             "获取所有标签失败",
//ERROR_COUNT_TAG_FAIL:            "统计标签失败",
//ERROR_ADD_TAG_FAIL:              "新增标签失败",
//ERROR_EDIT_TAG_FAIL:             "修改标签失败",
//ERROR_DELETE_TAG_FAIL:           "删除标签失败",
//ERROR_EXPORT_TAG_FAIL:           "导出标签失败",
//ERROR_IMPORT_TAG_FAIL:           "导入标签失败",

//ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
//ERROR_ADD_ARTICLE_FAIL:          "新增文章失败",
//ERROR_DELETE_ARTICLE_FAIL:       "删除文章失败",
//ERROR_CHECK_EXIST_ARTICLE_FAIL:  "检查文章是否存在失败",
//ERROR_EDIT_ARTICLE_FAIL:         "修改文章失败",
//ERROR_COUNT_ARTICLE_FAIL:        "统计文章失败",
//ERROR_GET_ARTICLES_FAIL:         "获取多个文章失败",
//ERROR_GET_ARTICLE_FAIL:          "获取单个文章失败",
//ERROR_GEN_ARTICLE_POSTER_FAIL:   "生成文章海报失败",

//ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
//ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
//ERROR_AUTH_TOKEN:                "Token生成失败",
//ERROR_AUTH:                      "Token错误",
//ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
//ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
//ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
