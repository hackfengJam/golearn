package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func refundAction() {
	// data
	var list []string
	list = []string{`<xml><return_code>SUCCESS</return_code><appid><![CDATA[wx314dd070d881531f]]></appid><mch_id><![CDATA[1496104772]]></mch_id><nonce_str><![CDATA[d4be57b1c66dbbe3125877ffea99296c]]></nonce_str><req_info><![CDATA[XUVfsYUPnHiyDVtD39dqigl5WOOMNdZxHWKXsZgsF4KiFLA5TssR/WPGnokQ3T/Xlxleenk+XCqhi6I+VGwK1k9LJe2I+CAOXOt7wJAEUjNesJxvuMSvSETCW+FjQH4/AtlXD8INn8fEnUVHsiuJU85pXF5d/05fZgTnS7/Nu/lJshb875tHw/6XVtLP7bWcLvPLevtS17fLlk/OR1dczX1E1pfUGejf8Dzsyg8ZySweSpZt4f0z7lp3rN/Jf06IiSXwm74w8QWZuqRNcymIHTLeePf41RjMwsqcMP/OgHK86djhHy56YKtGP/SA/125vSKfLZPlPwlpcu44bm1pmNINHd2l3II3BFE3n7DUfGf6+Cfmm/GD91rGaMRnFkSscR35APr1HXmEql892xV96VVW4Q7XND3jQXIMk6fGK5hDyCJkpGSyAoQHmIW+HS6jjokmPddSOtjtmFY6gUiBkloEnvyMuu1pdAtCO4BrA7Zp5o90asTGdKJCWTnapQSxvV8VYmfPE8ydP6QdBUSYwkuKBpSiXchhLwVMCdN6zveXMq0rvbMUSbJcVTI1tGRmDBZhol2CjRj/y/dsfD6Q0lo/70mvYRbvur5FsnPa02+CqbXU4ac4SJ/qbNz3tSSFCuJ6ivD4S/vSYjAk1xeR8CY44l/sAP5bcacaK3TQ7AF/8MsowryCZVokZNgp+2gi7Vi6dbjXmLej8ZcdzP6QD0RS62xOhF3aUyRVrx1R0LG5K2jzP08KPGxc+H/IQ6jCsrpdtK34NV0jNcDIXyw9tADPQy92gfagM0UYCua/2n8Y4lshPpFowbqZZSd9EK5RKLhUgvhIibS1LSC3YTOHLuAWFCQy/o8pk0FIq3iTbljDsy42vncPlvLUTOjFqVJ/I3/uJ0ir87EbWTbpkvcU6J3g6vLOJnnigkiHhK7e8ED0neky0DA9BKeYVHHj/nTXw1QJjnfAGzIjZBnMYxiblrQPa1ynhV7x49e5ecvytulfhumjOTnvL4kEFKiF6qTGInXCSRUl0isGkpvxFYKBjCI7nbim6U5rZF9seqxLMEaANBEg//8viZhIhGxrrFL3GRmCvZ9zfoWqHQVaqZEeeUhrAv8ih2E4ddiocSzfVQw=]]></req_info></xml>
`}

	// req
	url := "http://pre.api.mengtuiapp.com/notify/wxpay/refund.action"
	headers := map[string]string{}
	client := &http.Client{}
	var (
		req  *http.Request
		resp *http.Response
	)

	// post
	var err error
	for _, v := range list {
		// construct req
		req, err = http.NewRequest(
			http.MethodPost,
			url, bytes.NewBuffer([]byte(v)))
		if err != nil {
			panic(err)
		}

		// add header
		for k, v := range headers {
			req.Header.Add(k, v)
		}

		// req
		resp, err = client.Do(req)
		if err != nil {
			panic(err)
		}

		// resp
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// log
		fmt.Printf("resp.header: %v, status: %s, resp: %v \n", resp.Header.Get("X-Reqid"), resp.Status, string(data))
	}

	// close
	if resp != nil {
		_ = resp.Body.Close()
	}
}

func main() {

}
