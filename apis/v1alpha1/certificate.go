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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CertificateSpec defines the desired state of Certificate.
type CertificateSpec struct {

	// The Amazon Resource Name (ARN) of the private certificate authority (CA)
	// that will be used to issue the certificate. If you do not provide an ARN
	// and you are trying to request a private certificate, ACM will attempt to
	// issue a public certificate. For more information about private CAs, see the
	// Amazon Web Services Private Certificate Authority (https://docs.aws.amazon.com/privateca/latest/userguide/PcaWelcome.html)
	// user guide. The ARN must have the following form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	CertificateAuthorityARN *string `json:"certificateAuthorityARN,omitempty"`
	// Fully qualified domain name (FQDN), such as www.example.com, that you want
	// to secure with an ACM certificate. Use an asterisk (*) to create a wildcard
	// certificate that protects several sites in the same domain. For example,
	// *.example.com protects www.example.com, site.example.com, and images.example.com.
	//
	// In compliance with RFC 5280 (https://datatracker.ietf.org/doc/html/rfc5280),
	// the length of the domain name (technically, the Common Name) that you provide
	// cannot exceed 64 octets (characters), including periods. To add a longer
	// domain name, specify it in the Subject Alternative Name field, which supports
	// names up to 253 octets in length.
	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName"`
	// The domain name that you want ACM to use to send you emails so that you can
	// validate domain ownership.
	DomainValidationOptions []*DomainValidationOption `json:"domainValidationOptions,omitempty"`
	// Specifies the algorithm of the public and private key pair that your certificate
	// uses to encrypt data. RSA is the default key algorithm for ACM certificates.
	// Elliptic Curve Digital Signature Algorithm (ECDSA) keys are smaller, offering
	// security comparable to RSA keys but with greater computing efficiency. However,
	// ECDSA is not supported by all network clients. Some AWS services may require
	// RSA keys, or only support ECDSA keys of a particular size, while others allow
	// the use of either RSA and ECDSA keys to ensure that compatibility is not
	// broken. Check the requirements for the AWS service where you plan to deploy
	// your certificate.
	//
	// Default: RSA_2048
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty"`
	// Currently, you can use this parameter to specify whether to add the certificate
	// to a certificate transparency log. Certificate transparency makes it possible
	// to detect SSL/TLS certificates that have been mistakenly or maliciously issued.
	// Certificates that have not been logged typically produce an error message
	// in a browser. For more information, see Opting Out of Certificate Transparency
	// Logging (https://docs.aws.amazon.com/acm/latest/userguide/acm-bestpractices.html#best-practices-transparency).
	Options *CertificateOptions `json:"options,omitempty"`
	// Additional FQDNs to be included in the Subject Alternative Name extension
	// of the ACM certificate. For example, add the name www.example.net to a certificate
	// for which the DomainName field is www.example.com if users can reach your
	// site by using either name. The maximum number of domain names that you can
	// add to an ACM certificate is 100. However, the initial quota is 10 domain
	// names. If you need more than 10 names, you must request a quota increase.
	// For more information, see Quotas (https://docs.aws.amazon.com/acm/latest/userguide/acm-limits.html).
	//
	// The maximum length of a SAN DNS name is 253 octets. The name is made up of
	// multiple labels separated by periods. No label can be longer than 63 octets.
	// Consider the following examples:
	//
	//   - (63 octets).(63 octets).(63 octets).(61 octets) is legal because the
	//     total length is 253 octets (63+1+63+1+63+1+61) and no label exceeds 63
	//     octets.
	//
	//   - (64 octets).(63 octets).(63 octets).(61 octets) is not legal because
	//     the total length exceeds 253 octets (64+1+63+1+63+1+61) and the first
	//     label exceeds 63 octets.
	//
	//   - (63 octets).(63 octets).(63 octets).(62 octets) is not legal because
	//     the total length of the DNS name (63+1+63+1+63+1+62) exceeds 253 octets.
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty"`
	// One or more resource tags to associate with the certificate.
	Tags []*Tag `json:"tags,omitempty"`
}

