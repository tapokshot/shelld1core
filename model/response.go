package model

type Response interface {
	Code() int
	Encode() ([]byte, error)
}

//func Error(err error) Response {
//	return Response{
//		Errors: []Error{
//			{Msg: err.Error()},
//		},
//	}
//}
//
//func ErrorsResponse(err []error) Response {
//	errResp := make([]Error, 0, len(err))
//	for _, e := range err {
//		errResp = append(errResp, Error{Msg: e.Error()})
//	}
//	return Response{
//		Errors: errResp,
//	}
//}
