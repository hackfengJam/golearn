package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	test2048KeyPub := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDCG4VT1+s0kGqQ8FHrfkNMSshO
ggyoZpJAQ3+MzDxJ8S+QQghlKsLT8xLn3S9ZEOSl6hCDL/DRqBUlKmzztiaVkxaP
GPwXU3rtiilK/l1Q7wwWUr6sJMcf88sxg1f8bO9wfAv4fwbkBfRGzHdTKO+xCp3P
YysbbB3iIwKhSkY9QwIDAQAB
-----END PUBLIC KEY-----`
	test2048KeyPri := `-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAMIbhVPX6zSQapDw
Uet+Q0xKyE6CDKhmkkBDf4zMPEnxL5BCCGUqwtPzEufdL1kQ5KXqEIMv8NGoFSUq
bPO2JpWTFo8Y/BdTeu2KKUr+XVDvDBZSvqwkxx/zyzGDV/xs73B8C/h/BuQF9EbM
d1Mo77EKnc9jKxtsHeIjAqFKRj1DAgMBAAECgYB7e76J5CaBPg5NPlUfFygA9OIQ
77LVvsrzjv0puRPxhjnX7+ofUeC3rT8tugxRAOo0kn8Gtgzhk6hACIlfUVWhXAo6
onpXLT8GaiCGO98di9/sYHTLOaSDWN9yIO1cBpC/gLUEgCx5oALbD7ZA8JPox7KP
W1xNLkFitmgTsNEnIQJBAP1YFofuE/h3IaFFv5E6FVtcXPaaUWjkk4AIDv+pj0+v
c6Q9KlBG2QNlZGgAVuTjn5lFu5/qIAy0UgWGxNaqeZcCQQDEJHTnFg8ZJzDsHOR7
eRbp1m5GQDuyxj+LASIR6uoOjsRRrTkW8crtW+a5Vvp0mWJZbPW8v0vRreLJ+8+5
NZc1AkAqNY3zccgkAn046G0FXj8GrLnUYFul9UdZ8n3FPNPiu+GxgHtXSqaHeVIk
PcI8emwx1jtvZkuWskhkIVMUTOzbAkEAug2y7nkYPU3VtEL74LvOJmYHGJBZkI1J
PTwH62MgPkriom9kVgVp7plcVLbSwMO2bQlUWIRFEVKWa+527kKNeQJBAMplBBjN
tNdRQIqHmeNN6gwW1vJJr56N07VAFQM6t9YKjzCFrPTLJAVls26tSa+l0hI2eZMe
hKHe9UENOfl8/0o=
-----END PRIVATE KEY-----`

	blockPub, _ := pem.Decode([]byte(test2048KeyPub))
	if blockPub == nil {
		panic("failed to parse PEM block containing the public key")
	}
	pub, err := x509.ParsePKIXPublicKey(blockPub.Bytes)

	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	fmt.Println(pub)

	blockPri, _ := pem.Decode([]byte(test2048KeyPri))
	if blockPri == nil {
		panic("failed to parse PEM block containing the public key")
	}
	pri, err := x509.ParsePKCS8PrivateKey(blockPri.Bytes)

	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}
	fmt.Println(pri)

}
