[![license](https://img.shields.io/:license-mit-blue.svg)](https://github.com/ozgur-soft/parasut.go/blob/master/LICENSE.md)
[![documentation](https://pkg.go.dev/badge/github.com/ozgur-soft/parasut.go)](https://pkg.go.dev/github.com/ozgur-soft/parasut.go/src)

# Parasut.go
An easy-to-use parasut.com API (v4) with golang

# Installation
```bash
go get github.com/ozgur-soft/parasut.go
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
		request.Contact.Data.Type = "contacts"           // << Değişiklik yapmayınız !
		request.Contact.Data.Attributes.AccountType = "" // "customer" (Müşteri) || "supplier" (Tedarikçi)
		request.Contact.Data.Attributes.ContactType = "" // "company" (Şirket) || "person" (Şahıs)
		request.Contact.Data.Attributes.Name = ""        // Firma Ünvanı
		request.Contact.Data.Attributes.ShortName = ""   // Kısa İsim
		request.Contact.Data.Attributes.TaxNumber = ""   // Vergi Numarası
		request.Contact.Data.Attributes.TaxOffice = ""   // Vergi Dairesi
		request.Contact.Data.Attributes.Country = ""     // Ülke
		request.Contact.Data.Attributes.City = ""        // İl
		request.Contact.Data.Attributes.District = ""    // İlçe
		request.Contact.Data.Attributes.Address = ""     // Adres
		request.Contact.Data.Attributes.Phone = ""       // Telefon
		request.Contact.Data.Attributes.Fax = ""         // Faks
		request.Contact.Data.Attributes.Email = ""       // E-posta adresi
		request.Contact.Data.Attributes.IBAN = ""        // IBAN numarası
		response := api.CreateContact(request)
		pretty, _ := json.MarshalIndent(response.Contact, " ", "\t")
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
		request.Contact.Data.Type = "contacts" // << Değişiklik yapmayınız !
		request.Contact.Data.ID = ""           // Müşteri/Tedarikçi ID
		response := api.DeleteContact(request)
		pretty, _ := json.MarshalIndent(response.Contact, " ", "\t")
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
		request.Contact.Data.Type = "contacts" // << Değişiklik yapmayınız !
		request.Contact.Data.ID = ""           // Müşteri/Tedarikçi ID
		response := api.ArchiveContact(request)
		pretty, _ := json.MarshalIndent(response.Contact, " ", "\t")
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
		request.Contact.Data.Type = "contacts" // << Değişiklik yapmayınız !
		request.Contact.Data.ID = ""           // Müşteri/Tedarikçi ID
		response := api.UnarchiveContact(request)
		pretty, _ := json.MarshalIndent(response.Contact, " ", "\t")
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
		request.Contact.Data.Type = "contacts" // << Değişiklik yapmayınız !
		request.Contact.Data.ID = ""           // Müşteri/Tedarikçi ID
		response := api.ShowContact(request)
		pretty, _ := json.MarshalIndent(response.Contact, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Peşin satış faturası oluşturma
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
		request.SalesInvoice.Data.Type = "sales_invoices"             // << Değişiklik yapmayınız !
		request.SalesInvoice.Data.Attributes.ItemType = "invoice"     // << Değişiklik yapmayınız !
		request.SalesInvoice.Data.Attributes.Description = ""         // Fatura başlığı
		request.SalesInvoice.Data.Attributes.IssueDate = ""           // Fatura tarihi (Yıl-Ay-Gün)
		request.SalesInvoice.Data.Attributes.ShipmentIncluded = false // İrsaliyeli fatura
		request.SalesInvoice.Data.Attributes.CashSale = true          // Peşin satış
		request.SalesInvoice.Data.Attributes.PaymentDate = ""         // Peşin ödeme tarihi (Yıl-Ay-Gün)
		request.SalesInvoice.Data.Attributes.PaymentDescription = ""  // Peşin ödeme açıklaması
		request.SalesInvoice.Data.Attributes.PaymentAccountID = ""    // Paraşüt Banka ID (zorunlu)
		request.SalesInvoice.Data.Attributes.Currency = "TRL"         // Para birimi : "TRL", "USD", "EUR", "GBP"

		request.SalesInvoice.Data.Relationships.Contact = new(parasut.SingleRelationShip)
		request.SalesInvoice.Data.Relationships.Contact.Data = new(parasut.RelationShip)
		request.SalesInvoice.Data.Relationships.Contact.Data.Type = "contacts" // << Değişiklik yapmayınız !
		request.SalesInvoice.Data.Relationships.Contact.Data.ID = ""           // Müşteri ID

		detail := request.SalesInvoice.Data.Relationships.Details.Detail
		detail.Type = "sales_invoice_details" // << Değişiklik yapmayınız !
		detail.Attributes.Quantity = "1"      // Ürün miktarı
		detail.Attributes.UnitPrice = "1.00"  // Ürün birim fiyatı
		detail.Attributes.VatRate = "18"      // Ürün KDV oranı
		detail.Relationships.Product = new(parasut.SingleRelationShip)
		detail.Relationships.Product.Data = new(parasut.RelationShip)
		detail.Relationships.Product.Data.Type = "products" // << Değişiklik yapmayınız !
		detail.Relationships.Product.Data.ID = ""           // Paraşüt Ürün ID (zorunlu)
		request.SalesInvoice.Data.Relationships.Details.Data = append(request.SalesInvoice.Data.Relationships.Details.Data, detail)

		response := api.CreateSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoice, " ", "\t")
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
		request.SalesInvoice.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoice.Data.ID = ""                 // Satış faturası ID
		response := api.DeleteSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoice, " ", "\t")
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
		request.SalesInvoice.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoice.Data.ID = ""                 // Satış faturası ID
		response := api.CancelSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoice, " ", "\t")
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
		request.SalesInvoice.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoice.Data.ID = ""                 // Satış faturası ID
		response := api.ArchiveSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoice, " ", "\t")
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
		request.SalesInvoice.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoice.Data.ID = ""                 // Satış faturası ID
		response := api.UnarchiveSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoice, " ", "\t")
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
		request.SalesInvoice.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoice.Data.ID = ""                 // Satış faturası ID
		response := api.ShowSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoice, " ", "\t")
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
				request.EInvoice.Data.Type = "e_invoices" // << Değişiklik yapmayınız !
				request.EInvoice.Data.Relationships.Invoice = new(parasut.SingleRelationShip)
				request.EInvoice.Data.Relationships.Invoice.Data = new(parasut.RelationShip)
				request.EInvoice.Data.Relationships.Invoice.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
				request.EInvoice.Data.Relationships.Invoice.Data.ID = ""                 // Satış faturası ID
				request.EInvoice.Data.Attributes.To = data.Attributes.EInvoiceAddress
				request.EInvoice.Data.Attributes.Scenario = "" // "basic" (Temel e-Fatura) || "commercial" (Ticari e-Fatura)
				request.EInvoice.Data.Attributes.Note = ""     // Fatura notu
				response := api.CreateEInvoice(request)
				pretty, _ := json.MarshalIndent(response.EInvoice, " ", "\t")
				fmt.Println(string(pretty))
			}
		} else { // e-Arşiv ise
			request := new(parasut.Request)
			request.EArchive.Data.Type = "e_archives" // << Değişiklik yapmayınız !
			request.EArchive.Data.Relationships.SalesInvoice = new(parasut.SingleRelationShip)
			request.EArchive.Data.Relationships.SalesInvoice.Data = new(parasut.RelationShip)
			request.EArchive.Data.Relationships.SalesInvoice.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
			request.EArchive.Data.Relationships.SalesInvoice.Data.ID = ""                 // Satış faturası ID
			request.EArchive.Data.Attributes.Note = ""                                    // Fatura notu
			// Internet satışı (Varsa)
			request.EArchive.Data.Attributes.InternetSale.URL = ""             // İnternet satışının yapıldığı url
			request.EArchive.Data.Attributes.InternetSale.PaymentType = ""     // "KREDIKARTI/BANKAKARTI" "EFT/HAVALE" "KAPIDAODEME" "ODEMEARACISI" (Ödeme yöntemi)
			request.EArchive.Data.Attributes.InternetSale.PaymentPlatform = "" // Ödeme platformu (iyzico,payu,banka adı vb.)
			request.EArchive.Data.Attributes.InternetSale.PaymentDate = ""     // Ödeme tarihi (Yıl-Ay-Gün)
			response := api.CreateEArchive(request)
			pretty, _ := json.MarshalIndent(response.EArchive, " ", "\t")
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
		request.SalesInvoice.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoice.Data.ID = ""                 // Satış faturası ID
		response := api.ShowSalesInvoice(request)
		docid := response.SalesInvoice.Data.Relationships.ActiveEDocument.Data.ID
		doctype := response.SalesInvoice.Data.Relationships.ActiveEDocument.Data.Type
		if doctype == "e_invoices" { // e-Fatura ise
			request := new(parasut.Request)
			request.EInvoice.Data.Type = doctype
			request.EInvoice.Data.ID = docid
			response := api.ShowEInvoice(request)
			pretty, _ := json.MarshalIndent(response.EInvoice, " ", "\t")
			fmt.Println(string(pretty))
		} else if doctype == "e_archives" { // e-Arşiv ise
			request := new(parasut.Request)
			request.EArchive.Data.Type = doctype
			request.EArchive.Data.ID = docid
			response := api.ShowEArchive(request)
			pretty, _ := json.MarshalIndent(response.EArchive, " ", "\t")
			fmt.Println(string(pretty))
		}
	}
}
```

# Resmileştirilmiş faturaya ait PDF url adresini görüntüleme
```go
package main

import (
	"fmt"

	parasut "github.com/ozgur-soft/parasut.go/src"
)

func main() {
	config := parasut.Config{CompanyID: "", ClientID: "", ClientSecret: "", Username: "", Password: ""}
	api := &parasut.API{Config: config}
	auth := api.Authorize()
	if auth {
		request := new(parasut.Request)
		request.SalesInvoice.Data.Type = "sales_invoices" // << Değişiklik yapmayınız !
		request.SalesInvoice.Data.ID = ""                 // Satış faturası ID
		response := api.ShowSalesInvoice(request)
		docid := response.SalesInvoice.Data.Relationships.ActiveEDocument.Data.ID
		doctype := response.SalesInvoice.Data.Relationships.ActiveEDocument.Data.Type
		if doctype == "e_invoices" { // e-Fatura ise
			request := new(parasut.Request)
			request.EInvoice.Data.Type = "e_document_pdfs" // << Değişiklik yapmayınız !
			request.EInvoicePDF.Data.ID = docid
			response := api.ShowEInvoicePDF(request)
			pdfurl := response.EInvoicePDF.Data.Attributes.URL
			fmt.Println(pdfurl)
		} else if doctype == "e_archives" { // e-Arşiv ise
			request := new(parasut.Request)
			request.EArchive.Data.Type = "e_document_pdfs" // << Değişiklik yapmayınız !
			request.EArchivePDF.Data.ID = docid
			response := api.ShowEArchivePDF(request)
			pdfurl := response.EArchivePDF.Data.Attributes.URL
			fmt.Println(pdfurl)
		}
	}
}
```