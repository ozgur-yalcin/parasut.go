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

	// Müşteri kaydı oluşturmak için
	if auth {
		request := parasut.Request{}
		request.Contacts.Data.Type = "contacts"
		request.Contacts.Data.Attr.AccountType = "customer" // customer (Müşteri) || supplier (Tedarikçi)
		request.Contacts.Data.Attr.Name = ""                // Firma Ünvanı
		request.Contacts.Data.Attr.ShortName = ""           // Kısa İsim
		request.Contacts.Data.Attr.ContactType = ""         // company (Şirket) || person (Şahıs)
		request.Contacts.Data.Attr.TaxNumber = ""           // Vergi Numarası
		request.Contacts.Data.Attr.TaxOffice = ""           // Vergi Dairesi
		request.Contacts.Data.Attr.City = ""                // İl
		request.Contacts.Data.Attr.District = ""            // İlçe
		request.Contacts.Data.Attr.Address = ""             // Adres
		request.Contacts.Data.Attr.Phone = ""               // Telefon
		request.Contacts.Data.Attr.Fax = ""                 // Faks
		request.Contacts.Data.Attr.Email = ""               // Email
		request.Contacts.Data.Attr.IBAN = ""                // IBAN numarası
		invoice := api.CreateContact(request)
		pretty, _ := json.MarshalIndent(invoice.Contacts, " ", "\t")
		fmt.Println(string(pretty))
	}
}
