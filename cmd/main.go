package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const urlTpl = "https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s"

func main() {
	println("Starting ...")

	http.HandleFunc("/api/dyndns", dynDnsHandler)
	http.HandleFunc("/_health", healthHandler)

	err := http.ListenAndServe(":6221", nil)

	if err != nil {
		panic(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func dynDnsHandler(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()

	username, password, _ := r.BasicAuth()

	rsp, err := putRequest(
		username,
		password,
		params.Get("hostname"),
		params.Get("ipAddress"),
		params.Get("zoneId"),
		params.Get("dnsId"),
	)

	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		print(err)
	} else {
		if rsp.StatusCode > 299 {
			w.WriteHeader(rsp.StatusCode)

			var b []byte
			_, _ = rsp.Body.Read(b)
			_, _ = w.Write(b)

			return
		}
		w.WriteHeader(204)
	}
}

func putRequest(email, apiKey, hostname, ipAddress, cloudflareZoneId, cloudflareDnsId string) (*http.Response, error) {
	data := &map[string]interface{}{
		"type":    "A",
		"name":    hostname,
		"content": ipAddress,
		"ttl":     120,
		"proxied": false,
	}

	client := &http.Client{}

	url := fmt.Sprintf(urlTpl, cloudflareZoneId, cloudflareDnsId)

	bb, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bb))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Email", email)
	req.Header.Set("X-Auth-Key", apiKey)

	return client.Do(req)
}
