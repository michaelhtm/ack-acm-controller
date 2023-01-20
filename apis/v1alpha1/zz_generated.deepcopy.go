//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Certificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateDetail) DeepCopyInto(out *CertificateDetail) {
	*out = *in
	if in.CertificateARN != nil {
		in, out := &in.CertificateARN, &out.CertificateARN
		*out = new(string)
		**out = **in
	}
	if in.CertificateAuthorityARN != nil {
		in, out := &in.CertificateAuthorityARN, &out.CertificateAuthorityARN
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.DomainValidationOptions != nil {
		in, out := &in.DomainValidationOptions, &out.DomainValidationOptions
		*out = make([]*DomainValidation, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DomainValidation)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ExtendedKeyUsages != nil {
		in, out := &in.ExtendedKeyUsages, &out.ExtendedKeyUsages
		*out = make([]*ExtendedKeyUsage, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ExtendedKeyUsage)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(string)
		**out = **in
	}
	if in.ImportedAt != nil {
		in, out := &in.ImportedAt, &out.ImportedAt
		*out = (*in).DeepCopy()
	}
	if in.InUseBy != nil {
		in, out := &in.InUseBy, &out.InUseBy
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IssuedAt != nil {
		in, out := &in.IssuedAt, &out.IssuedAt
		*out = (*in).DeepCopy()
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	if in.KeyAlgorithm != nil {
		in, out := &in.KeyAlgorithm, &out.KeyAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.KeyUsages != nil {
		in, out := &in.KeyUsages, &out.KeyUsages
		*out = make([]*KeyUsage, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeyUsage)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.NotAfter != nil {
		in, out := &in.NotAfter, &out.NotAfter
		*out = (*in).DeepCopy()
	}
	if in.NotBefore != nil {
		in, out := &in.NotBefore, &out.NotBefore
		*out = (*in).DeepCopy()
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(CertificateOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.RenewalEligibility != nil {
		in, out := &in.RenewalEligibility, &out.RenewalEligibility
		*out = new(string)
		**out = **in
	}
	if in.RenewalSummary != nil {
		in, out := &in.RenewalSummary, &out.RenewalSummary
		*out = new(RenewalSummary)
		(*in).DeepCopyInto(*out)
	}
	if in.RevocationReason != nil {
		in, out := &in.RevocationReason, &out.RevocationReason
		*out = new(string)
		**out = **in
	}
	if in.RevokedAt != nil {
		in, out := &in.RevokedAt, &out.RevokedAt
		*out = (*in).DeepCopy()
	}
	if in.Serial != nil {
		in, out := &in.Serial, &out.Serial
		*out = new(string)
		**out = **in
	}
	if in.SignatureAlgorithm != nil {
		in, out := &in.SignatureAlgorithm, &out.SignatureAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(string)
		**out = **in
	}
	if in.SubjectAlternativeNames != nil {
		in, out := &in.SubjectAlternativeNames, &out.SubjectAlternativeNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateDetail.
func (in *CertificateDetail) DeepCopy() *CertificateDetail {
	if in == nil {
		return nil
	}
	out := new(CertificateDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateList) DeepCopyInto(out *CertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Certificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateList.
func (in *CertificateList) DeepCopy() *CertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateOptions) DeepCopyInto(out *CertificateOptions) {
	*out = *in
	if in.CertificateTransparencyLoggingPreference != nil {
		in, out := &in.CertificateTransparencyLoggingPreference, &out.CertificateTransparencyLoggingPreference
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateOptions.
func (in *CertificateOptions) DeepCopy() *CertificateOptions {
	if in == nil {
		return nil
	}
	out := new(CertificateOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	if in.CertificateAuthorityARN != nil {
		in, out := &in.CertificateAuthorityARN, &out.CertificateAuthorityARN
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.DomainValidationOptions != nil {
		in, out := &in.DomainValidationOptions, &out.DomainValidationOptions
		*out = make([]*DomainValidationOption, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DomainValidationOption)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.KeyAlgorithm != nil {
		in, out := &in.KeyAlgorithm, &out.KeyAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(CertificateOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.SubjectAlternativeNames != nil {
		in, out := &in.SubjectAlternativeNames, &out.SubjectAlternativeNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStatus) DeepCopyInto(out *CertificateStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.ExtendedKeyUsages != nil {
		in, out := &in.ExtendedKeyUsages, &out.ExtendedKeyUsages
		*out = make([]*ExtendedKeyUsage, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ExtendedKeyUsage)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(string)
		**out = **in
	}
	if in.ImportedAt != nil {
		in, out := &in.ImportedAt, &out.ImportedAt
		*out = (*in).DeepCopy()
	}
	if in.InUseBy != nil {
		in, out := &in.InUseBy, &out.InUseBy
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IssuedAt != nil {
		in, out := &in.IssuedAt, &out.IssuedAt
		*out = (*in).DeepCopy()
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	if in.KeyUsages != nil {
		in, out := &in.KeyUsages, &out.KeyUsages
		*out = make([]*KeyUsage, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeyUsage)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.NotAfter != nil {
		in, out := &in.NotAfter, &out.NotAfter
		*out = (*in).DeepCopy()
	}
	if in.NotBefore != nil {
		in, out := &in.NotBefore, &out.NotBefore
		*out = (*in).DeepCopy()
	}
	if in.RenewalEligibility != nil {
		in, out := &in.RenewalEligibility, &out.RenewalEligibility
		*out = new(string)
		**out = **in
	}
	if in.RenewalSummary != nil {
		in, out := &in.RenewalSummary, &out.RenewalSummary
		*out = new(RenewalSummary)
		(*in).DeepCopyInto(*out)
	}
	if in.RevocationReason != nil {
		in, out := &in.RevocationReason, &out.RevocationReason
		*out = new(string)
		**out = **in
	}
	if in.RevokedAt != nil {
		in, out := &in.RevokedAt, &out.RevokedAt
		*out = (*in).DeepCopy()
	}
	if in.Serial != nil {
		in, out := &in.Serial, &out.Serial
		*out = new(string)
		**out = **in
	}
	if in.SignatureAlgorithm != nil {
		in, out := &in.SignatureAlgorithm, &out.SignatureAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStatus.
func (in *CertificateStatus) DeepCopy() *CertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSummary) DeepCopyInto(out *CertificateSummary) {
	*out = *in
	if in.CertificateARN != nil {
		in, out := &in.CertificateARN, &out.CertificateARN
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.Exported != nil {
		in, out := &in.Exported, &out.Exported
		*out = new(bool)
		**out = **in
	}
	if in.ExtendedKeyUsages != nil {
		in, out := &in.ExtendedKeyUsages, &out.ExtendedKeyUsages
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.HasAdditionalSubjectAlternativeNames != nil {
		in, out := &in.HasAdditionalSubjectAlternativeNames, &out.HasAdditionalSubjectAlternativeNames
		*out = new(bool)
		**out = **in
	}
	if in.ImportedAt != nil {
		in, out := &in.ImportedAt, &out.ImportedAt
		*out = (*in).DeepCopy()
	}
	if in.InUse != nil {
		in, out := &in.InUse, &out.InUse
		*out = new(bool)
		**out = **in
	}
	if in.IssuedAt != nil {
		in, out := &in.IssuedAt, &out.IssuedAt
		*out = (*in).DeepCopy()
	}
	if in.KeyAlgorithm != nil {
		in, out := &in.KeyAlgorithm, &out.KeyAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.KeyUsages != nil {
		in, out := &in.KeyUsages, &out.KeyUsages
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NotAfter != nil {
		in, out := &in.NotAfter, &out.NotAfter
		*out = (*in).DeepCopy()
	}
	if in.NotBefore != nil {
		in, out := &in.NotBefore, &out.NotBefore
		*out = (*in).DeepCopy()
	}
	if in.RenewalEligibility != nil {
		in, out := &in.RenewalEligibility, &out.RenewalEligibility
		*out = new(string)
		**out = **in
	}
	if in.RevokedAt != nil {
		in, out := &in.RevokedAt, &out.RevokedAt
		*out = (*in).DeepCopy()
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SubjectAlternativeNameSummaries != nil {
		in, out := &in.SubjectAlternativeNameSummaries, &out.SubjectAlternativeNameSummaries
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSummary.
func (in *CertificateSummary) DeepCopy() *CertificateSummary {
	if in == nil {
		return nil
	}
	out := new(CertificateSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainValidation) DeepCopyInto(out *DomainValidation) {
	*out = *in
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.ResourceRecord != nil {
		in, out := &in.ResourceRecord, &out.ResourceRecord
		*out = new(ResourceRecord)
		(*in).DeepCopyInto(*out)
	}
	if in.ValidationDomain != nil {
		in, out := &in.ValidationDomain, &out.ValidationDomain
		*out = new(string)
		**out = **in
	}
	if in.ValidationEmails != nil {
		in, out := &in.ValidationEmails, &out.ValidationEmails
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ValidationMethod != nil {
		in, out := &in.ValidationMethod, &out.ValidationMethod
		*out = new(string)
		**out = **in
	}
	if in.ValidationStatus != nil {
		in, out := &in.ValidationStatus, &out.ValidationStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainValidation.
func (in *DomainValidation) DeepCopy() *DomainValidation {
	if in == nil {
		return nil
	}
	out := new(DomainValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainValidationOption) DeepCopyInto(out *DomainValidationOption) {
	*out = *in
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.ValidationDomain != nil {
		in, out := &in.ValidationDomain, &out.ValidationDomain
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainValidationOption.
func (in *DomainValidationOption) DeepCopy() *DomainValidationOption {
	if in == nil {
		return nil
	}
	out := new(DomainValidationOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedKeyUsage) DeepCopyInto(out *ExtendedKeyUsage) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OID != nil {
		in, out := &in.OID, &out.OID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedKeyUsage.
func (in *ExtendedKeyUsage) DeepCopy() *ExtendedKeyUsage {
	if in == nil {
		return nil
	}
	out := new(ExtendedKeyUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filters) DeepCopyInto(out *Filters) {
	*out = *in
	if in.ExtendedKeyUsage != nil {
		in, out := &in.ExtendedKeyUsage, &out.ExtendedKeyUsage
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KeyTypes != nil {
		in, out := &in.KeyTypes, &out.KeyTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KeyUsage != nil {
		in, out := &in.KeyUsage, &out.KeyUsage
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filters.
func (in *Filters) DeepCopy() *Filters {
	if in == nil {
		return nil
	}
	out := new(Filters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyUsage) DeepCopyInto(out *KeyUsage) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyUsage.
func (in *KeyUsage) DeepCopy() *KeyUsage {
	if in == nil {
		return nil
	}
	out := new(KeyUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenewalSummary) DeepCopyInto(out *RenewalSummary) {
	*out = *in
	if in.DomainValidationOptions != nil {
		in, out := &in.DomainValidationOptions, &out.DomainValidationOptions
		*out = make([]*DomainValidation, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DomainValidation)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.RenewalStatus != nil {
		in, out := &in.RenewalStatus, &out.RenewalStatus
		*out = new(string)
		**out = **in
	}
	if in.RenewalStatusReason != nil {
		in, out := &in.RenewalStatusReason, &out.RenewalStatusReason
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenewalSummary.
func (in *RenewalSummary) DeepCopy() *RenewalSummary {
	if in == nil {
		return nil
	}
	out := new(RenewalSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecord) DeepCopyInto(out *ResourceRecord) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecord.
func (in *ResourceRecord) DeepCopy() *ResourceRecord {
	if in == nil {
		return nil
	}
	out := new(ResourceRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}