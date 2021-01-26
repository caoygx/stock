package util

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

func FileGetContent(rurl string) []byte {
	r, err := RequestByParams(rurl, "", "get")
	if err != nil {
		//panic(err)
		var zero = []byte("")
		return zero
	}
	content := GetBody(r)
	return content
}

func PrintBody(r *http.Response) {
	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}
func GetBody(r *http.Response) []byte {
	defer func() { _ = r.Body.Close() }()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return content
	//return string(content)
}

func Post(rurl string, params string, method string) ([]byte, error) {
	return FileGetContent2(rurl, params, method)
}
func FileGetContent2(rurl string, params string, method string) ([]byte, error) {
	//time.Sleep(time.Second * time.Duration(RandInt64(1, 15)))
	r, err := RequestByParams(rurl, params, method)
	if err != nil {
		fmt.Println("err:", err)
		var zero = []byte("")
		return zero, err
	}
	content := GetBody(r)
	return content, nil
}

func RequestByParams(rurl string, params string, method string) (*http.Response, error) {

	fmt.Println(rurl)
	//
	/*proxyUrl, err := url.Parse("http://127.0.0.1:8888")
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}*/
	var err error
	myClient := &http.Client{}

	var request *http.Request

	if method == "post" {
		data := strings.NewReader(params)
		request, err = http.NewRequest(http.MethodPost, rurl, data)
		request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	} else {
		request, err = http.NewRequest(http.MethodGet, rurl, nil)
	}

	request.Header.Add("Accept", "application/json, text/plain, */*")
	//request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")

	//request.Header.Add("Referer", "https://data.appgrowing.cn/leaflet/brand")
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36 SocketLog(tabid=241&client_id=)")
	request.Header.Add("cookie", "st_asi=20200925190338206-119098306875-0142492434-page.f10v2.cwfx-2; st_inirUrl=https%3A%2F%2Fm.baidu.com%2Ffrom%3D844b%2Fbd_page_type%3D1%2Fssid%3D9ae77169686a6e1d04%2Fuid%3D0%2Fpu%3Dusm%25402%252Csz%25401320_2001%252Cta%2540iphone_1_12.0_25_12.0%2Fbaiduid%3D48848C376022410EAC98D679F7343EFD%2Fw%3D0_10_%2Ft%3Diphone%2Fl%3D1%2Ftc; st_psi=20200925190338206-119098306875-0142492434; st_pvi=48734839687341; st_sn=5; st_sp=2020-01-10%2009%3A25%3A58; qgqp_b_id=f6a98b12ff1ec83b550498807d4ccc5a; st_si=31171944906971; wap_ck1=true; wap_ck2=true")
	//req.Header.Set("Cookie", "name=anny")

	//r, err := http.DefaultClient.Do(request)
	r, err := myClient.Do(request)
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	//fmt.Println(r)
	//printBody(r)
	return r, err

}

func Download(imgUrl string) string {

	dir, _ := os.Getwd()
	dir = dir + string(os.PathSeparator) + "down" + string(os.PathSeparator)
	f, err := os.Stat(dir)
	fmt.Println(f)
	if err != nil {
		os.MkdirAll(dir, os.ModePerm)
	}
	fmt.Println(dir)
	//os.Exit(1)

	fileName := path.Base(imgUrl)

	res, err := RequestByParams(imgUrl, "", "get")
	if err != nil {
		return ""
	}
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	savePath := dir + fileName
	file, err := os.Create(savePath)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, err := io.Copy(writer, reader)
	if err == nil {
		fmt.Printf("Total length: %d", written)
		return savePath
	}
	return ""

}

/*
func printBody(r *http.Response) {
	defer func() { _ = r.Body.Close() }()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}
func GetBody(r *http.Response) []byte {
	defer func() { _ = r.Body.Close() }()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return content
	//return string(content)
}*/

func requestByParams(rurl string) (*http.Response, error) {

	//rurl = "http://ad.uzipm.com/adMaterial/index?ret_format=json"
	request, err := http.NewRequest(http.MethodGet, rurl, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36 SocketLog(tabid=392&client_id=)")
	params := make(url.Values)
	params.Add("ret_format", "json")
	//request.URL.RawQuery = params.Encode()

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	//fmt.Println(r)
	//printBody(r)
	return r, err

	/*
		proxyUrl, err := url.Parse("http://127.0.0.1:8888")
		myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
		fmt.Println(rurl)
		request, err := http.NewRequest(http.MethodGet, rurl, nil)
		request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36 SocketLog(tabid=392&client_id=)")

		request.Header.Add("Accept-Encoding", "gzip, deflate, br") //must

		//r, err := http.DefaultClient.Do(request)
		r, err := myClient.Do(request)
		if err != nil {
			//panic(err)
			fmt.Println(err)
		}
		//fmt.Println(r)
		//printBody(r)
		return r, err
	*/
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
