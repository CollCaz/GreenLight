package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var sytaxError *json.SyntaxError
		var unmarshalTypeErro *json.UnmarshalTypeError
		var invalidUnmarshalErro *json.InvalidUnmarshalError
		var maxBytesErro *http.MaxBytesError

		switch {
		case errors.As(err, &sytaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", sytaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeErro):
			if unmarshalTypeErro.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeErro.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeErro.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body msut not be empty")

		case strings.HasPrefix(err.Error(), "json: unkown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unkown field ")
			return fmt.Errorf("body contains unkown key %s", fieldName)

		case errors.As(err, &maxBytesErro):
			return fmt.Errorf("body must not be larger the %d bytes", maxBytesErro.Limit)

		case errors.As(err, &invalidUnmarshalErro):
			panic(err)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}
	return nil
}
