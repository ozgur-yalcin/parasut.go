# parasut
An easy-to-use parasut.com API (Paraşüt v4) with golang


# Bağlantı ve kimlik doğrulama bilgileri (/config/config.go)
```go
package config

const (
	APIURL       = "https://api.parasut.com/v4/"
	TokenURL     = "https://api.parasut.com/oauth/token"
	CompanyID    = "" // Paraşüt tarafından belirlenen firma numarasını yazınız
	ClientID     = "" // Paraşüt tarafından belirlenen müşteri numarasını yazınız
	ClientSecret = "" // Paraşüt tarafından belirlenen müşteri anahtarını yazınız
	Username     = "" // Paraşüte giriş yaparken kullandığınız kullanıcı adını yazınız
	Password     = "" // Paraşüte giriş yaparken kullandığınız şifreyi yazınız
	GrantType    = "password" // << Burada değişiklik yapmayınız !
	RedirectURI  = "urn:ietf:wg:oauth:2.0:oob" // << Burada değişiklik yapmayınız !
)
```

# Satış faturası kaydı oluşturma
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
		api := parasut.API{}
		auth := api.Authorize()
		if auth {
			request := parasut.Request{}
			request.SalesInvoices.Data.Type = "sales_invoices"              // << Burada değişiklik yapmayınız !
			request.SalesInvoices.Data.Attributes.ItemType = "invoice"      // << Burada değişiklik yapmayınız !
			request.SalesInvoices.Data.Attributes.Description = ""          // Fatura başlığı
			request.SalesInvoices.Data.Attributes.TaxNumber = "11111111111" // Vergi numarası
			request.SalesInvoices.Data.Attributes.TaxOffice = ""            // Vergi dairesi
			request.SalesInvoices.Data.Attributes.IssueDate = ""            // Fatura tarihi (Yıl-Ay-Gün)
			request.SalesInvoices.Data.Attributes.Currency = "TRL"          // "TRL" || "USD" || "EUR" || "GBP" (Para birimi)
			request.SalesInvoices.Data.Attributes.BillingPhone = ""         // Telefon numarası
			request.SalesInvoices.Data.Attributes.BillingFax = ""           // Fax numarası
			request.SalesInvoices.Data.Attributes.BillingAddress = ""       // Fatura adresi
			request.SalesInvoices.Data.Attributes.City = ""                 // İl
			request.SalesInvoices.Data.Attributes.District = ""             // İlçe

			request.SalesInvoices.Data.Relationships.Contact.Data.Type = "contacts" // << Burada değişiklik yapmayınız !
			request.SalesInvoices.Data.Relationships.Contact.Data.ID = ""           // Müşteri/Tedarikçi ID (varsa)

			request.SalesInvoices.Data.Relationships.Category.Data.Type = "item_categories" // << Burada değişiklik yapmayınız !
			request.SalesInvoices.Data.Relationships.Category.Data.ID = ""                  // Kategori ID (varsa)

			request.SalesInvoices.Data.Relationships.Details.Fill.Relationships.Product.Data.Type = "products" // << Burada değişiklik yapmayınız !
			request.SalesInvoices.Data.Relationships.Details.Fill.Relationships.Product.Data.ID = ""           // Ürün ID

			request.SalesInvoices.Data.Relationships.Details.Fill.Type = "sales_invoice_details"     // << Burada değişiklik yapmayınız !
			request.SalesInvoices.Data.Relationships.Details.Fill.Attributes.Quantity = "0"          // Ürün miktarı
			request.SalesInvoices.Data.Relationships.Details.Fill.Attributes.UnitPrice = "0"         // Ürün birim fiyatı
			request.SalesInvoices.Data.Relationships.Details.Fill.Attributes.VatRate = "0"           // Ürün KDV oranı
			request.SalesInvoices.Data.Relationships.Details.Fill.Attributes.DiscountType = "amount" // "amount" || "percentage" (İndirim türü)
			request.SalesInvoices.Data.Relationships.Details.Fill.Attributes.DiscountValue = "0"     // İndirim oranı
			request.SalesInvoices.Data.Relationships.Details.Data = append(request.SalesInvoices.Data.Relationships.Details.Data, request.SalesInvoices.Data.Relationships.Details.Fill)

			response := api.CreateSalesInvoice(request)
			pretty, _ := json.MarshalIndent(response.SalesInvoices, " ", "\t")
			fmt.Println(string(pretty))
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
	"parasut/src"
)

