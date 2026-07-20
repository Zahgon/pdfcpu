package sign

import "encoding/asn1"

var (
	oidETSIQCPublicWithSSCD    = asn1.ObjectIdentifier{0, 4, 0, 1456, 1, 1}
	oidSigPolicy               = asn1.ObjectIdentifier{0, 4, 0, 2023, 1, 1}
	oidQualSealPolicy          = asn1.ObjectIdentifier{0, 4, 0, 2023, 1, 2}
	oidAdvSigPolicy            = asn1.ObjectIdentifier{0, 4, 0, 2023, 1, 3}
	oidAdvSigLTVPolicy         = asn1.ObjectIdentifier{0, 4, 0, 2023, 1, 4}
	oidQESLTVPolicy            = asn1.ObjectIdentifier{0, 4, 0, 2023, 1, 5}
	oidQCESign                 = asn1.ObjectIdentifier{0, 4, 0, 194112, 1, 2}
	oidQCESeal                 = asn1.ObjectIdentifier{0, 4, 0, 194112, 1, 3}
	oidQWebAuthCert            = asn1.ObjectIdentifier{0, 4, 0, 194112, 1, 4}
	oidRSAESOAEP               = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 10}
	oidData                    = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 1}
	oidMessageDigest           = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 4}
	oidSigningTime             = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 5}
	oidTSTInfo                 = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 1, 4}
	oidSigningCertificate      = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 2, 12}
	oidTimestampToken          = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 2, 14}
	oidSigPolicyID             = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 2, 15}
	oidCommitmentType          = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 2, 16}
	oidContentTimestamp        = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 2, 20}
	oidCompleteCertificateRefs = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 2, 21}
	oidCompleteRevocationRefs  = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 2, 22}
	oidCertificateValues       = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 2, 23}
	oidRevocationValues        = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 2, 24}
	oidArchiveTimestamp        = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 2, 27}
	oidSigningCertificateV2    = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 2, 47}
	oidProofOfOrigin           = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 6, 1}
	oidProofOfReceipt          = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 6, 2}
	oidProofOfDelivery         = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 6, 3}
	oidProofOfSender           = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 6, 4}
	oidProofOfApproval         = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 6, 5}
	oidProofOfCreation         = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 6, 6}
	oidRevocationInfoArchival  = asn1.ObjectIdentifier{1, 2, 840, 113583, 1, 1, 8}
	oidOCSPNoCheck             = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 48, 1, 5}
)
