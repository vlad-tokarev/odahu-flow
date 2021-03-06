//
//    Copyright 2019 EPAM Systems
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package routes

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	odahuflow_errors "github.com/odahu/odahu-flow/packages/operator/pkg/errors"
	k8_serror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

const (
	MaxSize                 = 500
	FirstPage               = 0
	SizeURLParamName        = "size"
	PageURLParamName        = "page"
	DisabledAPIErrorMessage = "This API is disabled"
)

type HTTPResult struct {
	// Success of error message
	Message string `json:"message"`
}

// We should develop a custom exception on the repository layer.
// But we rely on kubernetes exceptions for now.
// TODO: implement Odahuflow exceptions
func CalculateHTTPStatusCode(err error) int {
	errStatus, ok := err.(*k8_serror.StatusError)
	if ok {
		return int(errStatus.ErrStatus.Code)
	}

	if odahuflow_errors.IsNotFoundError(err) {
		return http.StatusNotFound
	}

	if odahuflow_errors.IsAlreadyExistError(err) {
		return http.StatusConflict
	}

	if odahuflow_errors.IsForbiddenError(err) {
		return http.StatusForbidden
	}

	return http.StatusInternalServerError
}

func URLParamsToFilter(c *gin.Context, filter interface{}, fields map[string]int) (size int, page int, err error) {
	urlParameters := c.Request.URL.Query()
	size = MaxSize
	page = FirstPage

	for name, value := range urlParameters {
		switch name {
		case SizeURLParamName:
			if len(value) > 1 {
				return size, page, errors.New("the size URL parameter must be only one")
			}
			size, err = strconv.Atoi(value[0])
			if err != nil {
				return size, page, err
			}
		case PageURLParamName:
			if len(value) > 1 {
				return size, page, errors.New("the page URL parameter must be only one")
			}
			page, err = strconv.Atoi(value[0])
			if err != nil {
				return size, page, err
			}
		default:
			fieldNumber, ok := fields[name]
			if !ok {
				return size, page, fmt.Errorf("cannot find %s url parameter", name)
			}

			reflect.ValueOf(filter).Elem().Field(fieldNumber).Set(reflect.ValueOf(value))
		}
	}

	return size, page, nil
}

func DisableAPIMiddleware(enabledAPI bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !enabledAPI && c.Request.Method != http.MethodGet {
			c.AbortWithStatusJSON(http.StatusBadRequest, HTTPResult{Message: DisabledAPIErrorMessage})
		}
	}
}

// Because k8s has only "seconds" precision therefore we should operate the same precision in tests to
// compare timings in appropriate way
func GetTimeNowTruncatedToSeconds() metav1.Time {
	t1 := time.Now()
	t2 := t1.Truncate(time.Second)
	return metav1.Time{Time: t2}
}
