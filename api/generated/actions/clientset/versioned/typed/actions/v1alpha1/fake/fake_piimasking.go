/*
Copyright 2022.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/odigos-io/odigos/api/actions/v1alpha1"
	actionsv1alpha1 "github.com/odigos-io/odigos/api/generated/actions/applyconfiguration/actions/v1alpha1"
	typedactionsv1alpha1 "github.com/odigos-io/odigos/api/generated/actions/clientset/versioned/typed/actions/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakePiiMaskings implements PiiMaskingInterface
type fakePiiMaskings struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.PiiMasking, *v1alpha1.PiiMaskingList, *actionsv1alpha1.PiiMaskingApplyConfiguration]
	Fake *FakeActionsV1alpha1
}

func newFakePiiMaskings(fake *FakeActionsV1alpha1, namespace string) typedactionsv1alpha1.PiiMaskingInterface {
	return &fakePiiMaskings{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.PiiMasking, *v1alpha1.PiiMaskingList, *actionsv1alpha1.PiiMaskingApplyConfiguration](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("piimaskings"),
			v1alpha1.SchemeGroupVersion.WithKind("PiiMasking"),
			func() *v1alpha1.PiiMasking { return &v1alpha1.PiiMasking{} },
			func() *v1alpha1.PiiMaskingList { return &v1alpha1.PiiMaskingList{} },
			func(dst, src *v1alpha1.PiiMaskingList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.PiiMaskingList) []*v1alpha1.PiiMasking { return gentype.ToPointerSlice(list.Items) },
			func(list *v1alpha1.PiiMaskingList, items []*v1alpha1.PiiMasking) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
