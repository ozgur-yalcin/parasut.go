package parasut

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Config struct {
	CompanyID    string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

type API struct {
	Config Config

	Authentication struct {
		AccessToken  string `json:"access_token,omitempty"`
		TokenType    string `json:"token_type,omitempty"`
		ExpiresIn    string `json:"expires_in,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
		Scope        string `json:"scope,omitempty"`
		CreatedAt    string `json:"created_at,omitempty"`
	}
}

type Request struct {
	Contact struct {
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Name        string `json:"name,omitempty"`
				ShortName   string `json:"short_name,omitempty"`
				Email       string `json:"email,omitempty"`
				AccountType string `json:"account_type,omitempty"`
				ContactType string `json:"contact_type,omitempty"`
				IBAN        string `json:"iban,omitempty"`
				TaxOffice   string `json:"tax_office,omitempty"`
				TaxNumber   string `json:"tax_number,omitempty"`
				Country     string `json:"country,omitempty"`
				City        string `json:"city,omitempty"`
				District    string `json:"district,omitempty"`
				Address     string `json:"address,omitempty"`
				Phone       string `json:"phone,omitempty"`
				Fax         string `json:"fax,omitempty"`
				IsAbroad    bool   `json:"is_abroad,omitempty"`
				Archived    bool   `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Category      *SingleRelationShip `json:"category,omitempty"`
				ContactPortal *SingleRelationShip `json:"contact_portal,omitempty"`
				ContactPeople *MultiRelationShip  `json:"contact_people,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	Employee struct {
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Name     string `json:"name,omitempty"`
				TCKN     string `json:"tckn,omitempty"`
				Email    string `json:"email,omitempty"`
				IBAN     string `json:"iban,omitempty"`
				Archived bool   `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Category          *SingleRelationShip `json:"category,omitempty"`
				ManagedByUser     *SingleRelationShip `json:"managed_by_user,omitempty"`
				ManagedByUserRole *MultiRelationShip  `json:"managed_by_user_role,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	SalesInvoice struct {
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				InvoiceSeries       string      `json:"invoice_series,omitempty"`
				InvoiceID           json.Number `json:"invoice_id,omitempty"`
				ExchangeRate        json.Number `json:"exchange_rate,omitempty"`
				WithholdingRate     json.Number `json:"withholding_rate,omitempty"`
				VatWithholdingRate  json.Number `json:"vat_withholding_rate,omitempty"`
				TotalDiscount       json.Number `json:"total_discount,omitempty"`
				InvoiceDiscount     json.Number `json:"invoice_discount,omitempty"`
				InvoiceDiscountType string      `json:"invoice_discount_type,omitempty"`
				Currency            string      `json:"currency,omitempty"`
				ItemType            string      `json:"item_type,omitempty"`
				Description         string      `json:"description,omitempty"`
				IssueDate           string      `json:"issue_date,omitempty"`
				DueDate             string      `json:"due_date,omitempty"`
				BillingAddress      string      `json:"billing_address,omitempty"`
				BillingPhone        string      `json:"billing_phone,omitempty"`
				BillingFax          string      `json:"billing_fax,omitempty"`
				TaxOffice           string      `json:"tax_office,omitempty"`
				TaxNumber           string      `json:"tax_number,omitempty"`
				PayerTaxNumbers     []string    `json:"payer_tax_numbers,omitempty"`
				Country             string      `json:"country,omitempty"`
				City                string      `json:"city,omitempty"`
				District            string      `json:"district,omitempty"`
				OrderNo             string      `json:"order_no,omitempty"`
				OrderDate           string      `json:"order_date,omitempty"`
				ShipmentAddress     string      `json:"shipment_addres,omitempty"`
				ShipmentIncluded    bool        `json:"shipment_included,omitempty"`
				PaymentAccountID    string      `json:"payment_account_id,omitempty"`
				PaymentDate         string      `json:"payment_date,omitempty"`
				PaymentDescription  string      `json:"payment_description,omitempty"`
				CashSale            bool        `json:"cash_sale,omitempty"`
				IsAbroad            bool        `json:"is_abroad,omitempty"`
				Archived            bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Details struct {
					Detail struct {
						Type       string `json:"type,omitempty"`
						ID         string `json:"id,omitempty"`
						Attributes struct {
							Quantity              json.Number `json:"quantity,omitempty"`
							UnitPrice             json.Number `json:"unit_price,omitempty"`
							VatRate               json.Number `json:"vat_rate,omitempty"`
							DiscountValue         json.Number `json:"discount_value,omitempty"`
							ExciseDutyValue       json.Number `json:"excise_duty_value,omitempty"`
							CommunicationsTaxRate json.Number `json:"communications_tax_rate,omitempty"`
							ProductID             string      `json:"product_id,omitempty"`
							Description           string      `json:"description,omitempty"`
							DiscountType          string      `json:"discount_type,omitempty"`
							ExciseDutyType        string      `json:"excise_duty_type,omitempty"`
						} `json:"attributes,omitempty"`
						Relationships struct {
							Product *SingleRelationShip `json:"product,omitempty"`
						} `json:"relationships,omitempty"`
					} `json:"-"`
					Data []interface{} `json:"data,omitempty"`
				} `json:"details,omitempty"`
				Contact  *SingleRelationShip `json:"contact,omitempty"`
				Category *SingleRelationShip `json:"category,omitempty"`
				Tags     *MultiRelationShip  `json:"tags,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	EArchive struct {
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				VatWithholdingCode     string `json:"vat_withholding_code,omitempty"`
				VatExemptionReasonCode string `json:"vat_exemption_reason_code,omitempty"`
				VatExemptionReason     string `json:"vat_exemption_reason,omitempty"`
				Note                   string `json:"note,omitempty"`
				ExciseDutyCodes        struct {
					Product             string `json:"product,omitempty"`
					SalesExciseDutyCode string `json:"sales_excise_duty_code,omitempty"`
				} `json:"excise_duty_codes,omitempty"`
				InternetSale struct {
					URL             string `json:"url,omitempty"`
					PaymentType     string `json:"payment_type,omitempty"`
					PaymentPlatform string `json:"payment_platform,omitempty"`
					PaymentDate     string `json:"payment_date,omitempty"`
				} `json:"internet_sale,omitempty"`
				Shipment struct {
					Title string `json:"title,omitempty"`
					Name  string `json:"name,omitempty"`
					VKN   string `json:"vkn,omitempty"`
					TCKN  string `json:"tckn,omitempty"`
					Date  string `json:"date,omitempty"`
				} `json:"shipment,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				SalesInvoice *SingleRelationShip `json:"sales_invoice,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	EInvoice struct {
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				To                     string `json:"to,omitempty"`
				Scenario               string `json:"scenario,omitempty"`
				VatWithholdingCode     string `json:"vat_withholding_code,omitempty"`
				VatExemptionReasonCode string `json:"vat_exemption_reason_code,omitempty"`
				VatExemptionReason     string `json:"vat_exemption_reason,omitempty"`
				Note                   string `json:"note,omitempty"`
				ExciseDutyCodes        struct {
					Product             string `json:"product,omitempty"`
					SalesExciseDutyCode string `json:"sales_excise_duty_code,omitempty"`
				} `json:"excise_duty_codes,omitempty"`
				InternetSale struct {
					URL             string `json:"url,omitempty"`
					PaymentType     string `json:"payment_type,omitempty"`
					PaymentPlatform string `json:"payment_platform,omitempty"`
					PaymentDate     string `json:"payment_date,omitempty"`
				} `json:"internet_sale,omitempty"`
				Shipment struct {
					Title string `json:"title,omitempty"`
					Name  string `json:"name,omitempty"`
					VKN   string `json:"vkn,omitempty"`
					TCKN  string `json:"tckn,omitempty"`
					Date  string `json:"date,omitempty"`
				} `json:"shipment,omitempty"`
				CustomRequirementParams struct {
					Products    []*EInvoiceProduct `json:"products,omitempty"`
					Integration struct {
						Data struct {
							AdditionalInvoiceType string `json:"additional_invoice_type,omitempty"`
							SupplierCode          string `json:"supplier_code,omitempty"`
							TaxPayerCode          string `json:"tax_payer_code,omitempty"`
							TaxPayerName          string `json:"tax_payer_name,omitempty"`
							FileNumber            string `json:"file_number,omitempty"`
							TermStartDate         string `json:"term_start_date,omitempty"`
							TermEndDate           string `json:"term_end_date,omitempty"`
						} `json:"data,omitempty"`
					} `json:"integration,omitempty"`
				} `json:"custom_requirement_params,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Invoice *SingleRelationShip `json:"invoice,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	EArchivePDF SingleRelationShip

	EInvoicePDF SingleRelationShip

	EInvoiceInboxes struct {
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				VKN string `json:"vkn,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	Transaction SingleRelationShip

	TrackableJob SingleRelationShip
}

type Response struct {
	Contact struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Balance     json.Number `json:"balance,omitempty"`
				TrlBalance  json.Number `json:"trl_balance,omitempty"`
				UsdBalance  json.Number `json:"usd_balance,omitempty"`
				EurBalance  json.Number `json:"eur_balance,omitempty"`
				GbpBalance  json.Number `json:"gbp_balance,omitempty"`
				CreatedAt   string      `json:"created_at,omitempty"`
				UpdatedAt   string      `json:"updated_at,omitempty"`
				Name        string      `json:"name,omitempty"`
				ShortName   string      `json:"short_name,omitempty"`
				Email       string      `json:"email,omitempty"`
				AccountType string      `json:"account_type,omitempty"`
				ContactType string      `json:"contact_type,omitempty"`
				IBAN        string      `json:"iban,omitempty"`
				TaxOffice   string      `json:"tax_office,omitempty"`
				TaxNumber   string      `json:"tax_number,omitempty"`
				Country     string      `json:"country,omitempty"`
				City        string      `json:"city,omitempty"`
				District    string      `json:"district,omitempty"`
				Address     string      `json:"address,omitempty"`
				Phone       string      `json:"phone,omitempty"`
				Fax         string      `json:"fax,omitempty"`
				IsAbroad    bool        `json:"is_abroad,omitempty"`
				Archived    bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Category      *SingleRelationShip `json:"category,omitempty"`
				ContactPortal *SingleRelationShip `json:"contact_portal,omitempty"`
				ContactPeople *MultiRelationShip  `json:"contact_people,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
		Included []struct {
			ID            string      `json:"id,omitempty"`
			Type          string      `json:"type,omitempty"`
			Attributes    interface{} `json:"attributes,omitempty"`
			Relationships interface{} `json:"relationships,omitempty"`
		} `json:"included,omitempty"`
	}

	Employee struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Balance    json.Number `json:"balance,omitempty"`
				TrlBalance json.Number `json:"trl_balance,omitempty"`
				UsdBalance json.Number `json:"usd_balance,omitempty"`
				EurBalance json.Number `json:"eur_balance,omitempty"`
				GbpBalance json.Number `json:"gbp_balance,omitempty"`
				CreatedAt  string      `json:"created_at,omitempty"`
				UpdatedAt  string      `json:"updated_at,omitempty"`
				Name       string      `json:"name,omitempty"`
				TCKN       string      `json:"tckn,omitempty"`
				Email      string      `json:"email,omitempty"`
				IBAN       string      `json:"iban,omitempty"`
				Archived   bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Category          *SingleRelationShip `json:"category,omitempty"`
				ManagedByUser     *SingleRelationShip `json:"managed_by_user,omitempty"`
				ManagedByUserRole *MultiRelationShip  `json:"managed_by_user_role,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
		Included []struct {
			ID            string      `json:"id,omitempty"`
			Type          string      `json:"type,omitempty"`
			Attributes    interface{} `json:"attributes,omitempty"`
			Relationships interface{} `json:"relationships,omitempty"`
		} `json:"included,omitempty"`
	}

	SalesInvoice struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				InvoiceSeries          string      `json:"invoice_series,omitempty"`
				InvoiceNo              string      `json:"invoice_no,omitempty"`
				InvoiceID              json.Number `json:"invoice_id,omitempty"`
				ExchangeRate           json.Number `json:"exchange_rate,omitempty"`
				WithholdingRate        json.Number `json:"withholding_rate,omitempty"`
				VatWithholdingRate     json.Number `json:"vat_withholding_rate,omitempty"`
				NetTotal               json.Number `json:"net_total,omitempty"`
				GrossTotal             json.Number `json:"gross_total,omitempty"`
				Withholding            json.Number `json:"withholding,omitempty"`
				TotalExciseDuty        json.Number `json:"total_excise_duty,omitempty"`
				TotalCommunicationsTax json.Number `json:"total_communications_tax,omitempty"`
				TotalVat               json.Number `json:"total_vat,omitempty"`
				VatWithholding         json.Number `json:"vat_withholding,omitempty"`
				BeforeTaxesTotal       json.Number `json:"before_taxes_total,omitempty"`
				Remaining              json.Number `json:"remaining,omitempty"`
				RemainingInTrl         json.Number `json:"remaining_in_trl,omitempty"`
				TotalDiscount          json.Number `json:"total_discount,omitempty"`
				TotalInvoiceDiscount   json.Number `json:"total_invoice_discount,omitempty"`
				InvoiceDiscount        json.Number `json:"invoice_discount,omitempty"`
				InvoiceDiscountType    string      `json:"invoice_discount_type,omitempty"`
				Currency               string      `json:"currency,omitempty"`
				PaymentStatus          string      `json:"payment_status,omitempty"`
				ItemType               string      `json:"item_type,omitempty"`
				Description            string      `json:"description,omitempty"`
				CreatedAt              string      `json:"created_at,omitempty"`
				UpdatedAt              string      `json:"updated_at,omitempty"`
				IssueDate              string      `json:"issue_date,omitempty"`
				DueDate                string      `json:"due_date,omitempty"`
				BillingAddress         string      `json:"billing_address,omitempty"`
				BillingPhone           string      `json:"billing_phone,omitempty"`
				BillingFax             string      `json:"billing_fax,omitempty"`
				TaxOffice              string      `json:"tax_office,omitempty"`
				TaxNumber              string      `json:"tax_number,omitempty"`
				PayerTaxNumbers        []string    `json:"payer_tax_numbers,omitempty"`
				Country                string      `json:"country,omitempty"`
				City                   string      `json:"city,omitempty"`
				District               string      `json:"district,omitempty"`
				OrderNo                string      `json:"order_no,omitempty"`
				OrderDate              string      `json:"order_date,omitempty"`
				ShipmentAddress        string      `json:"shipment_addres,omitempty"`
				ShipmentIncluded       bool        `json:"shipment_included,omitempty"`
				CashSale               bool        `json:"cash_sale,omitempty"`
				IsAbroad               bool        `json:"is_abroad,omitempty"`
				Archived               bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Category        *SingleRelationShip `json:"category,omitempty"`
				Contact         *SingleRelationShip `json:"contact,omitempty"`
				Details         *MultiRelationShip  `json:"details,omitempty"`
				Payments        *MultiRelationShip  `json:"payments,omitempty"`
				Tags            *MultiRelationShip  `json:"tags,omitempty"`
				SalesOffer      *SingleRelationShip `json:"sales_offer,omitempty"`
				Sharings        *MultiRelationShip  `json:"sharings,omitempty"`
				RecurrencePlan  *MultiRelationShip  `json:"recurrence_plan,omitempty"`
				ActiveEDocument *SingleRelationShip `json:"active_e_document,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
		Included []struct {
			ID            string      `json:"id,omitempty"`
			Type          string      `json:"type,omitempty"`
			Attributes    interface{} `json:"attributes,omitempty"`
			Relationships interface{} `json:"relationships,omitempty"`
		} `json:"included,omitempty"`
	}

	EArchive struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				CreatedAt        string `json:"created_at,omitempty"`
				UpdatedAt        string `json:"updated_at,omitempty"`
				PrintedAt        string `json:"printed_at,omitempty"`
				CancellableUntil string `json:"cancellable_until,omitempty"`
				UUID             string `json:"uuid,omitempty"`
				VKN              string `json:"vkn,omitempty"`
				InvoiceNumber    string `json:"invoice_number,omitempty"`
				Note             string `json:"note,omitempty"`
				Status           string `json:"status,omitempty"`
				IsPrinted        bool   `json:"is_printed,omitempty"`
				IsSigned         bool   `json:"is_signed,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				SalesInvoice *SingleRelationShip `json:"sales_invoice,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
		Included []struct {
			ID            string      `json:"id,omitempty"`
			Type          string      `json:"type,omitempty"`
			Attributes    interface{} `json:"attributes,omitempty"`
			Relationships interface{} `json:"relationships,omitempty"`
		} `json:"included,omitempty"`
	}

	EInvoice struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				ExternalID   string      `json:"external_id,omitempty"`
				UUID         string      `json:"uuid,omitempty"`
				EnvUUID      string      `json:"env_uuid,omitempty"`
				FromAddress  string      `json:"from_address,omitempty"`
				FromVKN      string      `json:"from_vkn,omitempty"`
				ToAddress    string      `json:"to_address,omitempty"`
				ToVKN        string      `json:"to_vkn,omitempty"`
				Direction    string      `json:"direction,omitempty"`
				Scenario     string      `json:"scenario,omitempty"`
				ResponseType string      `json:"response_type,omitempty"`
				ContactName  string      `json:"contact_name,omitempty"`
				NetTotal     json.Number `json:"net_total,omitempty"`
				Currency     string      `json:"currency,omitempty"`
				ItemType     string      `json:"item_type,omitempty"`
				Note         string      `json:"note,omitempty"`
				Status       string      `json:"status,omitempty"`
				CreatedAt    string      `json:"created_at,omitempty"`
				UpdatedAt    string      `json:"updated_at,omitempty"`
				IssueDate    string      `json:"issue_date,omitempty"`
				IsExpired    bool        `json:"is_expired,omitempty"`
				IsAnswerable bool        `json:"is_answerable,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Invoice *SingleRelationShip `json:"invoice,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
		Included []struct {
			ID            string      `json:"id,omitempty"`
			Type          string      `json:"type,omitempty"`
			Attributes    interface{} `json:"attributes,omitempty"`
			Relationships interface{} `json:"relationships,omitempty"`
		} `json:"included,omitempty"`
	}

	EArchivePDF struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				URL       string `json:"url,omitempty"`
				ExpiresAt string `json:"expires_at,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	EInvoicePDF struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				URL       string `json:"url,omitempty"`
				ExpiresAt string `json:"expires_at,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	EInvoiceInboxes struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data []struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				VKN                 string `json:"vkn,omitempty"`
				EInvoiceAddress     string `json:"e_invoice_address,omitempty"`
				Name                string `json:"name,omitempty"`
				InboxType           string `json:"inbox_type,omitempty"`
				AddressRegisteredAt string `json:"address_registered_at,omitempty"`
				RegisteredAt        string `json:"registered_at,omitempty"`
				CreatedAt           string `json:"created_at,omitempty"`
				UpdatedAt           string `json:"updated_at,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	Transaction struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Description    string      `json:"description,omitempty"`
				Type           string      `json:"transaction_type,omitempty"`
				Date           string      `json:"date,omitempty"`
				AmountInTrl    json.Number `json:"amount_in_trl,omitempty"`
				DebitAmount    json.Number `json:"debit_amount,omitempty"`
				DebitCurrency  string      `json:"debit_currency,omitempty"`
				CreditAmount   json.Number `json:"credit_amount,omitempty"`
				CreditCurrency string      `json:"credit_currency,omitempty"`
				CreatedAt      string      `json:"created_at,omitempty"`
				UpdatedAt      string      `json:"updated_at,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
		Included []struct {
			ID            string      `json:"id,omitempty"`
			Type          string      `json:"type,omitempty"`
			Attributes    interface{} `json:"attributes,omitempty"`
			Relationships interface{} `json:"relationships,omitempty"`
		} `json:"included,omitempty"`
	}

	TrackableJob struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Status string `json:"status,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}
}

type RelationShip struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type SingleRelationShip struct {
	Data *RelationShip `json:"data,omitempty"`
}

type MultiRelationShip struct {
	Data []*RelationShip `json:"data,omitempty"`
}

type EInvoiceProduct struct {
	ID   string `json:"product_id,omitempty"`
	Data struct {
		BuyerCode string `json:"buyer_code,omitempty"`
	} `json:"data,omitempty"`
}

func (api *API) Authorize() bool {
	apidata := url.Values{}
	apidata.Add("client_id", api.Config.ClientID)
	apidata.Add("client_secret", api.Config.ClientSecret)
	apidata.Add("username", api.Config.Username)
	apidata.Add("password", api.Config.Password)
	apidata.Add("grant_type", "password")
	apidata.Add("redirect_uri", "urn:ietf:wg:oauth:2.0:oob")
	client := new(http.Client)
	req, err := http.NewRequest("POST", "https://api.parasut.com/oauth/token", strings.NewReader(apidata.Encode()))
	if err != nil {
		log.Println(err)
		return false
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&api.Authentication)
	return true
}

func (api *API) CreateContact(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/contacts?include=category,contact_portal,contact_people"
	request.Contact.Data.Type = "contacts"
	contactdata, _ := json.Marshal(request.Contact)
	client := new(http.Client)
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(contactdata))
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Contact)
	return response
}

func (api *API) ShowContact(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/contacts/" + request.Contact.Data.ID + "?include=category,contact_portal,contact_people"
	client := new(http.Client)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Contact)
	return response
}

func (api *API) DeleteContact(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/contacts/" + request.Contact.Data.ID
	client := new(http.Client)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Contact)
	return response
}

func (api *API) ArchiveContact(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/contacts/" + request.Contact.Data.ID + "/archive"
	client := new(http.Client)
	req, err := http.NewRequest("PATCH", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Contact)
	return response
}

func (api *API) UnarchiveContact(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/contacts/" + request.Contact.Data.ID + "/unarchive"
	client := new(http.Client)
	req, err := http.NewRequest("PATCH", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Contact)
	return response
}

func (api *API) CreateEmployee(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/employees?include=category,managed_by_user,managed_by_user_role"
	request.Employee.Data.Type = "employees"
	employeedata, _ := json.Marshal(request.Employee)
	client := new(http.Client)
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(employeedata))
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Employee)
	return response
}

func (api *API) ShowEmployee(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/employees/" + request.Employee.Data.ID + "?include=category,managed_by_user,managed_by_user_role"
	client := new(http.Client)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Employee)
	return response
}

func (api *API) DeleteEmployee(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/employees/" + request.Employee.Data.ID
	client := new(http.Client)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Employee)
	return response
}

func (api *API) ArchiveEmployee(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/employees/" + request.Employee.Data.ID + "/archive"
	client := new(http.Client)
	req, err := http.NewRequest("PATCH", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Employee)
	return response
}

func (api *API) UnarchiveEmployee(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/employees/" + request.Employee.Data.ID + "/unarchive"
	client := new(http.Client)
	req, err := http.NewRequest("PATCH", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Employee)
	return response
}

func (api *API) CreateSalesInvoice(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/sales_invoices?include=category,contact,details,details.product,details.warehouse,payments,payments.transaction,tags,sharings,recurrence_plan,active_e_document"
	request.SalesInvoice.Data.Type = "sales_invoices"
	salesinvoicedata, _ := json.Marshal(request.SalesInvoice)
	client := new(http.Client)
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(salesinvoicedata))
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.SalesInvoice)
	return response
}

func (api *API) ShowSalesInvoice(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/sales_invoices/" + request.SalesInvoice.Data.ID + "?include=category,contact,details,details.product,details.warehouse,payments,payments.transaction,tags,sharings,recurrence_plan,active_e_document"
	client := new(http.Client)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.SalesInvoice)
	return response
}

func (api *API) CancelSalesInvoice(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/sales_invoices/" + request.SalesInvoice.Data.ID + "/cancel"
	client := new(http.Client)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.SalesInvoice)
	return response
}

func (api *API) DeleteSalesInvoice(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/sales_invoices/" + request.SalesInvoice.Data.ID
	client := new(http.Client)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.SalesInvoice)
	return response
}

func (api *API) ArchiveSalesInvoice(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/sales_invoices/" + request.SalesInvoice.Data.ID + "/archive"
	client := new(http.Client)
	req, err := http.NewRequest("PATCH", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.SalesInvoice)
	return response
}

func (api *API) UnarchiveSalesInvoice(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/sales_invoices/" + request.SalesInvoice.Data.ID + "/unarchive"
	client := new(http.Client)
	req, err := http.NewRequest("PATCH", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.SalesInvoice)
	return response
}

func (api *API) CreateEArchive(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/e_archives"
	request.EArchive.Data.Type = "e_archives"
	earchivedata, _ := json.Marshal(request.EArchive)
	client := new(http.Client)
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(earchivedata))
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.EArchive)
	return response
}

func (api *API) ShowEArchive(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/e_archives/" + request.EArchive.Data.ID
	client := new(http.Client)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.EArchive)
	return response
}

func (api *API) CreateEInvoice(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/e_invoices"
	request.EInvoice.Data.Type = "e_invoices"
	einvoicedata, _ := json.Marshal(request.EInvoice)
	client := new(http.Client)
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(einvoicedata))
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.EInvoice)
	return response
}

func (api *API) ShowEInvoice(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/e_invoices/" + request.EInvoice.Data.ID
	client := new(http.Client)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.EInvoice)
	return response
}

func (api *API) ShowEArchivePDF(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/e_archives/" + request.EArchivePDF.Data.ID + "/pdf"
	client := new(http.Client)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.EArchivePDF)
	return response
}

func (api *API) ShowEInvoicePDF(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/e_invoices/" + request.EInvoicePDF.Data.ID + "/pdf"
	client := new(http.Client)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.EInvoicePDF)
	return response
}

func (api *API) ListEInvoiceInboxes(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/e_invoice_inboxes?filter[vkn]=" + request.EInvoiceInboxes.Data.Attributes.VKN
	client := new(http.Client)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.EInvoiceInboxes)
	return response
}

func (api *API) ShowTransaction(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/transactions/" + request.Transaction.Data.ID + "?include=debit_account,credit_account,payments"
	client := new(http.Client)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Transaction)
	return response
}

func (api *API) DeleteTransaction(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/transactions/" + request.Transaction.Data.ID
	client := new(http.Client)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.Transaction)
	return response
}

func (api *API) TrackJob(request *Request) (response Response) {
	endpoint := "https://api.parasut.com/v4/" + api.Config.CompanyID + "/trackable_jobs/" + request.TrackableJob.Data.ID
	client := new(http.Client)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response.TrackableJob)
	return response
}
