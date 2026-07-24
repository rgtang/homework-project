package utils

import "net/http"

type AppError struct {
	Code       int    `json:"code"`    // 业务错误码，也可直接与 HTTP Status 一致
	Message    string `json:"message"` // 对前端友好的提示文案
	StatusCode int    `json:"-"`       // HTTP Status Code (不序列化到 JSON)
	RawErr     error  `json:"-"`       // 原始底层 Error（用于日志排查，不暴露给前端）
}

func (e *AppError) Error() string {
	if e.RawErr != nil {
		return e.RawErr.Error()
	}
	return e.Message
}

// 工厂函数：附加原始错误信息
func (e *AppError) WithErr(err error) *AppError {
	return &AppError{
		Code:       e.Code,
		Message:    e.Message,
		StatusCode: e.StatusCode,
		RawErr:     err,
	}
}

// --- 预定义常见业务错误 ---
var (
	// 400 客户端错误
	ErrInvalidParam = &AppError{Code: 40001, StatusCode: http.StatusBadRequest, Message: "请求参数格式错误或缺失"}
	ErrUserExists   = &AppError{Code: 40002, StatusCode: http.StatusBadRequest, Message: "用户名或邮箱已被注册"}

	// 401 认证错误
	ErrUnauthorized = &AppError{Code: 40101, StatusCode: http.StatusUnauthorized, Message: "用户未认证或凭证已过期"}
	ErrInvalidToken = &AppError{Code: 40102, StatusCode: http.StatusUnauthorized, Message: "无效的 Token 签名"}
	ErrLoginFailed  = &AppError{Code: 40103, StatusCode: http.StatusUnauthorized, Message: "用户名或密码错误"}

	// 403 权限错误
	ErrForbidden = &AppError{Code: 40301, StatusCode: http.StatusForbidden, Message: "无权操作此资源"}

	// 404 资源不存在
	ErrNotFoundPost    = &AppError{Code: 40401, StatusCode: http.StatusNotFound, Message: "请求的文章不存在"}
	ErrNotFoundComment = &AppError{Code: 40402, StatusCode: http.StatusNotFound, Message: "请求的评论不存在"}
	ErrNotFoundUser    = &AppError{Code: 40403, StatusCode: http.StatusNotFound, Message: "未找到指定的账户"}

	// 500 服务器错误
	ErrDatabaseError  = &AppError{Code: 50001, StatusCode: http.StatusInternalServerError, Message: "数据库服务内部错误"}
	ErrInternalServer = &AppError{Code: 50000, StatusCode: http.StatusInternalServerError, Message: "系统内部错误，请联系管理员"}
)
