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

package deployment

import (
	"errors"
	"fmt"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis/deployment"
	md_repository "github.com/odahu/odahu-flow/packages/operator/pkg/repository/deployment"
	"github.com/odahu/odahu-flow/packages/operator/pkg/validation"
	"go.uber.org/multierr"
	"strings"
)

const (
	URLPrefixEmptyErrorMessage = "URL Prefix must not be empty"
	EmptyTargetErrorMessage    = "model deployment targets must contain at least one element"
	OneTargetErrorMessage      = "it must have 100 weight or nil value if there is only one target"
	MissedWeightErrorMessage   = "weights must be present if there are more than one model deployment targets"
	TotalWeightErrorMessage    = "total target weight does not equal 100"
	URLPrefixSlashErrorMessage = "the URL prefix must start with slash"
	ForbiddenPrefix            = "the URL prefix %s is forbidden"
	ErrorMessageTemplate       = "%s: %s"
	ValidationMrErrorMessage   = "Validation of model route is failed"
)

var (
	MaxWeight         = int32(100)
	ForbiddenPrefixes = []string{
		"/model", "/feedback",
	}
)

type MrValidator struct {
	mdRepository md_repository.Repository
}

func NewMrValidator(mdRepository md_repository.Repository) *MrValidator {
	return &MrValidator{
		mdRepository: mdRepository,
	}
}

func (mrv *MrValidator) ValidatesAndSetDefaults(mr *deployment.ModelRoute) (err error) {
	err = multierr.Append(err, validation.ValidateID(mr.ID))

	err = multierr.Append(err, mrv.validateMainParameters(mr))

	err = multierr.Append(err, mrv.validateModelDeploymentTargets(mr))

	return
}

func (mrv *MrValidator) validateMainParameters(mr *deployment.ModelRoute) (err error) {
	if len(mr.Spec.URLPrefix) == 0 {
		err = multierr.Append(err, errors.New(URLPrefixEmptyErrorMessage))
	} else {
		if !strings.HasPrefix(mr.Spec.URLPrefix, "/") {
			err = multierr.Append(err, errors.New(URLPrefixSlashErrorMessage))
		} else {
			for _, prefix := range ForbiddenPrefixes {
				if strings.HasPrefix(mr.Spec.URLPrefix, prefix) {
					err = multierr.Append(err, fmt.Errorf(ForbiddenPrefix, prefix))
					break
				}
			}
		}
	}
	if mr.Spec.Mirror != nil && len(*mr.Spec.Mirror) != 0 {
		if _, k8sError := mrv.mdRepository.GetModelDeployment(*mr.Spec.Mirror); k8sError != nil {
			err = multierr.Append(err, k8sError)
		}
	}

	if err != nil {
		return fmt.Errorf(ErrorMessageTemplate, ValidationMrErrorMessage, err.Error())
	}

	return
}

func (mrv *MrValidator) validateModelDeploymentTargets(mr *deployment.ModelRoute) (err error) {
	switch len(mr.Spec.ModelDeploymentTargets) {
	case 0:
		err = multierr.Append(err, errors.New(EmptyTargetErrorMessage))
	case 1:
		mdt := mr.Spec.ModelDeploymentTargets[0]

		if _, k8sError := mrv.mdRepository.GetModelDeployment(mdt.Name); k8sError != nil {
			err = multierr.Append(err, k8sError)
		}
		if mdt.Weight == nil {
			logMR.Info("Weight parameter is nil. Set the default value",
				"Model Route name", mr.ID, "weight", MaxWeight)
			mr.Spec.ModelDeploymentTargets[0].Weight = &MaxWeight
		} else if *mdt.Weight != 100 {
			err = multierr.Append(err, errors.New(OneTargetErrorMessage))
		}
	default:
		weightSum := int32(0)

		for _, mdt := range mr.Spec.ModelDeploymentTargets {
			if _, k8sError := mrv.mdRepository.GetModelDeployment(mdt.Name); k8sError != nil {
				err = multierr.Append(err, k8sError)
			}

			if mdt.Weight == nil {
				err = multierr.Append(err, errors.New(MissedWeightErrorMessage))
				continue
			}

			weightSum += *mdt.Weight
		}

		if weightSum != 100 {
			err = multierr.Append(err, errors.New(TotalWeightErrorMessage))
		}
	}

	return err
}
