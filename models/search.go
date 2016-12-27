package models

import (
	"errors"
	"net/http"
	"strconv"
	"fmt"
)

const (
	RADIUS = 1000 // search in meters
	LIMIT = 20
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
		if limit > 100 {
			return sp, errors.New("Max limit is exceeded")
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
			} else if rad > 10000 {
				return sp, errors.New("Max radius is exceeded")
			} else {
				sp.Rad = rad
			}
		} else {
			sp.Rad = RADIUS
		}
	}

	return sp, nil
}

func (p *SearchParams) getPrev() string {
	if p.Offset == 0 {
		return ""
	}
	newOffset := p.Offset - p.Limit
	if newOffset < 0 {
		newOffset = 0
	}
	if p.IsGeo {
		return fmt.Sprintf("?lon=%g&lat=%g&rad=%g&limit=%d&offset=%d",
			p.Lon,
			p.Lat,
			p.Rad,
			p.Limit,
			newOffset,
		)
	} else {
		return fmt.Sprintf("?limit=%d&offset=%d",
			p.Limit,
			newOffset,
		)
	}
}

func (p *SearchParams) getNext() string {
	newOffset := p.Offset + p.Limit
	if p.IsGeo {
		return fmt.Sprintf("?lon=%g&lat=%g&rad=%g&limit=%d&offset=%d",
			p.Lon,
			p.Lat,
			p.Rad,
			p.Limit,
			newOffset,
		)
	} else {
		return fmt.Sprintf("?limit=%d&offset=%d",
			p.Limit,
			newOffset,
		)
	}
}