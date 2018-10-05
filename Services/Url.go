package Services

import (

	"encoding/json"
	"net/http"
	
	"fmt"
	"encoding/base64"
)


type Shorten struct{
	URL string
}

type URLoutput struct{
	longUrl string `json:"longUrl"`
}

func ShorternURL(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Shorten
	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}
  	str := UrlEncoded(GetIpvalue.URL)
	outputstring := "http://foxu.tech:8087/?url=" +str
	w.Write([]byte(outputstring))


	}

func Getredirect(w http.ResponseWriter, r *http.Request) {
	
	  param1 := r.URL.Query().Get("url")
	decodestring := Decodestring(param1)
	http.Redirect(w, r, decodestring, http.StatusSeeOther)
	}

func UrlEncoded(str string) string {
  data := []byte(str)
	str1 := base64.StdEncoding.EncodeToString(data)
	
	return str1
}

func Decodestring(str string )string{
	
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		
	}
	
	data1 := string(data[:])
	return data1
}
	