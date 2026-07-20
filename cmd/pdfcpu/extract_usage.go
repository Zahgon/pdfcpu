package main

const usageLongExtract = `Export inFile's images, fonts, content or pages into outDir.

      mode ... extraction mode
     pages ... Please refer to "pdfcpu selectedpages"
    inFile ... input PDF file, use - to read from stdin
    outDir ... output directory, use - with mode page and one selected page to write to stdout

 The extraction modes are:

  image ... extract images
   font ... extract font files (supported font types: TrueType)
content ... extract raw page content
   page ... extract single page PDFs
   meta ... extract all metadata (page selection does not apply)

Pipeline example:
   aws s3 cp s3://acme-archive/contract.pdf - \
      | pdfcpu extract -m page -p 3 - - \
      | aws s3 cp - s3://acme-archive/pages/contract-page-3.pdf
`
