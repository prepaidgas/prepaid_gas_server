package http

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/SteMak/prepaid_gas_server/config"
	"github.com/SteMak/prepaid_gas_server/structs"
)

var (
	err error
)

func Init() error {
	// TODO: Add view endpoints
	http.HandleFunc("/", Validator)

	return http.ListenAndServe(":"+strconv.Itoa(config.ValidatorPort), nil)
}

func Validator(w http.ResponseWriter, r *http.Request) {
	var message structs.Message

	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	err = message.ValidateEarlyLiquidation(20)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	sign, err := structs.SignMessage(message)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, hex.EncodeToString(sign))
}