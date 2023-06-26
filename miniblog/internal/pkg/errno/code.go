package errno

var (
	// OK 代表请求成功
	OK = &Errno{HTTP: 200, Code: "", Message: "ok"}
	// 所有未知服务器错误
	InternalServerError = &Errno{HTTP: 500, Code: "InternalError", Message: "Internal server error."}
	// not found
	ErrPageNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.PageNotFound", Message: "Page not found."}
)
