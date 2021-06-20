package job

import (
	"github.com/junhaideng/51job/internal"
	"github.com/junhaideng/51job/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"go.uber.org/zap"
)

type Job struct {
	Log         *zap.Logger
	Keyword     string
	TotalPage   int
	client      http.Client
	requestPool sync.Pool
	f           *internal.Write
}

// NewJob returns a new job with the specific parameters
func NewJob(log *zap.Logger, f io.WriterAt, keyword string) *Job {
	return &Job{
		Log:     log,
		Keyword: keyword,
		requestPool: sync.Pool{
			New: func() interface{} {
				req, err := http.NewRequest("GET", fmt.Sprintf(Search_URL, keyword, 1), nil)
				if err != nil {
					log.Error("New request err", zap.String("msg", err.Error()))
					return nil
				}
				req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
				return req
			},
		},
		f: internal.NewWriter(f),
	}
}

func (j *Job) newURL(page int) *url.URL {
	u, _ := url.Parse(fmt.Sprintf(Search_URL, j.Keyword, page))
	return u
}

func (j *Job) sendRequest(page int) (*Response, error) {
	req, ok := j.requestPool.Get().(*http.Request)
	if !ok {
		return nil, ErrNewRequest
	}
	// 设置url，不同的页网址不一样
	req.URL = j.newURL(page)
	r, err := j.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	resp := Response{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	data = utils.Convert(data)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// 获取总的页数
func (j *Job) getPage() {
	resp, err := j.sendRequest(1)
	if err != nil {
		j.Log.Info("Send request err", zap.String("err", err.Error()))
		return
	}
	j.TotalPage, _ = strconv.Atoi(resp.TotalPage)
}

// 处理每一个response内容
func (j *Job) handle(resp *Response) {
	var buff strings.Builder
	// 需要记录的所有数据,可修改
	for _, element := range resp.EngineSearchResult {
		buff.WriteString(element.JobName)
		buff.WriteByte(',')
		buff.WriteString(element.JobHref)
		buff.WriteByte(',')
		buff.WriteString(element.CompanyName)
		buff.WriteByte(',')
		buff.WriteString(element.CompanyHref)
		buff.WriteByte(',')
		buff.WriteString(element.ProvidesalaryText)
		buff.WriteByte(',')
		buff.WriteString(element.WorkareaText)
		buff.WriteByte(',')
		buff.WriteString(element.Issuedate)
		buff.WriteByte(',')
		buff.WriteString(element.Jobwelf)
		buff.WriteByte('\n')

	}
	j.f.Write([]byte(buff.String()))
}

func (j *Job) Run() {
	j.f.Write([]byte("job_name, job_href, company_name, company_href, salary, workarea, issuedate, jobwelf\n"))

	j.getPage()
	wg := sync.WaitGroup{}
	for i := 1; i <= j.TotalPage; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			resp, err := j.sendRequest(page)
			if err != nil {
				j.Log.Sugar().Errorf("send request err",
					zap.String("msg", err.Error()),
					zap.Int("page", page))
			}
			j.handle(resp)
		}(i)
	}
	wg.Wait()
}
