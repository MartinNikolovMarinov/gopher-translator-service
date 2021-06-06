package responses

import (
	"encoding/json"
	"net/http"

	"github.com/gopher-translator-service/internal/errors"
	"github.com/gopher-translator-service/pkg/logger"
	"github.com/gopher-translator-service/pkg/util"
)

func ErrorResponse(rw http.ResponseWriter, err error, log logger.Logger) error {
	if err == nil {
		return nil
	}

	if log != nil {
		log.Warnf("[Server Responded With Error] %s\n", err.Error())
	}

	switch persedErr := err.(type) {
	case *errors.HttpErr:
		rw.WriteHeader(persedErr.Status)
	default:
		rw.WriteHeader(http.StatusInternalServerError)
	}

	errBytes, err := json.Marshal(err)
	if err != nil {
		return err
	}
	if err := util.WriteAllBytes(rw, errBytes); err != nil {
		return err
	}

	return nil
}

func JsonResponse(rw http.ResponseWriter, data interface{}) error {
	if data == nil {
		return nil
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := util.WriteAllBytes(rw, dataBytes); err != nil {
		return err
	}

	return nil
}
