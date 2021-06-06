package handlers

import (
	"net/http"

	"github.com/gopher-translator-service/internal/cache"
	serviceErr "github.com/gopher-translator-service/internal/errors"
	"github.com/gopher-translator-service/internal/responses"
	"github.com/gopher-translator-service/pkg/logger"
)

type HistoryResponse struct {
	History []map[string]string `json:"history"`
}

func HistoryHandler(log logger.Logger, translationCache cache.KeyValueCache) http.HandlerFunc {
	ret := func(rw http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		kvs := translationCache.GetAll()
		hr := HistoryResponse{
			History: make([]map[string]string, 0, len(kvs)),
		}

		for _, v := range kvs {
			m := make(map[string]string, 1)
			m[v.Key] = v.Value
			hr.History = append(hr.History, m)
		}

		if err := responses.JsonResponse(rw, &hr); err != nil {
			httpErr := serviceErr.NewHttpErr(http.StatusInternalServerError, "failed to write history response", err)
			responses.ErrorResponse(rw, httpErr, log)
			return
		}
	}
	return ret
}
