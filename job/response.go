package job

type Response struct {
	TopAds                []interface{}               `json:"top_ads"`               
	AuctionAds            []interface{}               `json:"auction_ads"`           
	MarketAds             []interface{}               `json:"market_ads"`            
	EngineSearchResult    []EngineSearchResultElement `json:"engine_search_result"`  
	JobidCount            string                      `json:"jobid_count"`           
	BannerAds             string                      `json:"banner_ads"`            
	IsCollapseexpansion   string                      `json:"is_collapseexpansion"`  
	CoAds                 []interface{}               `json:"co_ads"`                
	KeywordRecommendation KeywordRecommendation       `json:"keyword_recommendation"`
	SearchCondition       map[string]string           `json:"search_condition"`      
	SearchedCondition     string                      `json:"searched_condition"`    
	CurrPage              string                      `json:"curr_page"`             
	TotalPage             string                      `json:"total_page"`            
	KeywordAds            []interface{}               `json:"keyword_ads"`           
	JobSearchAssistance   []map[string]string         `json:"job_search_assistance"` 
	SEOTitle              string                      `json:"seo_title"`             
	SEODescription        string                      `json:"seo_description"`       
	SEOKeywords           string                      `json:"seo_keywords"`          
}

type EngineSearchResultElement struct {
	Type              Type       `json:"type"`              
	Jt                Jt         `json:"jt"`                
	Tags              []string   `json:"tags"`              
	AdTrack           string     `json:"ad_track"`          
	Jobid             string     `json:"jobid"`             
	Coid              string     `json:"coid"`              
	Effect            string     `json:"effect"`            
	IsSpecialJob      string     `json:"is_special_job"`    
	JobHref           string     `json:"job_href"`          
	JobName           string     `json:"job_name"`          
	JobTitle          string     `json:"job_title"`         
	CompanyHref       string     `json:"company_href"`      
	CompanyName       string     `json:"company_name"`      
	ProvidesalaryText string     `json:"providesalary_text"`
	Workarea          string     `json:"workarea"`          
	WorkareaText      string     `json:"workarea_text"`     
	Updatedate        Updatedate `json:"updatedate"`        
	Iscommunicate     string     `json:"iscommunicate"`     
	CompanytypeText   string     `json:"companytype_text"`  
	Degreefrom        string     `json:"degreefrom"`        
	Workyear          string     `json:"workyear"`          
	Issuedate         string     `json:"issuedate"`         
	IsFromXyz         string     `json:"isFromXyz"`         
	IsIntern          string     `json:"isIntern"`          
	Jobwelf           string     `json:"jobwelf"`           
	JobwelfList       []string   `json:"jobwelf_list"`      
	Isdiffcity        string     `json:"isdiffcity"`        
	AttributeText     []string   `json:"attribute_text"`    
	CompanysizeText   string     `json:"companysize_text"`  
	CompanyindText    string     `json:"companyind_text"`   
	Adid              string     `json:"adid"`              
}

type KeywordRecommendation struct {
	Title    string  `json:"title"`    
	DataType string  `json:"data_type"`
	Keyword  string  `json:"keyword"`  
	Data     []Datum `json:"data"`     
}

type Datum struct {
	Href  string `json:"href"` 
	Text  string `json:"text"` 
	Click string `json:"click"`
}

type Jt string
const (
	The0_0 Jt = "0_0"
	The0_1 Jt = "0_1"
)

type Type string
const (
	EngineSearchResult Type = "engine_search_result"
)

type Updatedate string
const (
	The0618 Updatedate = "06-18"
	The0619 Updatedate = "06-19"
)
