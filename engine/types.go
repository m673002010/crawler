package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

// 空结构体
func NilParser([]byte) ParserResult {
	return ParserResult{}
}
