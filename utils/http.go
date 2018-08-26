package utils

import (
	"net/http"
	"net/url"
	"time"
	"reflect"
	"log"
)



func NewHttpClient(timeout int,proxyUrl string)  *http.Client{
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}

	transport := &http.Transport{Proxy: proxy}
	if proxyUrl==""{
		return &http.Client{Timeout:  time.Duration(timeout) * time.Second}
	}
	return &http.Client{Transport: transport, Timeout:  time.Duration(timeout) * time.Second}
}


/**
  @retry  重试次数
  @method 调用的函数，
  @params 参数,顺序一定要按照实际调用函数入参顺序一样
  @return 返回
*/
func RE(retry int, method interface{}, params ...interface{}) interface{} {

	invokeM := reflect.ValueOf(method)
	if invokeM.Kind() != reflect.Func {
		panic("method not a function")
		return nil
	}

	var value []reflect.Value = make([]reflect.Value, len(params))
	var i int = 0
	for ; i < len(params); i++ {
		value[i] = reflect.ValueOf(params[i])
	}

	var retV interface{}
	var retryC int = 0
_CALL:
	if retryC > 0 {
		log.Println("sleep....", time.Duration(retryC*500*int(time.Millisecond)))
		time.Sleep(time.Duration(retryC * 500 * int(time.Millisecond)))
	}

	retValues := invokeM.Call(value)

	for _, vl := range retValues {
		if vl.Type().String() == "error" {
			if !vl.IsNil() {
				log.Println(vl)
				retryC++
				if retryC <= retry {
					log.Printf("Invoke Method[%s] Error , Begin Retry Call [%d] ...", invokeM.String(), retryC)
					goto _CALL
				} else {
					log.Println("Invoke Method Fail ???" + invokeM.String())
					//panic("Invoke Method Fail ???" + invokeM.String())
				}
			}
		} else {
			retV = vl.Interface()
		}
	}

	return retV
}

