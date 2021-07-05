package gplayapi

import (
	"io"
	"net/http"
	"strings"

	"github.com/Juby210/gplayapi-go/gpproto"
	"google.golang.org/protobuf/proto"
)

func ptrBool(b bool) *bool {
	return &b
}

func ptrStr(str string) *string {
	return &str
}

func ptrInt32(i int32) *int32 {
	return &i
}

func doReq(r *http.Request) ([]byte, error) {
	res, err := httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func parseResponse(res string) map[string]string {
	ret := map[string]string{}
	for _, ln := range strings.Split(res, "\n") {
		keyVal := strings.SplitN(ln, "=", 2)
		if len(keyVal) >= 2 {
			ret[keyVal[0]] = keyVal[1]
		}
	}
	return ret
}

func (client *GooglePlayClient) doAuthedReq(r *http.Request) (_ *gpproto.Payload, err error) {
	client.setDefaultHeaders(r)
	b, err := doReq(r)
	if err != nil {
		return
	}
	resp := &gpproto.ResponseWrapper{}
	err = proto.Unmarshal(b, resp)
	if err != nil {
		return
	}
	return resp.Payload, nil
}
