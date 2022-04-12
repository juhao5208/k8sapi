package models

/**
 * @author  巨昊
 * @date  2021/9/7 21:11
 * @version 1.15.3
 */

type CertModel struct {
	CN        string //域名
	Algorithm string //算法
	Issuer    string //签发者
	BeginTime string //生效时间
	EndTime   string //到期时间
}
