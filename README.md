# parasut
An easy-to-use parasut.com API (v4) with golang

# Müşteri/Tedarikçi kaydı oluşturma
```go
package main

import (
	"encoding/json"
	"fmt"
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Firma numarası
	config.ClientID = ""     // Müşteri numarası
	config.ClientSecret = "" // Müşteri anahtarı
	config.Username = ""     // Kullanıcı adı
	config.Password = ""     // Şifre
}

func main() {
	api := parasut.API{}
	auth := api.Authorize()
	if auth {
		request := parasut.Request{}
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

		request.Contacts.Data.Relationships.Category.Data.Type = "item_categories"  // << Değişiklik yapmayınız !
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
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Firma numarası
	config.ClientID = ""     // Müşteri numarası
	config.ClientSecret = "" // Müşteri anahtarı
	config.Username = ""     // Kullanıcı adı
	config.Password = ""     // Şifre
}

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
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Firma numarası
	config.ClientID = ""     // Müşteri numarası
	config.ClientSecret = "" // Müşteri anahtarı
	config.Username = ""     // Kullanıcı adı
	config.Password = ""     // Şifre
}

func main() {
	api := parasut.API{}
	auth := api.Authorize()
	if auth {
		request := parasut.Request{}
		request.Employees.Data.Type = "employees"    // << Değişiklik yapmayınız !
		request.Employees.Data.Attributes.Name = ""  // İsim
		request.Employees.Data.Attributes.Email = "" // E-posta adresi
		request.Employees.Data.Attributes.TCKN = ""  // TC Kimlik Numarası
		request.Employees.Data.Attributes.IBAN = ""  // IBAN numarası

		request.Employees.Data.Relationships.Category.Data.Type = "item_categories"  // << Değişiklik yapmayınız !
		request.Employees.Data.Relationships.Category.Data.ID = ""                   // Kategori ID (varsa)

		response := api.CreateEmployee(request)
		pretty, _ := json.MarshalIndent(response.Employees, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Çalışan bilgilerini görüntüleme
```go
package main

import (
	"encoding/json"
	"fmt"
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Firma numarası
	config.ClientID = ""     // Müşteri numarası
	config.ClientSecret = "" // Müşteri anahtarı
	config.Username = ""     // Kullanıcı adı
	config.Password = ""     // Şifre
}

func main() {
	api := parasut.API{}
	auth := api.Authorize()
	if auth {
		request := parasut.Request{}
		request.Employees.Data.ID = "" // Çalışan ID
		response := api.ShowEmployee(request)
		pretty, _ := json.MarshalIndent(response.Employees, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Satış faturası kaydı oluşturma
```go
package main

import (
	"encoding/json"
	"fmt"
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Firma numarası
	config.ClientID = ""     // Müşteri numarası
	config.ClientSecret = "" // Müşteri anahtarı
	config.Username = ""     // Kullanıcı adı
	config.Password = ""     // Şifre
}

func main() {
	api := parasut.API{}
	auth := api.Authorize()
	if auth {
		request := parasut.Request{}
		request.SalesInvoices.Data.Type = "sales_invoices"              // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.Attributes.ItemType = "invoice"      // << Değişiklik yapmayınız !
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

		request.SalesInvoices.Data.Relationships.Contact.Data.Type = "contacts" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.Relationships.Contact.Data.ID = ""           // Müşteri ID

		request.SalesInvoices.Data.Relationships.Category.Data.Type = "item_categories" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.Relationships.Category.Data.ID = ""                  // Kategori ID (varsa)

		request.SalesInvoices.Data.Relationships.Details.Fill.Relationships.Product.Data.Type = "products" // << Değişiklik yapmayınız !
		request.SalesInvoices.Data.Relationships.Details.Fill.Relationships.Product.Data.ID = ""           // Ürün ID

		request.SalesInvoices.Data.Relationships.Details.Fill.Type = "sales_invoice_details"     // << Değişiklik yapmayınız !
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
```

# Satış faturası bilgilerini görüntüleme
```go
package main

import (
	"encoding/json"
	"fmt"
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Firma numarası
	config.ClientID = ""     // Müşteri numarası
	config.ClientSecret = "" // Müşteri anahtarı
	config.Username = ""     // Kullanıcı adı
	config.Password = ""     // Şifre
}

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

# Satış faturasına ödeme kaydı oluşturma
```go
package main

import (
	"encoding/json"
	"fmt"
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Firma numarası
	config.ClientID = ""     // Müşteri numarası
	config.ClientSecret = "" // Müşteri anahtarı
	config.Username = ""     // Kullanıcı adı
	config.Password = ""     // Şifre
}

func main() {
	api := parasut.API{}
	auth := api.Authorize()
	if auth {
		request := parasut.Request{}
		request.SalesInvoices.Data.ID = ""                    // Satış Faturası ID
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
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Firma numarası
	config.ClientID = ""     // Müşteri numarası
	config.ClientSecret = "" // Müşteri anahtarı
	config.Username = ""     // Kullanıcı adı
	config.Password = ""     // Şifre
}

func main() {
	api := parasut.API{}
	auth := api.Authorize()
	if auth {
		request := parasut.Request{}
		request.EInvoiceInboxes.Data.Type = "e_invoice_inboxes" // << Değişiklik yapmayınız !
		request.EInvoiceInboxes.Data.Attributes.VKN = ""        // Vergi numarası
		response := api.ListEInvoiceInboxes(request)
		pretty, _ := json.MarshalIndent(response.EInvoiceInboxes, " ", "\t")
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
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Firma numarası
	config.ClientID = ""     // Müşteri numarası
	config.ClientSecret = "" // Müşteri anahtarı
	config.Username = ""     // Kullanıcı adı
	config.Password = ""     // Şifre
}

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
			request.EArchives.Data.Type = "e_archives" // << Değişiklik yapmayınız !
			request.EArchives.Data.ID = docid
			response := api.ShowEArchive(request)
			pretty, _ := json.MarshalIndent(response.EArchives, " ", "\t")
			fmt.Println(string(pretty))
		}
		if doctype == "e_invoices" { // Fatura tipi e-Fatura ise
			request := parasut.Request{}
			request.EInvoices.Data.Type = "e_invoices" // << Değişiklik yapmayınız !
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
	"parasut/config"
	"parasut/src"
)

func init() {
	config.CompanyID = ""    // Firma numarası
	config.ClientID = ""     // Müşteri numarası
	config.ClientSecret = "" // Müşteri anahtarı
	config.Username = ""     // Kullanıcı adı
	config.Password = ""     // Şifre
}

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
			request.EArchives.Data.Type = "e_document_pdfs" // << Değişiklik yapmayınız !
			request.EArchivePDF.Data.ID = docid
			response := api.ShowEArchivePDF(request)
			pdfurl := response.EArchivePDF.Data.Attributes.URL
			fmt.Println(pdfurl)
		}
		if doctype == "e_invoices" { // Fatura tipi e-Fatura ise
			request := parasut.Request{}
			request.EInvoices.Data.Type = "e_document_pdfs" // << Değişiklik yapmayınız !
			request.EInvoicePDF.Data.ID = docid
			response := api.ShowEInvoicePDF(request)
			pdfurl := response.EInvoicePDF.Data.Attributes.URL
			fmt.Println(pdfurl)
		}
	}
}
```
