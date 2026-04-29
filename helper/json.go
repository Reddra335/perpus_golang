package helper

import (
	"encoding/json"
	"net/http"
)

func WriteFromBody(request *http.Request, result any) {
	decode := json.NewDecoder(request.Body)
	err := decode.Decode(&result)
	ErrorT(err)

}

func WriteToBody(writer http.ResponseWriter, result any) {
	writer.Header().Add("Content-Type", "Aplication/Json")
	encode := json.NewEncoder(writer)
	err := encode.Encode(result)
	ErrorT(err)
}
