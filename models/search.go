package models

import (
	"errors"
	"net/http"
	"strconv"
)

const (
	RADIUS = 500 // search in meters
	LIMIT = 100
	OFFSET = 0
)

type (
	SearchParams struct {
		Lon    float64
		Lat    float64
		Rad    float64

		Limit  int
		Offset int

		IsGeo  bool
	}
)

func NewSearchParams(r *http.Request) (*SearchParams, error) {

	sp := &SearchParams{Limit: LIMIT, Offset: OFFSET, IsGeo: false}

	queryParams := r.URL.Query()

	limitParam := queryParams.Get("limit")
	if len(limitParam) != 0 {
		limit, err := strconv.Atoi(limitParam)
		if err != nil || limit < 0 {
			return sp, errors.New("limit is incorrect")
		}
		sp.Limit = limit
	}

	offsetParam := queryParams.Get("offset")
	if len(offsetParam) != 0 {
		offset, err := strconv.Atoi(offsetParam)
		if err != nil || offset < 0 {
			return sp, errors.New("offset is incorrect")
		}
		sp.Offset = offset
	}

	lonParam := queryParams.Get("lon")
	latParam := queryParams.Get("lat")
	radParam := queryParams.Get("rad")

	if len(lonParam) != 0 || len(latParam) != 0 {
		sp.IsGeo = true
		if len(lonParam) == 0 {
			return sp, errors.New("lon is missed")
		}
		if len(latParam) == 0 {
			return sp, errors.New("lat is missed")
		}
		if lon, err := strconv.ParseFloat(lonParam, 64); err != nil {
			return sp, errors.New("lon is incorrect")
		} else {
			sp.Lon = lon
		}
		if lat, err := strconv.ParseFloat(latParam, 64); err != nil {
			return sp, errors.New("lat is incorrect")
		} else {
			sp.Lat = lat
		}

		if len(radParam) != 0 {
			if rad, err := strconv.ParseFloat(latParam, 64); err != nil || rad < 0 {
				return sp, errors.New("rad is incorrect")
			} else {
				sp.Rad = rad
			}
		} else {
			sp.Rad = RADIUS
		}
	}

	return sp, nil
}
