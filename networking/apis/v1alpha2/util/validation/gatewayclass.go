/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validation

import (
	validationutilv1beta1 "github.com/sharding-db/k8s-gateway-apis/networking/apis/v1beta1/util/validation"
)

// IsControllerNameValid checks that the provided controllerName complies with the expected
// format. It must be a non-empty domain prefixed path.
var IsControllerNameValid = validationutilv1beta1.IsControllerNameValid
