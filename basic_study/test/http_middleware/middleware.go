package thridparty

//
// import (
// 	"encoding/json"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"
// 	"time"
//
// 	"ptapp.cn/util/itime"
// )
//
// type HttpDo func(req *http.Request) (*http.Response, error)
//
// type HttpClientMiddleware func(HttpDo) HttpDo
//
// func (c *Client) HttpClientMiddlewareInit(HttpDo) HttpDo {
// 	do := c.Client.Do
// 	do = c.PreRequest()(do)
// 	do = c.Collector()(do)
// 	return do
// }
//
// func (c *Client) Collector() HttpClientMiddleware {
// 	return func(next HttpDo) HttpDo {
// 		return func(req *http.Request) (resp *http.Response, err error) {
// 			defer func(begin time.Time) {
// 				lvs := []string{"path", req.URL.Path, "code", strconv.FormatInt(int64(resp.StatusCode), 10)}
// 				c.Metrics.ReqCounter.With(lvs...).Add(1)
// 				c.Metrics.ReqLatency.With(lvs...).Observe(itime.TransformMilliSecond(time.Since(begin)))
// 			}(time.Now())
// 			resp, err = next(req)
// 			if err != nil {
// 				return
// 			}
// 			return
// 		}
// 	}
// }
//
// func (c *Client) PreRequest() HttpClientMiddleware {
// 	return func(next HttpDo) HttpDo {
// 		return func(req *http.Request) (resp *http.Response, err error) {
// 			// Set default headers
// 			req.Header.Set("Accept", "application/json")
// 			/*	req.Header.Set("Accept-Language", "en_US")*/
//
// 			// Default values for headers
// 			if req.Header.Get("Content-type") == "" {
// 				req.Header.Set("Content-type", "application/json")
// 			}
// 			resp, err = next(req)
// 			if err != nil {
// 				return
// 			}
// 			return
// 		}
// 	}
// }
//
// func (c *Client) PostResponse(v interface{}) HttpClientMiddleware {
// 	return func(next HttpDo) HttpDo {
// 		return func(req *http.Request) (resp *http.Response, err error) {
// 			resp, err = next(req)
// 			if err != nil {
// 				return
// 			}
// 			var data []byte
// 			if resp.StatusCode < 200 || resp.StatusCode > 299 {
// 				errResp := &ErrorResponse{Response: resp}
// 				data, err = ioutil.ReadAll(resp.Body)
//
// 				if err == nil && len(data) > 0 {
// 					SetBody(errResp, string(data))
// 					_ = json.Unmarshal(data, errResp)
// 				}
//
// 				return nil, errResp
// 			}
// 			if v == nil {
// 				return
// 			}
//
// 			if w, ok := v.(io.Writer); ok {
// 				_, _ = io.Copy(w, resp.Body)
// 				return
// 			}
// 			respBytes, err := ioutil.ReadAll(resp.Body)
// 			if err != nil {
// 				return nil, err
// 			}
// 			err = json.Unmarshal(respBytes, v)
// 			if err != nil {
// 				return nil, err
// 			}
// 			if w, ok := v.(AikucunWriter); ok {
// 				SetBody(w, string(respBytes))
// 			}
// 			return
// 		}
// 	}
// }
