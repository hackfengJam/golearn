package thridparty

//
// import (
// 	"net/http"
// 	"strconv"
// 	"time"
//
// 	"github.com/go-kit/kit/metrics"
// 	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
// 	stdprometheus "github.com/prometheus/client_golang/prometheus"
// 	"ptapp.cn/util/itime"
// )
//
// // ref https://prometheus.io/docs/concepts/metric_types/
// type MetricsCollector struct {
// 	HostName string
//
// 	// 服务调用计数
// 	ReqCounter metrics.Counter
// 	// 服务响应分布
// 	ReqLatency metrics.Histogram
// }
//
// // 定义一个数据收集接口，可以对接不同时序数据库, eg influxdb, prometheus
// type Collector interface {
// 	Collect(method string)
// }
//
// func (c *Client) Collect() HttpClientMiddleware {
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
// func NewMetricsCollector() *MetricsCollector {
// 	reqCounter := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
// 		Name: "request_cnt",
// 		Help: "Number of requests received.",
// 	}, []string{"method", "code", "service", "hostname"})
//
// 	requestLatency := kitprometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
// 		Name:    "request_duration",
// 		Help:    "Total duration of requests . not include writing on wire",
// 		Buckets: []float64{.1, 1, 10, 30, 50, 100, 300, 500, 1000, 3000},
// 	}, []string{"method", "code", "service", "hostname"})
//
// 	return &MetricsCollector{
// 		ReqCounter: reqCounter,
// 		ReqLatency: requestLatency,
// 	}
// }
//
// func metricsXCodeMiddleware(c *gin.Context) {
// 	start := time.Now()
// 	c.Next()
// 	end := time.Now()
// 	code := c.Writer.Header().Get("X-Code")
// 	metrcisCollect(c, code, end.Sub(start))
// 	httpStatusCollect(c)
// }
//
// func metrcisCollect(c *gin.Context, code string, duration time.Duration) {
// 	if code == "" { // code等于"" 说明是没有使用过router定义过的path，比如404 或者 OPTIONS等方法。 如果这些放过来记录会导致wildcard的path造成指标爆炸
// 		return
// 	}
// 	if path, okk := c.Keys[wildcardPath]; okk {
// 		apiCodeCounterVec.WithLabelValues(c.Request.Method, path.(string), code, hostname).Inc()
// 		apiDurationHistogramVec.WithLabelValues(c.Request.Method, path.(string), hostname).Observe(itime.TransformMilliSecond(duration))
// 	} else {
// 		apiCodeCounterVec.WithLabelValues(c.Request.Method, c.Request.URL.Path, code, hostname).Inc()
// 		apiDurationHistogramVec.WithLabelValues(c.Request.Method, c.Request.URL.Path, hostname).Observe(itime.TransformMilliSecond(duration))
// 	}
// }
//
// // 收集http status
// func httpStatusCollect(c *gin.Context) {
// 	httpCodeCounterVec.WithLabelValues(c.Request.Method, strconv.Itoa(c.Writer.Status()), hostname).Inc()
// }
