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

type SentenceReq struct {
	EnglishSentence string `json:"english-sentence"`
}

type SentenceResp struct {
	GopherSentence string `json:"gopher-sentence"`
}

func SentenceHandler(log logger.Logger, translationCache cache.KeyValueCache) http.HandlerFunc {
	ret := func(rw http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		var sentenceReq SentenceReq
		jd := json.NewDecoder(req.Body)
		if err := jd.Decode(&sentenceReq); err != nil {
			responses.ErrorResponse(rw, serviceErr.NewHttpErr(http.StatusBadRequest, "failed to decode body", err), log)
			return
		}

		sentenceReq.EnglishSentence = strings.TrimSpace(sentenceReq.EnglishSentence)
		if sentenceReq.EnglishSentence == "" {
			err := serviceErr.NewHttpErr(http.StatusBadRequest,
				"Please provide english-sentence in request body",
				errors.New("invalid request"))
			responses.ErrorResponse(rw, err, log)
			return
		}

		words := strings.Split(sentenceReq.EnglishSentence, " ")
		if !validators.IsValidEnglishSentence(words) {
			err := serviceErr.NewHttpErr(http.StatusBadRequest,
				fmt.Sprintf("'%s' is NOT a valid english sentence", sentenceReq.EnglishSentence),
				errors.New("invalid english-sentence"))
			responses.ErrorResponse(rw, err, log)
			return
		}

		var translated string
		if fromCache := translationCache.Get(sentenceReq.EnglishSentence); fromCache != "" {
			translated = fromCache
			log.Infoln(sentenceReq.EnglishSentence, "sentence translating from cache", translated)
		} else {
			translated = language.TranslateSentence(words)
			log.Infoln(sentenceReq.EnglishSentence, "sentence translated to", translated)
		}

		resp := SentenceResp{ GopherSentence: translated }
		if err := responses.JsonResponse(rw, &resp); err != nil {
			httpErr := serviceErr.NewHttpErr(http.StatusInternalServerError,
				"failed to write sentance translate response", err)
			responses.ErrorResponse(rw, httpErr, log)
			return
		}

		translationCache.Add(sentenceReq.EnglishSentence, resp.GopherSentence)
	}
	return ret
}
