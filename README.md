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

# Müşteri kaydı oluşturmak için
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
		request.Contacts.Data.Type = "contacts"             // << Burada değişiklik yapmayınız !
		request.Contacts.Data.Attr.AccountType = "customer" // << Burada değişiklik yapmayınız !
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
		request.Contacts.Data.Attr.Email = ""               // E-posta adresi
		request.Contacts.Data.Attr.IBAN = ""                // IBAN numarası
		response := api.CreateContact(request)
		pretty, _ := json.MarshalIndent(response.Contacts, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Çalışan kaydı oluşturmak için
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
		request.Employees.Data.Type = "employees" // << Burada değişiklik yapmayınız !
		request.Employees.Data.Attr.Name = ""     // İsim
		request.Employees.Data.Attr.Email = ""    // E-posta adresi
		request.Employees.Data.Attr.TCKN = ""     // TC Kimlik Numarası
		request.Employees.Data.Attr.IBAN = ""     // IBAN numarası
		response := api.CreateEmployee(request)
		pretty, _ := json.MarshalIndent(response.Employees, " ", "\t")
		fmt.Println(string(pretty))
	}
}
```

# Satış faturası bilgilerini görüntülemek için
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

# Resmileştirilmiş fatura bilgilerini görüntülemek için
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
		docid := response.SalesInvoices.Data.RelationShips.ActiveEDocument.Data.ID
		doctype := response.SalesInvoices.Data.RelationShips.ActiveEDocument.Data.Type
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

# Faturaya ait PDF url adresini görüntülemek için
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
		docid := response.SalesInvoices.Data.RelationShips.ActiveEDocument.Data.ID
		doctype := response.SalesInvoices.Data.RelationShips.ActiveEDocument.Data.Type
		if doctype == "e_archives" { // Fatura tipi e-Arşiv ise
			request := parasut.Request{}
			request.EArchivePDF.Data.ID = docid
			response := api.ShowEArchivePDF(request)
			pdfurl := response.EArchivePDF.Data.Attr.URL
			fmt.Println(pdfurl)
		}
		if doctype == "e_invoices" { // Fatura tipi e-Fatura ise
			request := parasut.Request{}
			request.EInvoicePDF.Data.ID = docid
			response := api.ShowEInvoicePDF(request)
			pdfurl := response.EInvoicePDF.Data.Attr.URL
			fmt.Println(pdfurl)
		}
	}
}
```
