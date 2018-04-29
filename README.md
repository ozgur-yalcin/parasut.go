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

	// Satış faturası bilgilerine ulaşmak için
	if auth {
		request := parasut.Request{}
		request.SalesInvoices.Data.ID = "" // Satış Faturası ID
		invoice := api.ShowSalesInvoice(request)
		pretty, _ := json.MarshalIndent(invoice.SalesInvoices, " ", "\t")
		fmt.Println(string(pretty))
	}

	// Resmileştirilmiş fatura bilgilerine ulaşmak için
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
		}
		if doctype == "e_invoices" {
			request := parasut.Request{}
			request.EInvoices.Data.ID = docid
			einvoice := api.ShowEInvoice(request)
			pretty, _ := json.MarshalIndent(einvoice.EInvoices, " ", "\t")
			fmt.Println(string(pretty))
		}
	}

	// PDF url'si görüntülemek için
	if auth {
		request := parasut.Request{}
		request.SalesInvoices.Data.ID = "" // Satış Faturası ID
		invoice := api.ShowSalesInvoice(request)
		docid := invoice.SalesInvoices.Data.RelationShips.ActiveEDocument.Data.ID
		doctype := invoice.SalesInvoices.Data.RelationShips.ActiveEDocument.Data.Type
		if doctype == "e_archives" {
			request := parasut.Request{}
			request.EArchivePDF.Data.ID = docid
			pdf := api.ShowEArchivePDF(request)
			pdfurl := pdf.EArchivePDF.Data.Attr.URL
			fmt.Println(pdfurl)
		}
		if doctype == "e_invoices" {
			request := parasut.Request{}
			request.EInvoicePDF.Data.ID = docid
			pdf := api.ShowEInvoicePDF(request)
			pdfurl := pdf.EInvoicePDF.Data.Attr.URL
			fmt.Println(pdfurl)
		}
	}
}
