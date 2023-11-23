package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/shenyisyn/goft-gin/goft"
	"io/ioutil"
	"k8sapi2/src/models"
	"log"
	"math/big"
	rd "math/rand"
	"os"
	"time"
)

/**
 * @author  巨昊
 * @date  2021/9/7 21:10
 * @version 1.15.3
 */

const CAFILE = "./test/certs/ca.crt"
const CAKEY = "./test/certs/ca.key"

func DeleteK8sUser(cn string) {
	err := os.Remove(fmt.Sprintf("./k8suser/#{cn}.pem"))
	goft.Error(err)
	err = os.RemoveAll(fmt.Sprintf("./k8suser/#{cn}_key.pem"))
	goft.Error(err)
}

//签发用户证书
func GenK8sUser(cn, o string) { //用户和用户组
	//caCert, caPriKey := parseK8sCA(CAFILE, CAKEY)
	//if cn == "" {
	//	goft.Error(fmt.Errorf("CN is required"))
	//}

	//解析k8s ca 和key 文件
	caFile, err := ioutil.ReadFile(CAFILE)
	if err != nil {
		log.Fatal(err)
	}
	caBlock, _ := pem.Decode(caFile)

	caCert, err := x509.ParseCertificate(caBlock.Bytes) //ca 证书对象
	if err != nil {
		log.Fatal(err)
	}
	//解析私钥
	keyFile, err := ioutil.ReadFile(CAKEY)
	if err != nil {
		log.Fatal(err)
	}
	keyBlock, _ := pem.Decode(keyFile)
	caPriKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes) //是要对象
	if err != nil {
		log.Fatal(err)
	}

	//构建证书模板
	certTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(rd.Int63()), //证书序列号
		Subject: pkix.Name{
			Country: []string{"CN"},
			//Organization:       []string{"填的话这里可以用作用户组"},
			//OrganizationalUnit: []string{"可填课不填"},
			Province:   []string{"beijing"},
			CommonName: "lisi", //CN
			Locality:   []string{"beijing"},
		},
		NotBefore:             time.Now(),                                                                 //证书有效期开始时间
		NotAfter:              time.Now().AddDate(1, 0, 0),                                                //证书有效期
		BasicConstraintsValid: true,                                                                       //基本的有效性约束
		IsCA:                  false,                                                                      //是否是根证书
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}, //证书用途(客户端认证，数据加密)
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		EmailAddresses:        []string{"callmejustin7@gmail.com"},
	}

	//生成公私钥--秘钥对
	priKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	//创建证书 对象
	clientCert, err := x509.CreateCertificate(rand.Reader, certTemplate, caCert, &priKey.PublicKey, caPriKey)
	if err != nil {
		return
	}

	//编码证书文件和私钥文件
	clientCertPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: clientCert,
	}

	clientCertFile, err := os.OpenFile("./test/certs/lisi.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	goft.Error(err)
	defer clientCertFile.Close()
	err = pem.Encode(clientCertFile, clientCertPem)
	goft.Error(err)

	buf := x509.MarshalPKCS1PrivateKey(priKey)
	keyPem := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: buf,
	}
	clientKeyFile, _ := os.OpenFile("./test/certs/lisi_key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	err = pem.Encode(clientKeyFile, keyPem)
	if err != nil {
		log.Fatal(err)
	}

}

func getCertType(alg x509.PublicKeyAlgorithm) string {
	switch alg {
	case x509.RSA:
		return "RSA"
	case x509.DSA:
		return "DSA"
	case x509.ECDSA:
		return "ECDSA"
	case x509.UnknownPublicKeyAlgorithm:
		return "Unknow"
	}
	return "Unknow"
}

//解析证书
func ParseCert(crt []byte) *models.CertModel {
	var cert tls.Certificate
	//解码证书
	certBlock, restPEMBlock := pem.Decode(crt)
	if certBlock == nil {
		return nil
	}
	cert.Certificate = append(cert.Certificate, certBlock.Bytes)
	//处理证书链
	certBlockChain, _ := pem.Decode(restPEMBlock)
	if certBlockChain != nil {
		cert.Certificate = append(cert.Certificate, certBlockChain.Bytes)
	}

	//解析证书
	x509Cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil
	} else {
		return &models.CertModel{
			CN:        x509Cert.Subject.CommonName,
			Issuer:    x509Cert.Issuer.CommonName,
			Algorithm: getCertType(x509Cert.PublicKeyAlgorithm),
			BeginTime: x509Cert.NotBefore.Format("2006-01-02 15:03:04"),
			EndTime:   x509Cert.NotAfter.Format("2006-01-02 15:03:04"),
		}
	}
}
