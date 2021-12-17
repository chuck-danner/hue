package hue

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var client *http.Client

var headerName = "hue-application-key"

func init() {
	//cert, err := ioutil.ReadFile("hue_bridge.pem")
	pem := []byte(`-----BEGIN CERTIFICATE-----
MIICMjCCAdigAwIBAgIUO7FSLbaxikuXAljzVaurLXWmFw4wCgYIKoZIzj0EAwIw
OTELMAkGA1UEBhMCTkwxFDASBgNVBAoMC1BoaWxpcHMgSHVlMRQwEgYDVQQDDAty
b290LWJyaWRnZTAiGA8yMDE3MDEwMTAwMDAwMFoYDzIwMzgwMTE5MDMxNDA3WjA5
MQswCQYDVQQGEwJOTDEUMBIGA1UECgwLUGhpbGlwcyBIdWUxFDASBgNVBAMMC3Jv
b3QtYnJpZGdlMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEjNw2tx2AplOf9x86
aTdvEcL1FU65QDxziKvBpW9XXSIcibAeQiKxegpq8Exbr9v6LBnYbna2VcaK0G22
jOKkTqOBuTCBtjAPBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNV
HQ4EFgQUZ2ONTFrDT6o8ItRnKfqWKnHFGmQwdAYDVR0jBG0wa4AUZ2ONTFrDT6o8
ItRnKfqWKnHFGmShPaQ7MDkxCzAJBgNVBAYTAk5MMRQwEgYDVQQKDAtQaGlsaXBz
IEh1ZTEUMBIGA1UEAwwLcm9vdC1icmlkZ2WCFDuxUi22sYpLlwJY81Wrqy11phcO
MAoGCCqGSM49BAMCA0gAMEUCIEBYYEOsa07TH7E5MJnGw557lVkORgit2Rm1h3B2
sFgDAiEA1Fj/C3AN5psFMjo0//mrQebo0eKd3aWRx+pQY08mk48=
-----END CERTIFICATE-----`)
	// if err != nil {
	// 	panic(err)
	// }
	root := x509.NewCertPool()
	root.AppendCertsFromPEM(pem)
	conf := &tls.Config{
		RootCAs:            root,
		InsecureSkipVerify: true, //ugh... can't figure out the SANs vs common name
	}
	dialContext := func(ctx context.Context, network, addr string) (net.Conn, error) {
		dialer := &net.Dialer{}
		if addr == "id:443" {
			addr = "ip:443"
		}

		return dialer.DialContext(ctx, network, addr)
	}

	client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: conf,
			DialContext:     dialContext,
		},
	}

}

func (h Hue) getRequest(url string, v interface{}) {
	h.request("GET", url, nil, v)
}

func (h Hue) putRequest(url string, body io.Reader, v interface{}) {
	h.request("PUT", url, body, v)
}

func (h Hue) request(method string, url string, body io.Reader, v interface{}) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err.Error())
	}
	request.Header.Add(headerName, h.Key)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(request)

	if err != nil {
		panic(err.Error())
	}
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(data, &v)

	if err != nil {
		panic(err.Error())
	}
}
