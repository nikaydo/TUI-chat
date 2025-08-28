package token

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

type Data struct {
	Ip        string   `json:"ip"`
	Isp       isp      `json:"isp"`
	Localtion location `json:"location"`
	Risk      risk     `json:"risk"`
}

type isp struct {
	Asn string `json:"asn"`
	Org string `json:"org"`
	Isp string `json:"isp"`
}

type location struct {
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Zipcode     string  `json:"zipcode"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Timezone    string  `json:"timezone"`
	Localtime   string  `json:"localtime"`
}

type risk struct {
	IsMobile     bool `json:"is_mobile"`
	IsVpn        bool `json:"is_vpn"`
	IsTor        bool `json:"is_tor"`
	IsProxy      bool `json:"is_proxy"`
	Isdatacenter bool `json:"is_datacenter"`
	RiskScore    int  `json:"risk_score"`
}

func GetToken() (string, error) {
	resp, err := http.Get("https://api.ipquery.io/?format=json")
	if err != nil {
		fmt.Println("Error fetching IP:", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}
	var d Data
	if err := json.Unmarshal(body, &d); err != nil {
		fmt.Println(err)
		return "", err
	}
	ip := net.ParseIP(d.Ip)
	ip4 := ip.To4()
	return base64.RawURLEncoding.EncodeToString([]byte(ip4)), err
}

func TokenToIP(token string) (net.IP, error) {
	data, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	if len(data) != 4 {
		return nil, fmt.Errorf("invalid data length")
	}
	return net.IPv4(data[0], data[1], data[2], data[3]), nil
}
