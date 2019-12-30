package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/natnaelawel/restfulexercise/handler"
)

func main()  {
		router := mux.NewRouter()
		router.PathPrefix("/assets/").
			Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("."+"/assets"))))
		router.HandleFunc("/convert", handler.ReadCurrencyData).Methods("GET")
		router.HandleFunc("/", handler.HomePage).Methods("GET")
		http.ListenAndServe(":12345", router)
}

//mapdata := map[string]interface{}{"Data":Datas.Rates}
//val := reflect.Indirect(reflect.ValueOf(Datas.Rates))
//mapdata := make(map[string]float64)
//json.Unmarshal(byte(Datas.Rates), &mapdata)
//fmt.Println(mapdata)
//var mapdata map[string]entity.Rates
//inrec, _ := json.Marshal(Datas)
//json.Unmarshal(data, &mapdata)
//for key,value := range mapdata  {
//
//}
//fmt.Println(mapdata["rates"])
//writer.Header().Set("Content-Type","application/json")
//json.NewEncoder(writer).Encode(&Datas)

//params := mux.Vars(request)
//fmt.Println(params["value"])
//var URL  = `http://data.fixer.io/api/latest?access_key=716fc6e5f8e9f4f84868c837f18f8e50`
//
////http://apilayer.net/api/live&access_key=3fd7d07251450fb9db97b357b23aecdd
////https://api.exchangeratesapi.io/latest
////https://free.currconv.com/api/v7/convert?q=USD_PHP&compact=ultra&apiKey=d947ebd4a524b083919b
////	response, err := http.Get("https://api.currconv.com/api/v7/convert?q=USD_PHP,PHP_USD&compact=ultra&apiKey=d947ebd4a524b083919b")
////	response, err := http.Get("https://free.currconv.com/api/v7/convert?q=USD_PHP&compact=ultra&apiKey=d947ebd4a524b083919b")
////	response, err := http.Get("http://apilayer.net/api/live&access_key=3fd7d07251450fb9db97b357b23aecdd")
//
//response, err := http.Get(URL)
////response = http.NewRequest("GET",URL,request.B)
//if err != nil {
//	fmt.Println(err)
//}else {
//	data,_ := ioutil.ReadAll(response.Body)
//	//	_ = json.NewDecoder(response.Body).Decode(&Datas)
//	err = json.Unmarshal(data, &Datas)
//	fmt.Println(Datas)
//	writer.Header().Set("Content-Type","application/json")
//	json.NewEncoder(writer).Encode(&Datas)
//}