// CertificateStatus defines the observed state of Certificate
type CertificateStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The time at which the certificate was requested.
	// +kubebuilder:validation:Optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// Contains a list of Extended Key Usage X.509 v3 extension objects. Each object
	// specifies a purpose for which the certificate public key can be used and
	// consists of a name and an object identifier (OID).
	// +kubebuilder:validation:Optional
	ExtendedKeyUsages []*ExtendedKeyUsage `json:"extendedKeyUsages,omitempty"`
	// The reason the certificate request failed. This value exists only when the
	// certificate status is FAILED. For more information, see Certificate Request
	// Failed (https://docs.aws.amazon.com/acm/latest/userguide/troubleshooting.html#troubleshooting-failed)
	// in the Certificate Manager User Guide.
	// +kubebuilder:validation:Optional
	FailureReason *string `json:"failureReason,omitempty"`
	// The date and time when the certificate was imported. This value exists only
	// when the certificate type is IMPORTED.
	// +kubebuilder:validation:Optional
	ImportedAt *metav1.Time `json:"importedAt,omitempty"`
	// A list of ARNs for the Amazon Web Services resources that are using the certificate.
	// A certificate can be used by multiple Amazon Web Services resources.
	// +kubebuilder:validation:Optional
	InUseBy []*string `json:"inUseBy,omitempty"`
	// The time at which the certificate was issued. This value exists only when
	// the certificate type is AMAZON_ISSUED.
	// +kubebuilder:validation:Optional
	IssuedAt *metav1.Time `json:"issuedAt,omitempty"`
	// The name of the certificate authority that issued and signed the certificate.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty"`
	// A list of Key Usage X.509 v3 extension objects. Each object is a string value
	// that identifies the purpose of the public key contained in the certificate.
	// Possible extension values include DIGITAL_SIGNATURE, KEY_ENCHIPHERMENT, NON_REPUDIATION,
	// and more.
	// +kubebuilder:validation:Optional
	KeyUsages []*KeyUsage `json:"keyUsages,omitempty"`
	// The time after which the certificate is not valid.
	// +kubebuilder:validation:Optional
	NotAfter *metav1.Time `json:"notAfter,omitempty"`
	// The time before which the certificate is not valid.
	// +kubebuilder:validation:Optional
	NotBefore *metav1.Time `json:"notBefore,omitempty"`
	// Specifies whether the certificate is eligible for renewal. At this time,
	// only exported private certificates can be renewed with the RenewCertificate
	// command.
	// +kubebuilder:validation:Optional
	RenewalEligibility *string `json:"renewalEligibility,omitempty"`
	// Contains information about the status of ACM's managed renewal (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html)
	// for the certificate. This field exists only when the certificate type is
	// AMAZON_ISSUED.
	// +kubebuilder:validation:Optional
	RenewalSummary *RenewalSummary `json:"renewalSummary,omitempty"`
	// The reason the certificate was revoked. This value exists only when the certificate
	// status is REVOKED.
	// +kubebuilder:validation:Optional
	RevocationReason *string `json:"revocationReason,omitempty"`
	// The time at which the certificate was revoked. This value exists only when
	// the certificate status is REVOKED.
	// +kubebuilder:validation:Optional
	RevokedAt *metav1.Time `json:"revokedAt,omitempty"`
	// The serial number of the certificate.
	// +kubebuilder:validation:Optional
	Serial *string `json:"serial,omitempty"`
	// The algorithm that was used to sign the certificate.
	// +kubebuilder:validation:Optional
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
	// The status of the certificate.
	//
	// A certificate enters status PENDING_VALIDATION upon being requested, unless
	// it fails for any of the reasons given in the troubleshooting topic Certificate
	// request fails (https://docs.aws.amazon.com/acm/latest/userguide/troubleshooting-failed.html).
	// ACM makes repeated attempts to validate a certificate for 72 hours and then
	// times out. If a certificate shows status FAILED or VALIDATION_TIMED_OUT,
	// delete the request, correct the issue with DNS validation (https://docs.aws.amazon.com/acm/latest/userguide/dns-validation.html)
	// or Email validation (https://docs.aws.amazon.com/acm/latest/userguide/email-validation.html),
	// and try again. If validation succeeds, the certificate enters status ISSUED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
	// The name of the entity that is associated with the public key contained in
	// the certificate.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty"`
	// The source of the certificate. For certificates provided by ACM, this value
	// is AMAZON_ISSUED. For certificates that you imported with ImportCertificate,
	// this value is IMPORTED. ACM does not provide managed renewal (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html)
	// for imported certificates. For more information about the differences between
	// certificates that you import and those that ACM provides, see Importing Certificates
	// (https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html)
	// in the Certificate Manager User Guide.
	// +kubebuilder:validation:Optional
	Type *string `json:"type_,omitempty"`
}

// Certificate is the Schema for the Certificates API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec,omitempty"`
	Status            CertificateStatus `json:"status,omitempty"`
}

// CertificateList contains a list of Certificate
// +kubebuilder:object:root=true
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}