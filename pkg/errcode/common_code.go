package errcode

var (
	Success                  = NewError(0, "Success")
	ServerError              = NewError(10000000, "Internal server error")
	InvalidParams            = NewError(10000001, "Invalid parameters")
	NotFound                 = NewError(10000002, "Not found")
	UnauthorizedAuthNotExist = NewError(10000003, "Authoriz failed with no matching AppKey and AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "Authoriz failed due to bad token")
	UnauthorizedTimeout       = NewError(10000005, "Authoriz failed with token due")
	UnauthorizedTokenGenerate = NewError(10000006, "Authoriz failed with broken token generation")
	TooManyRequests           = NewError(10000007, "To many requests")
)
