package job

import "errors"

// Search_URL is the api address which returns information
// the first placeholder is the keyword while the sencond is the page number
const Search_URL = "https://search.51job.com/list/010000,000000,0000,00,9,99,%s,2,%d.html"

var ErrNewRequest = errors.New("new request error")