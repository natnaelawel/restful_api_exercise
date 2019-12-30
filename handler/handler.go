package handler

import (
	"reflect"
	"net/http"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"html/template"
	"github.com/natnaelawel/restfulexercise/entity"
)

var Datas entity.Data
var Templ *template.Template

func init(){
	Templ = template.Must(template.ParseGlob("template/*.html"))
}

func HomePage(writer http.ResponseWriter, r *http.Request) {
	Templ.ExecuteTemplate(writer, "index.html",nil)
}
func ReadCurrencyData(writer http.ResponseWriter, request *http.Request) {
	var URL  = `http://data.fixer.io/api/latest?access_key=716fc6e5f8e9f4f84868c837f18f8e50`
	amount,_ := strconv.ParseFloat(request.FormValue("currency_amount"),64)
	fromOption := request.FormValue("fromOption")
	toOption := request.FormValue("toOption")

	response, err := http.Get(URL)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}else {
		data,_ := ioutil.ReadAll(response.Body)
		//	_ = json.NewDecoder(response.Body).Decode(&Datas)
		err = json.Unmarshal(data, &Datas)
		r := reflect.ValueOf(Datas.Rates)
		fromData := float64(reflect.Indirect(r).FieldByName(fromOption).Float())
		toData := float64(reflect.Indirect(r).FieldByName(toOption).Float())

		result := amount*(toData/fromData)

		templateData := struct {
			From string
			To string
			EnteredValue float64
			Value float64
			Currency float64
			Date string
		}{
			fromOption,
			toOption,
			amount,
			result,
			toData/fromData,
			Datas.Date,
		}
		Templ.ExecuteTemplate(writer,"home.html", templateData)
	}

}
