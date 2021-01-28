package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"hash"
	"io"
	"os"
)

/*

关于pkcs相关标准如下，摘自百度:

PKCS#1：定义RSA公开密钥算法加密和签名机制，主要用于组织PKCS#7中所描述的数字签名和数字信封[22]。
PKCS#3：定义Diffie-Hellman密钥交换协议[23]。
PKCS#5：描述一种利用从口令派生出来的安全密钥加密字符串的方法。使用MD2或MD5 从口令中派生密钥，并采用DES-CBC模式加密。主要用于加密从一个计算机传送到另一个计算机的私人密钥，不能用于加密消息[24]。
PKCS#6：描述了公钥证书的标准语法，主要描述X.509证书的扩展格式[25]。
PKCS#7：定义一种通用的消息语法，包括数字签名和加密等用于增强的加密机制，PKCS#7与PEM兼容，所以不需其他密码操作，就可以将加密的消息转换成PEM消息[26]。
PKCS#8：描述私有密钥信息格式，该信息包括公开密钥算法的私有密钥以及可选的属性集等[27]。
PKCS#9：定义一些用于PKCS#6证书扩展、PKCS#7数字签名和PKCS#8私钥加密信息的属性类型[28]。
PKCS#10：描述证书请求语法[29]。
PKCS#11：称为Cyptoki，定义了一套独立于技术的程序设计接口，用于智能卡和PCMCIA卡之类的加密设备[30]。
PKCS#12：描述个人信息交换语法标准。描述了将用户公钥、私钥、证书和其他相关信息打包的语法[31]。
PKCS#13：椭圆曲线密码体制标准[32]。
PKCS#14：伪随机数生成标准。
PKCS#15：密码令牌信息格式标准[33]。

*/
func getTestPriKey() *rsa.PrivateKey {
	// from http://web.chacuo.net/netrsakeypair
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

	blockPri, _ := pem.Decode([]byte(test2048KeyPri))
	if blockPri == nil {
		panic("failed to parse PEM block containing the public key")
	}
	pri, err := x509.ParsePKCS8PrivateKey(blockPri.Bytes)

	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}
	return pri.(*rsa.PrivateKey)
}

func getTestPubKey() rsa.PublicKey {
	test2048KeyPub := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDCG4VT1+s0kGqQ8FHrfkNMSshO
ggyoZpJAQ3+MzDxJ8S+QQghlKsLT8xLn3S9ZEOSl6hCDL/DRqBUlKmzztiaVkxaP
GPwXU3rtiilK/l1Q7wwWUr6sJMcf88sxg1f8bO9wfAv4fwbkBfRGzHdTKO+xCp3P
YysbbB3iIwKhSkY9QwIDAQAB
-----END PUBLIC KEY-----`

	blockPub, _ := pem.Decode([]byte(test2048KeyPub))
	if blockPub == nil {
		panic("failed to parse PEM block containing the public key")
	}
	pub, err := x509.ParsePKIXPublicKey(blockPub.Bytes)

	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	fmt.Println(pub)
	return pub.(rsa.PublicKey)
}

var (
	rng         io.Reader
	label       []byte
	hashes      hash.Hash
	test2048Key *rsa.PrivateKey
)

//func Decrypt(text string) string {
func Decrypt(text []byte) string {
	//test2048Key := getTestPriKey()
	ciphertext := text
	//ciphertext, _ := hex.DecodeString(text)
	//label := []byte("orders")

	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	//rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(hashes, rng, test2048Key, ciphertext, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return ""
	}

	//fmt.Printf("Plaintext: %s\n", string(plaintext))

	// Remember that encryption only provides confidentiality. The
	// ciphertext should be signed before authenticity is assumed and, even
	// then, consider that messages might be reordered.
	return string(plaintext)
}
func Encrypt(text string) []byte {
	//test2048Key := getTestPriKey()

	secretMessage := []byte(text)
	//label := []byte("orders")

	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	//rng := rand.Reader

	ciphertext, err := rsa.EncryptOAEP(hashes, rng, &test2048Key.PublicKey, secretMessage, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return nil
	}

	// Since encryption is a randomized function, ciphertext will be
	// different each time.
	//fmt.Printf("Ciphertext: %x\n", ciphertext)
	return ciphertext
}

func main() {
	test2048Key = getTestPriKey()
	rng = rand.Reader
	label = []byte("orders")
	hashes = sha256.New()

	text := "send reinforcements, we're going to advance"
	ciphertext := Encrypt(text)
	fmt.Println(Decrypt(ciphertext))
}
