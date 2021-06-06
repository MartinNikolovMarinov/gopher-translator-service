package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gopher-translator-service/internal/cache"
	serviceErr "github.com/gopher-translator-service/internal/errors"
	"github.com/gopher-translator-service/internal/language"
	"github.com/gopher-translator-service/internal/responses"
	"github.com/gopher-translator-service/internal/validators"
	"github.com/gopher-translator-service/pkg/logger"
)

type WordReq struct {
	EnglishWord string `json:"english-word"`
}

type WordResp struct {
	GopherWord string `json:"gopher-word"`
}

func WordHandler(log logger.Logger, translationCache cache.KeyValueCache) http.HandlerFunc {
	ret := func(rw http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		var wordReq WordReq
		jd := json.NewDecoder(req.Body)
		if err := jd.Decode(&wordReq); err != nil {
			responses.ErrorResponse(rw, serviceErr.NewHttpErr(http.StatusBadRequest, "failed to decode body", err), log)
			return
		}
		if wordReq.EnglishWord == "" {
			err := serviceErr.NewHttpErr(http.StatusBadRequest,
				"Please provide english-word in request body",
				errors.New("invalid request"))
			responses.ErrorResponse(rw, err, log)
			return
		}

		wordReq.EnglishWord = strings.TrimSpace(wordReq.EnglishWord)
		if !validators.IsValidEnglishWord(wordReq.EnglishWord) {
			err := serviceErr.NewHttpErr(http.StatusBadRequest,
				fmt.Sprintf("'%s' is NOT a valid english word", wordReq.EnglishWord),
				errors.New("invalid english-word"))
			responses.ErrorResponse(rw, err, log)
			return
		}

		var translated string
		if fromCache := translationCache.Get(wordReq.EnglishWord); fromCache != "" {
			translated = fromCache
			log.Infoln(wordReq.EnglishWord, "translating from cache", translated)
		} else {
			translated = language.TranslateWord(wordReq.EnglishWord)
			log.Infoln(wordReq.EnglishWord, "translated to", translated)
		}

		resp := WordResp{ GopherWord: translated }
		if err := responses.JsonResponse(rw, &resp); err != nil {
			httpErr := serviceErr.NewHttpErr(http.StatusInternalServerError, "failed to write translate response", err)
			responses.ErrorResponse(rw, httpErr, log)
			return
		}

		translationCache.Add(wordReq.EnglishWord, resp.GopherWord)
	}
	return ret
}
