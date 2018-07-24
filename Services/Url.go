package Services

import (
"io/ioutil"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
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
    clienturl := "https://www.googleapis.com/urlshortener/v1/url?key=AIzaSyDBBlepvakv6nRw4pyA0bIxx_Ri6wbAPqY" 
  	 values := map[string]string{"longUrl": GetIpvalue.URL}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(clienturl, "application/json", bytes.NewBuffer(jsonValue))	
	bytedata,_ := ioutil.ReadAll(resp.Body)
	w.Write(bytedata)

	}

