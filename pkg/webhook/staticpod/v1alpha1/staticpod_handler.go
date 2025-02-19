/*
Copyright 2023 The OpenYurt Authors.

Licensed under the Apache License, Version 2.0 (the License);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an AS IS BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"github.com/openyurtio/openyurt/pkg/apis/apps/v1alpha1"
	"github.com/openyurtio/openyurt/pkg/webhook/util"
)

// SetupWebhookWithManager sets up Cluster webhooks. 	mutate path, validatepath, error
func (webhook *StaticPodHandler) SetupWebhookWithManager(mgr ctrl.Manager) (string, string, error) {
	gvk, err := apiutil.GVKForObject(&v1alpha1.StaticPod{}, mgr.GetScheme())
	if err != nil {
		return "", "", err
	}
	return util.GenerateMutatePath(gvk),
		util.GenerateValidatePath(gvk),
		ctrl.NewWebhookManagedBy(mgr).
			For(&v1alpha1.StaticPod{}).
			WithDefaulter(webhook).
			WithValidator(webhook).
			Complete()
}

// +kubebuilder:webhook:path=/validate-apps-openyurt-io-v1alpha1-staticpod,mutating=false,failurePolicy=fail,sideEffects=None,admissionReviewVersions=v1;v1beta1,groups=apps.openyurt.io,resources=staticpods,verbs=create;update,versions=v1alpha1,name=validate.apps.v1alpha1.staticpod.openyurt.io
// +kubebuilder:webhook:path=/mutate-apps-openyurt-io-v1alpha1-staticpod,mutating=true,failurePolicy=fail,sideEffects=None,admissionReviewVersions=v1;v1beta1,groups=apps.openyurt.io,resources=staticpods,verbs=create;update,versions=v1alpha1,name=mutate.apps.v1alpha1.staticpod.openyurt.io

// Cluster implements a validating and defaulting webhook for Cluster.
type StaticPodHandler struct {
}

var _ webhook.CustomDefaulter = &StaticPodHandler{}
var _ webhook.CustomValidator = &StaticPodHandler{}
