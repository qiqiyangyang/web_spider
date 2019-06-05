package parser

import(
    "regexp"
)

func Parser(body []byte)(s []string){
    url_list := []string{}
    m := map[string]int{}
    reg := regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
    //为保证url地址不重复，使用map
	for _, d := range reg.FindAllString(string(body), -1){
		m[d] = 0
    }
    for k, _ := range m{
        url_list = append(url_list, k)
    }
	return url_list
}
 