package market

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/fromcode/market-api/internal/types"
	"github.com/fromcode/market-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { //mengatur handler pada url
		var market types.Markets

		// Membuat decoder JSON baru dari body permintaan (r.Body) dan mencoba
		// untuk men-decode (mengurai) data JSON ke dalam variabel 'market'
		err := json.NewDecoder(r.Body).Decode(&market)

		// Memeriksa apakah error yang terjadi adalah io.EOF (End of File).
		// Error ini biasanya terjadi jika body permintaan kosong.
		// Blok 'if' ini saat ini kosong, jadi tidak ada tindakan khusus yang diambil jika body kosong.
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty Body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//request validations
		if err := validator.New().Struct(market); err != nil {

			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidateError(validateErrs))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"succes": "OK"})
	}
}
