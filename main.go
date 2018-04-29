package main

import (
	"encoding/json"
	"fmt"
	"parasut/src"
)

func main() {
	request := parasut.Request{}
	auth := request.Authorize()
	if auth {
		invoiceid := "" // Satış Faturası ID
		invoice := request.ShowSalesInvoice(invoiceid)
		docid := invoice.SalesInvoices.Data.RelationShips.ActiveEDocument.Data.ID
		doctype := invoice.SalesInvoices.Data.RelationShips.ActiveEDocument.Data.Type
		if doctype == "e_archives" {
			earchive := request.ShowEArchive(docid)
			pretty, _ := json.MarshalIndent(earchive.EArchives, " ", "\t")
			fmt.Println(string(pretty))
			pdf := request.ShowEArchivePDF(docid)
			pdfurl := pdf.EArchivePDF.Data.Attr.URL
			fmt.Println(pdfurl)
		}
		if doctype == "e_invoices" {
			einvoice := request.ShowEInvoice(docid)
			pretty, _ := json.MarshalIndent(einvoice.EInvoices, " ", "\t")
			fmt.Println(string(pretty))
			pdf := request.ShowEInvoicePDF(docid)
			pdfurl := pdf.EInvoicePDF.Data.Attr.URL
			fmt.Println(pdfurl)
		}
	}
}
