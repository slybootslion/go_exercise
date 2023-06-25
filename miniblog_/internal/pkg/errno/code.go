package errno

var (
	OK                  = &Errno{}
	InternalServerError = &Errno{HTTP: 500, Code: "InternalError", Message: "Internal server error."}
	ErrPageNotFond      = &Errno{HTTP: 404, Code: "ResourceNotFound.PageNotFound", Message: "Page not found."}
)
