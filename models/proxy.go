package models

// Proxy struct (MODEL)
type Proxy struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

// RawProxy Method - Format raw proxy string
func (p *Proxy) RawProxy() string {
	raw := p.IP + ":" + string(p.Port)

	if p.User != "" && p.Pass != "" {
		raw = raw + "@" + p.User + ":" + p.Pass
	}

	return raw
}

// ConnProxy Method - Format connected http proxy
func (p *Proxy) ConnProxy() string {
	raw := p.IP + ":" + string(p.Port)

	if p.User != "" && p.Pass != "" {
		raw = p.User + ":" + p.Pass + "@" + raw
	}

	return raw
}
