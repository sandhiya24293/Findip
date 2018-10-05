package Services

import (
        "encoding/json"

        "io/ioutil"
        "log"
        "net/http"
        "strings"

        "github.com/PuerkitoBio/goquery"
)

func GetInnerSubstring(str string, prefix string, suffix string) string {
        var beginIndex, endIndex int
        beginIndex = strings.Index(str, prefix)
        if beginIndex == -1 {
                beginIndex = 0
                endIndex = 0
        } else if len(prefix) == 0 {
                beginIndex = 0
                endIndex = strings.Index(str, suffix)
                if endIndex == -1 || len(suffix) == 0 {
                        endIndex = len(str)
                }
        } else {
                beginIndex += len(prefix)
                endIndex = strings.Index(str[beginIndex:], suffix)
                if endIndex == -1 {
                        if strings.Index(str, suffix) < beginIndex {
                                endIndex = beginIndex
                        } else {
                                endIndex = len(str)
                        }
                } else {
                        if len(suffix) == 0 {
                                endIndex = len(str)
                        } else {
                                endIndex += beginIndex
                        }
                }
        }
return str[beginIndex:endIndex]
}

type ReadRank struct {
        Country     string
        Countryrank string
        Global      string
        Globalrank  string
}



type Domainstruct struct {
        Domain string
}

func Rating(w http.ResponseWriter, req *http.Request) {
        var country string
        var countryrank string
        var global string
        var globalrank string
        var domainstr Domainstruct
        err := json.NewDecoder(req.Body).Decode(&domainstr)
        if err != nil {
                log.Println("Error: Tenant vacate ", err)
        }
        url := "https://www.alexa.com/siteinfo/" + domainstr.Domain

        resp, _ := http.Get(url)

        data, _ := ioutil.ReadAll(resp.Body)
        p := strings.NewReader(string(data))
        doc, _ := goquery.NewDocumentFromReader(p)

        doc.Find("div").Each(func(i int, el *goquery.Selection) {

                el.Remove()

        })

   totaldata := doc.Text()
        trimglobal := GetInnerSubstring(totaldata, "!function(){", "}();")
        getvalue := GetInnerSubstring(trimglobal, "rank", "}")
        getfinalval := GetInnerSubstring(getvalue, ":{", "")

        values := strings.Split(getfinalval, ",")
        for i, v := range values {
                final := strings.Split(v, ":")
                if i == 0 {
                        country = final[0]
                        countryrank = final[1]
                } else {
                        global = final[0]
                        globalrank = final[1]
                }

        }
        datastruct := &ReadRank{}
        datastruct.Country = country
        datastruct.Countryrank = countryrank
        datastruct.Global = global
        datastruct.Globalrank = globalrank
        datasend, _ := json.Marshal(datastruct)
        w.Write(datasend)

}
