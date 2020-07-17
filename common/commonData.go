package common
import "net/http"

//http请求
type HttpRequest struct {
	httpRequest *http.Request
	depth uint32
}

//获取一个新的请求
func NewRequest(hr *http.Request , depth uint32) *HttpRequest {
	return &HttpRequest{hr,depth}
}

func (req *HttpRequest) GetHttpRequest() *http.Request {
	return req.httpRequest
}

func (req *HttpRequest) GetDepth() uint32 {
	return req.depth
}
//判断http请求是否有效
func (req *HttpRequest) IsValid() bool {
	return req.httpRequest !=nil && req.httpRequest.URL != nil
}

//http响应
type HttpResponse struct {
	httpResponse *http.Response
	depth uint32
}
//获取一个新的http响应
func NewResponse(hr *http.Response,depth uint32) *HttpResponse {
	return &HttpResponse{hr,depth}
}

func (resp *HttpResponse) GetHttpResponse() *http.Response {
	return resp.httpResponse
}

func (resp *HttpResponse) GetDepth() uint32 {
	return resp.depth
}

func (resp *HttpResponse) IsValid() bool {
	return resp.httpResponse != nil && resp.httpResponse.Body != nil
}

type Item map[string] interface {}
func (item Item) IsValid() bool {
	return item !=nil
}

