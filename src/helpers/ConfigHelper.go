package helpers

import (
	"github.com/shenyisyn/goft-gin/goft"
	"io/ioutil"
	"os"
)

/**
 * @author  巨昊
 * @date  2021/10/21 16:25
 * @version 1.15.3
 */

func CertData(path string ) []byte{
	f,err:=os.Open(path )
	goft.Error(err)
	defer f.Close()
	b,err:=ioutil.ReadAll(f)
	goft.Error(err)
	return b
}