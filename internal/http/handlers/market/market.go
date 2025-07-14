package market

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/fromcode/market-api/internal/types"
	"github.com/fromcode/market-api/internal/utils/response"
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

		// Pesan ini menandakan bahwa proses untuk "Membuat Product Baru" sedang dimulai.
		slog.Info("Membuat Product Baru")

		response.WriteJson(w, http.StatusCreated, map[string]string{"succes": "OK"})
	}
}
