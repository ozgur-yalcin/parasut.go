# parasut
An easy-to-use parasut.com API (Paraşüt v4) with golang

```go
package main

import (
	"encoding/json"
	"fmt"
	"parasut/src"
)

func main() {
	api := parasut.API{}
	auth := api.Authorize()
	if auth {
		request := parasut.Request{}
		request.SalesInvoices.Data.ID = "" // Satış Faturası ID
		invoice := api.ShowSalesInvoice(request)
		docid := invoice.SalesInvoices.Data.RelationShips.ActiveEDocument.Data.ID
		doctype := invoice.SalesInvoices.Data.RelationShips.ActiveEDocument.Data.Type
		if doctype == "e_archives" {
			request := parasut.Request{}
			request.EArchives.Data.ID = docid
			earchive := api.ShowEArchive(request)
			pretty, _ := json.MarshalIndent(earchive.EArchives, " ", "\t")
			fmt.Println(string(pretty))
			pdf := api.ShowEArchivePDF(docid)
			pdfurl := pdf.EArchivePDF.Data.Attr.URL
			fmt.Println(pdfurl)
		}
		if doctype == "e_invoices" {
			request := parasut.Request{}
			request.EInvoices.Data.ID = docid
			einvoice := api.ShowEInvoice(request)
			pretty, _ := json.MarshalIndent(einvoice.EInvoices, " ", "\t")
			fmt.Println(string(pretty))
			pdf := api.ShowEInvoicePDF(docid)
			pdfurl := pdf.EInvoicePDF.Data.Attr.URL
			fmt.Println(pdfurl)
		}
	}
}

