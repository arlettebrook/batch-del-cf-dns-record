package models

type DNSRecord struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type ResultInfo struct {
	TotalCount int `json:"total_count"`
	TotalPages int `json:"total_pages"`
}

type Result struct {
	Result     []DNSRecord `json:"result"`
	ResultInfo `json:"result_info"`
}
