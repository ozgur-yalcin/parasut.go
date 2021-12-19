[![license](https://img.shields.io/:license-mit-blue.svg)](https://github.com/ozgur-soft/parasut.go/blob/master/LICENSE.md)
[![documentation](https://pkg.go.dev/badge/github.com/ozgur-soft/parasut.go)](https://pkg.go.dev/github.com/ozgur-soft/parasut.go/src)

# Parasut.go
An easy-to-use parasut.com API (v4) with golang

# Installation
```bash
go get github.com/ozgur-soft/parasut.go
```

# Satış faturası kaydı oluşturma
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.SalesInvoices.Data.Type = "sales_invoices"         // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.Attributes.ItemType = "invoice" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.Attributes.Description = ""     // Fatura başlığı
		request.SalesInvoices.Data.Attributes.TaxNumber = ""       // Vergi numarası
		request.SalesInvoices.Data.Attributes.TaxOffice = ""       // Vergi dairesi
		request.SalesInvoices.Data.Attributes.IssueDate = ""       // Fatura tarihi (Yıl-Ay-Gün)
		request.SalesInvoices.Data.Attributes.Currency = "TRL"     // "TRL" || "USD" || "EUR" || "GBP" (Para birimi)
		request.SalesInvoices.Data.Attributes.BillingPhone = ""    // Telefon numarası
		request.SalesInvoices.Data.Attributes.BillingFax = ""      // Fax numarası
		request.SalesInvoices.Data.Attributes.BillingAddress = ""  // Fatura adresi
		request.SalesInvoices.Data.Attributes.Country = ""         // Ülke
		request.SalesInvoices.Data.Attributes.City = ""            // İl
		request.SalesInvoices.Data.Attributes.District = ""        // İlçe

		request.SalesInvoices.Data.Relationships.Contact = new(parasut.SingleRelationShip)
		request.SalesInvoices.Data.Relationships.Contact.Data = new(parasut.RelationShip)
		request.SalesInvoices.Data.Relationships.Contact.Data.Type = "contacts" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.Relationships.Contact.Data.ID = ""           // Müşteri ID

		request.SalesInvoices.Data.Relationships.Category = new(parasut.SingleRelationShip)
		request.SalesInvoices.Data.Relationships.Category.Data = new(parasut.RelationShip)
		request.SalesInvoices.Data.Relationships.Category.Data.Type = "item_categories" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.Relationships.Category.Data.ID = ""                  // Kategori ID (varsa)

		detail := request.SalesInvoices.Data.Relationships.Details.Detail
		detail.Type = "sales_invoice_details"     // << Değişiklik yapmayınız !
		detail.Attributes.Quantity = "0"          // Ürün miktarı
		detail.Attributes.UnitPrice = "0"         // Ürün birim fiyatı
		detail.Attributes.VatRate = "0"           // Ürün KDV oranı
		detail.Attributes.DiscountType = "amount" // "amount" || "percentage" (İndirim türü)
		detail.Attributes.DiscountValue = "0"     // İndirim oranı
		detail.Relationships.Product = new(parasut.SingleRelationShip)
		detail.Relationships.Product.Data = new(parasut.RelationShip)
		detail.Relationships.Product.Data.Type = "products" // << Değişiklik yapmayınız !
		detail.Relationships.Product.Data.ID = ""           // Ürün ID
		request.SalesInvoices.Data.Relationships.Details.Data = append(request.SalesInvoices.Data.Relationships.Details.Data, detail)

		response := api.CreateSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoices, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Satış faturası kaydını silme
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.SalesInvoices.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.ID = ""                 // Satış faturası ID
		response := api.DeleteSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoices, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Satış faturası kaydını iptal etme
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.SalesInvoices.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.ID = ""                 // Satış faturası ID
		response := api.CancelSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoices, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Satış faturası kaydını arşivleme
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.SalesInvoices.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.ID = ""                 // Satış faturası ID
		response := api.ArchiveSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoices, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Satış faturası kaydını arşivden çıkarma
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.SalesInvoices.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.ID = ""                 // Satış faturası ID
		response := api.UnarchiveSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoices, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Satış faturası kaydını görüntüleme
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.SalesInvoices.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.ID = ""                 // Satış faturası ID
		response := api.ShowSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoices, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Satış faturasına ödeme kaydı oluşturma
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.SalesInvoices.Data.Type = "sales_invoices"    // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.ID = ""                    // Satış faturası ID
		request.Payments.Data.Type = "payments"               // << Değişiklik yapmayınız !
		request.Payments.Data.Attributes.AccountID = ""       // Ödeme yapılan hesap ID
		request.Payments.Data.Attributes.Description = ""     // Ödeme açıklaması
		request.Payments.Data.Attributes.Date = ""            // Ödeme tarihi (Yıl-Ay-Gün)
		request.Payments.Data.Attributes.Amount = ""          // Ödeme tutarı
		request.Payments.Data.Attributes.Currency = "TRL"     // "TRL" || "USD" || "EUR" || "GBP" (Para birimi)
		request.Payments.Data.Attributes.ExchangeRate = "1.0" // Döviz Kuru
		response := api.PaySalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.Payments, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Vergi numarası ile E-Fatura mükellef bilgileri görüntüleme
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.EInvoiceInboxes.Data.Type = "e_invoice_inboxes" // << Değişiklik yapmayınız !
		request.EInvoiceInboxes.Data.Attributes.VKN = ""        // Vergi numarası
		response := api.ListEInvoiceInboxes(request)
		pretty, _ := json.MarshalIndent(response.EInvoiceInboxes, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Satış faturasını resmileştirme
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.EInvoiceInboxes.Data.Type = "e_invoice_inboxes" // << Değişiklik yapmayınız !
		request.EInvoiceInboxes.Data.Attributes.VKN = ""        // Vergi numarası sorgulama
		response := api.ListEInvoiceInboxes(request)
		if len(response.EInvoiceInboxes.Data) > 0 { // e-Fatura ise
			for _, data := range response.EInvoiceInboxes.Data {
				request := new(parasut.Request)
				request.EInvoices.Data.Type = "e_invoices" // << Değişiklik yapmayınız !
				request.EInvoices.Data.Relationships.Invoice = new(parasut.SingleRelationShip)
				request.EInvoices.Data.Relationships.Invoice.Data = new(parasut.RelationShip)
				request.EInvoices.Data.Relationships.Invoice.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
				request.EInvoices.Data.Relationships.Invoice.Data.ID = ""                 // Satış faturası ID

				request.EInvoices.Data.Attributes.To = data.Attributes.EInvoiceAddress
				request.EInvoices.Data.Attributes.Scenario = ""               // "basic" (Temel e-Fatura) || "commercial" (Ticari e-Fatura)
				request.EInvoices.Data.Attributes.Note = ""                   // Fatura notu
				request.EInvoices.Data.Attributes.VatExemptionReasonCode = "" // Firma KDV den muaf ise muafiyet sebebi kodu (Varsa)
				request.EInvoices.Data.Attributes.VatExemptionReason = ""     // Firma KDV den muaf ise muafiyet sebebi açıklaması (Varsa)
				request.EInvoices.Data.Attributes.VatWithholdingCode = ""     // Tevkifat oranına ait vergi kodu (Varsa)

				// Internet satışı (Varsa)
				request.EInvoices.Data.Attributes.InternetSale.URL = ""             // İnternet satışının yapıldığı url
				request.EInvoices.Data.Attributes.InternetSale.PaymentType = ""     // "KREDIKARTI/BANKAKARTI" "EFT/HAVALE" "KAPIDAODEME" "ODEMEARACISI" (Ödeme yöntemi)
				request.EInvoices.Data.Attributes.InternetSale.PaymentPlatform = "" // Ödeme platformu (iyzico,payu,banka adı vb.)
				request.EInvoices.Data.Attributes.InternetSale.PaymentDate = ""     // Ödeme tarihi (Yıl-Ay-Gün)

				response := api.CreateEInvoice(request)
				pretty, _ := json.MarshalIndent(response.EInvoices, " ", "\t")
				fmt.Println(string(pretty))
			}
		} else { // e-Arşiv ise
			request := new(parasut.Request)
			request.EArchives.Data.Type = "e_archives" // << Değişiklik yapmayınız !
			request.EArchives.Data.Relationships.SalesInvoice = new(parasut.SingleRelationShip)
			request.EArchives.Data.Relationships.SalesInvoice.Data = new(parasut.RelationShip)
			request.EArchives.Data.Relationships.SalesInvoice.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
			request.EArchives.Data.Relationships.SalesInvoice.Data.ID = ""                 // Satış faturası ID

			request.EArchives.Data.Attributes.Note = ""                   // Fatura notu
			request.EArchives.Data.Attributes.VatExemptionReasonCode = "" // Firma KDV den muaf ise muafiyet sebebi kodu (Varsa)
			request.EArchives.Data.Attributes.VatExemptionReason = ""     // Firma KDV den muaf ise muafiyet sebebi açıklaması (Varsa)
			request.EArchives.Data.Attributes.VatWithholdingCode = ""     // Tevkifat oranına ait vergi kodu (Varsa)

			// Internet satışı (Varsa)
			request.EArchives.Data.Attributes.InternetSale.URL = ""             // İnternet satışının yapıldığı url
			request.EArchives.Data.Attributes.InternetSale.PaymentType = ""     // "KREDIKARTI/BANKAKARTI" "EFT/HAVALE" "KAPIDAODEME" "ODEMEARACISI" (Ödeme yöntemi)
			request.EArchives.Data.Attributes.InternetSale.PaymentPlatform = "" // Ödeme platformu (iyzico,payu,banka adı vb.)
			request.EArchives.Data.Attributes.InternetSale.PaymentDate = ""     // Ödeme tarihi (Yıl-Ay-Gün)

			response := api.CreateEArchive(request)
			pretty, _ := json.MarshalIndent(response.EArchives, " ", "\t")
			fmt.Println(string(pretty))
		}
	}
}
```

# Resmileştirilmiş fatura kaydını görüntüleme
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.SalesInvoices.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.ID = ""                 // Satış faturası ID
		response := api.ShowSalesInvoice(request)
		docid := response.SalesInvoices.Data.Relationships.ActiveEDocument.Data.ID
		doctype := response.SalesInvoices.Data.Relationships.ActiveEDocument.Data.Type
		if doctype == "e_invoices" { // e-Fatura ise
			request := new(parasut.Request)
			request.EInvoices.Data.Type = doctype
			request.EInvoices.Data.ID = docid
			response := api.ShowEInvoice(request)
			pretty, _ := json.MarshalIndent(response.EInvoices, " ", "\t")
			fmt.Println(string(pretty))
		} else if doctype == "e_archives" { // e-Arşiv ise
			request := new(parasut.Request)
			request.EArchives.Data.Type = doctype
			request.EArchives.Data.ID = docid
			response := api.ShowEArchive(request)
			pretty, _ := json.MarshalIndent(response.EArchives, " ", "\t")
			fmt.Println(string(pretty))
		}
	}
}
```

# Resmileştirilmiş faturaya ait PDF url adresini görüntüleme
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.SalesInvoices.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.ID = ""                 // Satış faturası ID
		response := api.ShowSalesInvoice(request)
		docid := response.SalesInvoices.Data.Relationships.ActiveEDocument.Data.ID
		doctype := response.SalesInvoices.Data.Relationships.ActiveEDocument.Data.Type
		if doctype == "e_invoices" { // e-Fatura ise
			request := new(parasut.Request)
			request.EInvoices.Data.Type = "e_document_pdfs" // << Değişiklik yapmayınız !
			request.EInvoicePDF.Data.ID = docid
			response := api.ShowEInvoicePDF(request)
			pdfurl := response.EInvoicePDF.Data.Attributes.URL
			fmt.Println(pdfurl)
		} else if doctype == "e_archives" { // e-Arşiv ise
			request := new(parasut.Request)
			request.EArchives.Data.Type = "e_document_pdfs" // << Değişiklik yapmayınız !
			request.EArchivePDF.Data.ID = docid
			response := api.ShowEArchivePDF(request)
			pdfurl := response.EArchivePDF.Data.Attributes.URL
			fmt.Println(pdfurl)
		}
	}
}
```

# Müşteri/Tedarikçi kaydı oluşturma
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.Contacts.Data.Type = "contacts"           // << Değişiklik yapmayınız !
		request.Contacts.Data.Attributes.AccountType = "" // "customer" (Müşteri) || "supplier" (Tedarikçi)
		request.Contacts.Data.Attributes.Name = ""        // Firma Ünvanı
		request.Contacts.Data.Attributes.ShortName = ""   // Kısa İsim
		request.Contacts.Data.Attributes.ContactType = "" // "company" (Şirket) || "person" (Şahıs)
		request.Contacts.Data.Attributes.TaxNumber = ""   // Vergi Numarası
		request.Contacts.Data.Attributes.TaxOffice = ""   // Vergi Dairesi
		request.Contacts.Data.Attributes.City = ""        // İl
		request.Contacts.Data.Attributes.District = ""    // İlçe
		request.Contacts.Data.Attributes.Address = ""     // Adres
		request.Contacts.Data.Attributes.Phone = ""       // Telefon
		request.Contacts.Data.Attributes.Fax = ""         // Faks
		request.Contacts.Data.Attributes.Email = ""       // E-posta adresi
		request.Contacts.Data.Attributes.IBAN = ""        // IBAN numarası

		request.Contacts.Data.Relationships.Category = new(parasut.SingleRelationShip)
		request.Contacts.Data.Relationships.Category.Data = new(parasut.RelationShip)
		request.Contacts.Data.Relationships.Category.Data.Type = "item_categories" // << Değişiklik yapmayınız !
		request.Contacts.Data.Relationships.Category.Data.ID = ""                  // Kategori ID (varsa)

		response := api.CreateContact(request)
		pretty, _ := json.MarshalIndent(response.Contacts, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Müşteri/Tedarikçi kaydını silme
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.Contacts.Data.Type = "contacts" // << Değişiklik yapmayınız !
		request.Contacts.Data.ID = ""           // Müşteri/Tedarikçi ID
		response := api.DeleteContact(request)
		pretty, _ := json.MarshalIndent(response.Contacts, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Müşteri/Tedarikçi kaydını arşivleme
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.Contacts.Data.Type = "contacts" // << Değişiklik yapmayınız !
		request.Contacts.Data.ID = ""           // Müşteri/Tedarikçi ID
		response := api.ArchiveContact(request)
		pretty, _ := json.MarshalIndent(response.Contacts, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Müşteri/Tedarikçi kaydını arşivden çıkarma
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.Contacts.Data.Type = "contacts" // << Değişiklik yapmayınız !
		request.Contacts.Data.ID = ""           // Müşteri/Tedarikçi ID
		response := api.UnarchiveContact(request)
		pretty, _ := json.MarshalIndent(response.Contacts, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Müşteri/Tedarikçi kaydını görüntüleme
```go
package main

import (
	"encoding/json"
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.Contacts.Data.Type = "contacts" // << Değişiklik yapmayınız !
		request.Contacts.Data.ID = ""           // Müşteri/Tedarikçi ID
		response := api.ShowContact(request)
		pretty, _ := json.MarshalIndent(response.Contacts, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```
