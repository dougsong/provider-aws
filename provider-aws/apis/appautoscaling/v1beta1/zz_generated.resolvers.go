/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1"
	common "github.com/upbound/official-providers/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
		Extract:      resource.ExtractParamPath("resource_id", false),
		Reference:    mg.Spec.ForProvider.ResourceIDRef,
		Selector:     mg.Spec.ForProvider.ResourceIDSelector,
		To: reference.To{
			List:    &TargetList{},
			Managed: &Target{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScalableDimension),
		Extract:      resource.ExtractParamPath("scalable_dimension", false),
		Reference:    mg.Spec.ForProvider.ScalableDimensionRef,
		Selector:     mg.Spec.ForProvider.ScalableDimensionSelector,
		To: reference.To{
			List:    &TargetList{},
			Managed: &Target{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScalableDimension")
	}
	mg.Spec.ForProvider.ScalableDimension = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScalableDimensionRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceNamespace),
		Extract:      resource.ExtractParamPath("service_namespace", false),
		Reference:    mg.Spec.ForProvider.ServiceNamespaceRef,
		Selector:     mg.Spec.ForProvider.ServiceNamespaceSelector,
		To: reference.To{
			List:    &TargetList{},
			Managed: &Target{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceNamespace")
	}
	mg.Spec.ForProvider.ServiceNamespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceNamespaceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ScheduledAction.
func (mg *ScheduledAction) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
		Extract:      resource.ExtractParamPath("resource_id", false),
		Reference:    mg.Spec.ForProvider.ResourceIDRef,
		Selector:     mg.Spec.ForProvider.ResourceIDSelector,
		To: reference.To{
			List:    &TargetList{},
			Managed: &Target{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScalableDimension),
		Extract:      resource.ExtractParamPath("scalable_dimension", false),
		Reference:    mg.Spec.ForProvider.ScalableDimensionRef,
		Selector:     mg.Spec.ForProvider.ScalableDimensionSelector,
		To: reference.To{
			List:    &TargetList{},
			Managed: &Target{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScalableDimension")
	}
	mg.Spec.ForProvider.ScalableDimension = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScalableDimensionRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceNamespace),
		Extract:      resource.ExtractParamPath("service_namespace", false),
		Reference:    mg.Spec.ForProvider.ServiceNamespaceRef,
		Selector:     mg.Spec.ForProvider.ServiceNamespaceSelector,
		To: reference.To{
			List:    &TargetList{},
			Managed: &Target{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceNamespace")
	}
	mg.Spec.ForProvider.ServiceNamespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceNamespaceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Target.
func (mg *Target) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}
