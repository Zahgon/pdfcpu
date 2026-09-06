package main

const (
	usageLongCertificates = `Manage certificates.

           inFile ... .pem, .p7c, .cer, .crt file
             json ... output JSON

   Standard builds start with an empty trusted certificate directory.
   Builds created with -tags pdfcpu_eutl initialize this directory with an
   embedded snapshot of EU Trusted List certificate bundles.

   Please import any missing certificates.
`

	usageLongCertificatesList = `List installed trusted certificates.

             json ... output JSON

   Certificates are read from pdfcpu's trusted certificate directory.

Examples:
   pdfcpu certificates list
   pdfcpu certificates list --json
`

	usageLongSignatures = `Manage digital signatures.

           all ... validate all signatures (certified, approval, usage rights, digital timestamps)
          full ... detailed output including certificate chains, revocation status and any problems encountered
        inFile ... input PDF file, use - to read from stdin
       outFile ... output PDF file, use - to write to stdout

       Related configuration parameters: timeoutCRL,
                                         timeoutOCSP,
                                         preferredCertRevocationChecker

Pipeline examples:
   aws s3 cp s3://acme-docs/signed.pdf - \
      | pdfcpu signatures validate -

   aws s3 cp s3://acme-docs/signed.pdf - \
      | pdfcpu signatures remove - - \
      | aws s3 cp - s3://acme-docs/unsigned.pdf
	`

	usageLongSignaturesValidate = `pdfcpu validates signature integrity, reports available trust evidence and performs
a best-effort local assessment.

This command validates signed byte ranges and supported CMS/PKCS#7 signature data,
and reports signer metadata.
It also reports certificate chains, timestamps, revocation responses,
DSS data and PAdES-B-B results from supported profile checks where available.

The legacy adbe.x509.rsa_sha1 and adbe.pkcs7.sha1 profiles are supported for validation of existing PDFs only.

Certificate-path and revocation checks use pdfcpu's configured local certificate
store and available CRL/OCSP evidence on a best-effort basis.

This command does not perform legal-validity, eIDAS, enterprise policy or full long-term validation (LTV).

           all ... validate all signatures (certified, approval, usage rights, digital timestamps)
          full ... detailed output including certificate chains, revocation status and any problems encountered
        inFile ... input PDF file, use - to read from stdin

        Related configuration parameters: timeoutCRL,
                                          timeoutOCSP,
                                          preferredCertRevocationChecker
`
)