func main() {
	api := parasut.API{}
	auth := api.Authorize()
	if auth {
		request := parasut.Request{}
		request.Contacts.Data.Type = "contacts"           // << Burada değişiklik yapmayınız !
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

		request.Contacts.Data.Relationships.Category.Data.Type = "item_categories"  // << Burada değişiklik yapmayınız !
		request.Contacts.Data.Relationships.Category.Data.ID = ""                   // Kategori ID (varsa)

		response := api.CreateContact(request)
		pretty, _ := json.MarshalIndent(response.Contacts, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Müşteri/Tedarikçi bilgilerini görüntüleme
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
		request.Contacts.Data.ID = "" // Müşteri/Tedarikçi ID
		response := api.ShowContact(request)
		pretty, _ := json.MarshalIndent(response.Contacts, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Çalışan kaydı oluşturma
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
		request.Employees.Data.Type = "employees"    // << Burada değişiklik yapmayınız !
		request.Employees.Data.Attributes.Name = ""  // İsim
		request.Employees.Data.Attributes.Email = "" // E-posta adresi
		request.Employees.Data.Attributes.TCKN = ""  // TC Kimlik Numarası
		request.Employees.Data.Attributes.IBAN = ""  // IBAN numarası
		response := api.CreateEmployee(request)
		pretty, _ := json.MarshalIndent(response.Employees, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Satış faturası bilgilerini görüntüleme
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
		response := api.ShowSalesInvoice(request)
		pretty, _ := json.MarshalIndent(response.SalesInvoices, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Resmileştirilmiş fatura bilgilerini görüntüleme
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
		response := api.ShowSalesInvoice(request)
		docid := response.SalesInvoices.Data.Relationships.ActiveEDocument.Data.ID
		doctype := response.SalesInvoices.Data.Relationships.ActiveEDocument.Data.Type
		if doctype == "e_archives" { // Fatura tipi e-Arşiv ise
			request := parasut.Request{}
			request.EArchives.Data.ID = docid
			earchive := api.ShowEArchive(request)
			pretty, _ := json.MarshalIndent(earchive.EArchives, " ", "\t")
			fmt.Println(string(pretty))
		}
		if doctype == "e_invoices" { // Fatura tipi e-Fatura ise
			request := parasut.Request{}
			request.EInvoices.Data.ID = docid
			response := api.ShowEInvoice(request)
			pretty, _ := json.MarshalIndent(response.EInvoices, " ", "\t")
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
	"parasut/src"
)

func main() {
	api := parasut.API{}
	auth := api.Authorize()
	if auth {
		request := parasut.Request{}
		request.SalesInvoices.Data.ID = "" // Satış Faturası ID
		response := api.ShowSalesInvoice(request)
		docid := response.SalesInvoices.Data.Relationships.ActiveEDocument.Data.ID
		doctype := response.SalesInvoices.Data.Relationships.ActiveEDocument.Data.Type
		if doctype == "e_archives" { // Fatura tipi e-Arşiv ise
			request := parasut.Request{}
			request.EArchivePDF.Data.ID = docid
			response := api.ShowEArchivePDF(request)
			pdfurl := response.EArchivePDF.Data.Attributes.URL
			fmt.Println(pdfurl)
		}
		if doctype == "e_invoices" { // Fatura tipi e-Fatura ise
			request := parasut.Request{}
			request.EInvoicePDF.Data.ID = docid
			response := api.ShowEInvoicePDF(request)
			pdfurl := response.EInvoicePDF.Data.Attributes.URL
			fmt.Println(pdfurl)
		}
	}
}
```
