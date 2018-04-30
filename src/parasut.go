package parasut

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"parasut/config"
	"strings"

	"github.com/pasztorpisti/qs"
)

type API struct {
	Client struct {
		ClientID     string `json:"client_id,omitempty"`
		ClientSecret string `json:"client_secret,omitempty"`
		Username     string `json:"username,omitempty"`
		Password     string `json:"password,omitempty"`
		GrantType    string `json:"grant_type,omitempty"`
		RedirectURI  string `json:"redirect_uri,omitempty"`
	}

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
	Contacts struct {
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
				City        string `json:"city,omitempty"`
				District    string `json:"district,omitempty"`
				Address     string `json:"address,omitempty"`
				Phone       string `json:"phone,omitempty"`
				Fax         string `json:"fax,omitempty"`
				IsAbroad    bool   `json:"is_abroad,omitempty"`
				Archived    bool   `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Category struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"category,omitempty"`
				ContactPortal struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"contact_portal,omitempty"`
				ContactPeople struct {
					Data []struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"contact_people,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	Employees struct {
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
				Category struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"category,omitempty"`
				ManagedByUser struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"managed_by_user,omitempty"`
				ManagedByUserRole struct {
					Data []struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"managed_by_user_role,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	SalesInvoices struct {
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
				City                string      `json:"city,omitempty"`
				District            string      `json:"district,omitempty"`
				OrderNo             string      `json:"order_no,omitempty"`
				OrderDate           string      `json:"order_date,omitempty"`
				ShipmentAddress     string      `json:"shipment_addres,omitempty"`
				IsAbroad            bool        `json:"is_abroad,omitempty"`
				Archived            bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Details struct {
					Fill struct {
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
							Product struct {
								Data struct {
									Type string `json:"type,omitempty"`
									ID   string `json:"id,omitempty"`
								} `json:"data,omitempty"`
							} `json:"product,omitempty"`
						} `json:"relationships,omitempty"`
					} `json:"-"`
					Data []struct {
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
							Product struct {
								Data struct {
									Type string `json:"type,omitempty"`
									ID   string `json:"id,omitempty"`
								} `json:"data,omitempty"`
							} `json:"product,omitempty"`
						} `json:"relationships,omitempty"`
					} `json:"data,omitempty"`
				} `json:"details,omitempty"`
				Contact struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"contact,omitempty"`
				Category struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"category,omitempty"`
				Tags struct {
					Data []struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"tags,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	Payments struct {
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Description  string `json:"description,omitempty"`
				AccountID    string `json:"account_id,omitempty"`
				Date         string `json:"date,omitempty"`
				Amount       string `json:"amount,omitempty"`
				Currency     string `json:"currency,omitempty"`
				ExchangeRate string `json:"exchange_rate,omitempty"`
				Notes        string `json:"notes,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Payable struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"payable,omitempty"`
				Transaction struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"transaction,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	EArchives struct {
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
		} `json:"data,omitempty"`
	}

	EInvoices struct {
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
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	EArchivePDF struct {
		Data struct {
			Type string `json:"type,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"data,omitempty"`
	}

	EInvoicePDF struct {
		Data struct {
			Type string `json:"type,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"data,omitempty"`
	}
}

type Response struct {
	Contacts struct {
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
				City        string      `json:"city,omitempty"`
				District    string      `json:"district,omitempty"`
				Address     string      `json:"address,omitempty"`
				Phone       string      `json:"phone,omitempty"`
				Fax         string      `json:"fax,omitempty"`
				IsAbroad    bool        `json:"is_abroad,omitempty"`
				Archived    bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Category struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"category,omitempty"`
				ContactPortal struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"contact_portal,omitempty"`
				ContactPeople struct {
					Data []struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"contact_people,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	Employees struct {
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
				Category struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"category,omitempty"`
				ManagedByUser struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"managed_by_user,omitempty"`
				ManagedByUserRole struct {
					Data []struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"managed_by_user_role,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	Accounts struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Balance             json.Number `json:"balance,omitempty"`
				UsedFor             string      `json:"used_for,omitempty"`
				LastUsedAt          string      `json:"last_used_at,omitempty"`
				LastAdjustmentDate  string      `json:"last_adjustment_date,omitempty"`
				AssociateEmail      string      `json:"associate_email,omitempty"`
				CreatedAt           string      `json:"created_at,omitempty"`
				UpdatedAt           string      `json:"updated_at,omitempty"`
				Name                string      `json:"name,omitempty"`
				Currency            string      `json:"currency,omitempty"`
				AccountType         string      `json:"account_type,omitempty"`
				BankName            string      `json:"bank_name,omitempty"`
				BankBranch          string      `json:"bank_branch,omitempty"`
				BankAccountNo       string      `json:"bank_account_no,omitempty"`
				BankIntegrationType string      `json:"bank_integration_type,omitempty"`
				IBAN                string      `json:"iban,omitempty"`
				Archived            bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	Products struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				CreatedAt                   string      `json:"created_at,omitempty"`
				UpdatedAt                   string      `json:"updated_at,omitempty"`
				Code                        string      `json:"code,omitempty"`
				Name                        string      `json:"name,omitempty"`
				VatRate                     json.Number `json:"vat_rate,omitempty"`
				SalesExciseDuty             json.Number `json:"sales_excise_duty,omitempty"`
				SalesExciseDutyCode         string      `json:"sales_excise_duty_code,omitempty"`
				SalesExciseDutyType         string      `json:"sales_excise_duty_type,omitempty"`
				SalesInvoiceDetailsCount    json.Number `json:"sales_invoice_details_count,omitempty"`
				PurchaseExciseDuty          json.Number `json:"purchase_excise_duty,omitempty"`
				PurchaseExciseDutyType      string      `json:"purchase_excise_duty_type,omitempty"`
				PurchaseInvoiceDetailsCount json.Number `json:"purchase_invoice_details_count,omitempty"`
				Unit                        string      `json:"unit,omitempty"`
				Currency                    string      `json:"currency,omitempty"`
				BuyingCurrency              string      `json:"buying_currency,omitempty"`
				ListPrice                   json.Number `json:"list_price,omitempty"`
				ListPriceInTrl              json.Number `json:"list_price_in_trl,omitempty"`
				BuyingPrice                 json.Number `json:"buying_price,omitempty"`
				BuyingPriceInTrl            json.Number `json:"buying_price_in_trl,omitempty"`
				InitialStockCount           json.Number `json:"initial_stock_count,omitempty"`
				StockCount                  json.Number `json:"stock_count,omitempty"`
				CommunicationsTaxRate       json.Number `json:"communications_tax_rate,omitempty"`
				InventoryTracking           bool        `json:"inventory_tracking,omitempty"`
				Archived                    bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	BankFees struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Remaining      json.Number `json:"remaining,omitempty"`
				RemainingInTrl json.Number `json:"remaining_in_trl,omitempty"`
				TotalPaid      json.Number `json:"total_paid,omitempty"`
				NetTotal       json.Number `json:"net_total,omitempty"`
				ExchangeRate   json.Number `json:"exchange_rate,omitempty"`
				Currency       string      `json:"currency,omitempty"`
				CreatedAt      string      `json:"created_at,omitempty"`
				UpdatedAt      string      `json:"updated_at,omitempty"`
				IssueDate      string      `json:"issue_date,omitempty"`
				DueDate        string      `json:"due_date,omitempty"`
				Description    string      `json:"description,omitempty"`
				Archived       bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	Salaries struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Remaining      json.Number `json:"remaining,omitempty"`
				RemainingInTrl json.Number `json:"remaining_in_trl,omitempty"`
				TotalPaid      json.Number `json:"total_paid,omitempty"`
				NetTotal       json.Number `json:"net_total,omitempty"`
				ExchangeRate   json.Number `json:"exchange_rate,omitempty"`
				Currency       string      `json:"currency,omitempty"`
				CreatedAt      string      `json:"created_at,omitempty"`
				UpdatedAt      string      `json:"updated_at,omitempty"`
				IssueDate      string      `json:"issue_date,omitempty"`
				DueDate        string      `json:"due_date,omitempty"`
				Description    string      `json:"description,omitempty"`
				Archived       bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	Taxes struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Remaining      json.Number `json:"remaining,omitempty"`
				RemainingInTrl json.Number `json:"remaining_in_trl,omitempty"`
				TotalPaid      json.Number `json:"total_paid,omitempty"`
				NetTotal       json.Number `json:"net_total,omitempty"`
				CreatedAt      string      `json:"created_at,omitempty"`
				UpdatedAt      string      `json:"updated_at,omitempty"`
				IssueDate      string      `json:"issue_date,omitempty"`
				DueDate        string      `json:"due_date,omitempty"`
				Description    string      `json:"description,omitempty"`
				Archived       bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	Transactions struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				TransactionType   string      `json:"transaction_type,omitempty"`
				HumanizedTypeName string      `json:"humanized_type_name,omitempty"`
				CreatedAt         string      `json:"created_at,omitempty"`
				UpdatedAt         string      `json:"updated_at,omitempty"`
				Date              string      `json:"date,omitempty"`
				DebitAmount       json.Number `json:"debit_amount,omitempty"`
				AmountInTrl       json.Number `json:"amount_in_trl,omitempty"`
				DebitCurrency     string      `json:"debit_currency,omitempty"`
				CreditAmount      json.Number `json:"credit_amount,omitempty"`
				CreditCurrency    string      `json:"credit_currency,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	Payments struct {
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				Description  string `json:"description,omitempty"`
				AccountID    string `json:"account_id,omitempty"`
				Date         string `json:"date,omitempty"`
				Amount       string `json:"amount,omitempty"`
				Currency     string `json:"currency,omitempty"`
				ExchangeRate string `json:"exchange_rate,omitempty"`
				Notes        string `json:"notes,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Payable struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"payable,omitempty"`
				Transaction struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"transaction,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	SalesInvoices struct {
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
				City                   string      `json:"city,omitempty"`
				District               string      `json:"district,omitempty"`
				OrderNo                string      `json:"order_no,omitempty"`
				OrderDate              string      `json:"order_date,omitempty"`
				ShipmentAddress        string      `json:"shipment_addres,omitempty"`
				IsAbroad               bool        `json:"is_abroad,omitempty"`
				Archived               bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Details struct {
					Data []struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"details,omitempty"`
				Contact struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"contact,omitempty"`
				Category struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"category,omitempty"`
				Tags struct {
					Data []struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"tags,omitempty"`
				Payments struct {
					Data []struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"payments,omitempty"`
				Sharings struct {
					Data []struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"sharings,omitempty"`
				RecurrencePlan struct {
					Data []struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"recurrence_plan,omitempty"`
				ActiveEDocument struct {
					Data struct {
						Type string `json:"type,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"data,omitempty"`
				} `json:"active_e_document,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"data,omitempty"`
	}

	PurchaseBills struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
			Type       string `json:"type,omitempty"`
			ID         string `json:"id,omitempty"`
			Attributes struct {
				TotalPaid                   json.Number `json:"total_paid,omitempty"`
				NetTotal                    json.Number `json:"net_total,omitempty"`
				GrossTotal                  json.Number `json:"gross_total,omitempty"`
				TotalExciseDuty             json.Number `json:"total_excise_duty,omitempty"`
				TotalCommunicationsTax      json.Number `json:"total_communications_tax,omitempty"`
				TotalVat                    json.Number `json:"total_vat,omitempty"`
				TotalDiscount               json.Number `json:"total_discount,omitempty"`
				TotalInvoiceDiscount        json.Number `json:"total_invoice_discount,omitempty"`
				Remaining                   json.Number `json:"remaining,omitempty"`
				RemainingInTrl              json.Number `json:"remaining_in_trl,omitempty"`
				SharingsCount               json.Number `json:"sharings_count,omitempty"`
				EInvoicesCount              json.Number `json:"e_invoices_count,omitempty"`
				RemainingReimbursement      json.Number `json:"remaining_reimbursement,omitempty"`
				RemainingReimbursementInTrl json.Number `json:"remaining_reimbursement_in_trl,omitempty"`
				ExchangeRate                json.Number `json:"exchange_rate,omitempty"`
				VatWithholding              json.Number `json:"vat_withholding,omitempty"`
				Withholding                 json.Number `json:"withholding,omitempty"`
				WithholdingRate             json.Number `json:"withholding_rate,omitempty"`
				VatWithholdingRate          json.Number `json:"vat_withholding_rate,omitempty"`
				InvoiceDiscount             json.Number `json:"invoice_discount,omitempty"`
				PaymentStatus               string      `json:"payment_status,omitempty"`
				CreatedAt                   string      `json:"created_at,omitempty"`
				UpdatedAt                   string      `json:"updated_at,omitempty"`
				ItemType                    string      `json:"item_type,omitempty"`
				Description                 string      `json:"description,omitempty"`
				IssueDate                   string      `json:"issue_date,omitempty"`
				DueDate                     string      `json:"due_date,omitempty"`
				InvoiceNo                   string      `json:"invoice_no,omitempty"`
				Currency                    string      `json:"currency,omitempty"`
				InvoiceDiscountType         string      `json:"invoice_discount_type,omitempty"`
				IsDetailed                  bool        `json:"is_detailed,omitempty"`
				Archived                    bool        `json:"archived,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"data,omitempty"`
	}

	EInvoiceInboxes struct {
		Errors []struct {
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"errors,omitempty"`
		Data struct {
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

	EArchives struct {
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
		} `json:"data,omitempty"`
	}

	EInvoices struct {
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
		} `json:"data,omitempty"`
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
}

func (api *API) Authorize() bool {
	var data interface{}
	api.Client.ClientID = config.ClientID
	api.Client.ClientSecret = config.ClientSecret
	api.Client.Username = config.Username
	api.Client.Password = config.Password
	api.Client.GrantType = config.GrantType
	api.Client.RedirectURI = config.RedirectURI
	apidata, _ := qs.Marshal(api.Client)
	cli := http.Client{}
	req, err := http.NewRequest("POST", config.TokenURL, strings.NewReader(apidata))
	if err != nil {
		fmt.Println(err)
		return false
	}
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	json.Unmarshal(bytes, &api.Authentication)
	return true
}

func (api *API) CreateContact(request Request) (response Response) {
	var (
		apiurl string
		data   interface{}
	)
	apiurl = config.APIURL + config.CompanyID + "/contacts?include=category,contact_portal,contact_people"
	contactdata, _ := json.Marshal(request.Contacts)
	cli := http.Client{}
	req, err := http.NewRequest("POST", apiurl, bytes.NewReader(contactdata))
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return response
	}
	json.Unmarshal(bytes, &response.Contacts)
	return response
}

func (api *API) ShowContact(request Request) (response Response) {
	var (
		apiurl string
		data   interface{}
	)
	apiurl = config.APIURL + config.CompanyID + "/contacts/" + request.Contacts.Data.ID + "?include=category,contact_portal,contact_people"
	cli := http.Client{}
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return response
	}
	json.Unmarshal(bytes, &response.Contacts)
	return response
}

func (api *API) CreateEmployee(request Request) (response Response) {
	var (
		apiurl string
		data   interface{}
	)
	apiurl = config.APIURL + config.CompanyID + "/employees?include=category,managed_by_user,managed_by_user_role"
	employeedata, _ := json.Marshal(request.Employees)
	cli := http.Client{}
	req, err := http.NewRequest("POST", apiurl, bytes.NewReader(employeedata))
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return response
	}
	json.Unmarshal(bytes, &response.Employees)
	return response
}

func (api *API) ShowEmployee(request Request) (response Response) {
	var (
		apiurl string
		data   interface{}
	)
	apiurl = config.APIURL + config.CompanyID + "/employees/" + request.Employees.Data.ID + "?include=category,managed_by_user,managed_by_user_role"
	cli := http.Client{}
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return response
	}
	json.Unmarshal(bytes, &response.Employees)
	return response
}

func (api *API) CreateSalesInvoice(request Request) (response Response) {
	var (
		apiurl string
		data   interface{}
	)
	apiurl = config.APIURL + config.CompanyID + "/sales_invoices?include=category,contact,details,payments,tags,sharings,recurrence_plan,active_e_document"
	salesinvoicedata, _ := json.Marshal(request.SalesInvoices)
	cli := http.Client{}
	fmt.Println(string(salesinvoicedata))
	req, err := http.NewRequest("POST", apiurl, bytes.NewReader(salesinvoicedata))
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return response
	}
	json.Unmarshal(bytes, &response.SalesInvoices)
	return response
}

func (api *API) ShowSalesInvoice(request Request) (response Response) {
	var (
		apiurl string
		data   interface{}
	)
	apiurl = config.APIURL + config.CompanyID + "/sales_invoices/" + request.SalesInvoices.Data.ID + "?include=category,contact,details,payments,tags,sharings,recurrence_plan,active_e_document"
	cli := http.Client{}
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return response
	}
	json.Unmarshal(bytes, &response.SalesInvoices)
	return response
}

func (api *API) PaySalesInvoice(request Request) (response Response) {
	var (
		apiurl string
		data   interface{}
	)
	apiurl = config.APIURL + config.CompanyID + "/sales_invoices/" + request.SalesInvoices.Data.ID + "/payments"
	paymentdata, _ := json.Marshal(request.Payments)
	cli := http.Client{}
	fmt.Println(string(paymentdata))
	req, err := http.NewRequest("POST", apiurl, bytes.NewReader(paymentdata))
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return response
	}
	json.Unmarshal(bytes, &response.Payments)
	return response
}

func (api *API) ShowEArchive(request Request) (response Response) {
	var (
		apiurl string
		data   interface{}
	)
	apiurl = config.APIURL + config.CompanyID + "/e_archives/" + request.EArchives.Data.ID
	cli := http.Client{}
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return response
	}
	json.Unmarshal(bytes, &response.EArchives)
	return response
}

func (api *API) ShowEInvoice(request Request) (response Response) {
	var (
		apiurl string
		data   interface{}
	)
	apiurl = config.APIURL + config.CompanyID + "/e_invoices/" + request.EInvoices.Data.ID
	cli := http.Client{}
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return response
	}
	json.Unmarshal(bytes, &response.EInvoices)
	return response
}

func (api *API) ShowEArchivePDF(request Request) (response Response) {
	var (
		apiurl string
		data   interface{}
	)
	apiurl = config.APIURL + config.CompanyID + "/e_archives/" + request.EArchivePDF.Data.ID + "/pdf"
	cli := http.Client{}
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return response
	}
	json.Unmarshal(bytes, &response.EArchivePDF)
	return response
}

func (api *API) ShowEInvoicePDF(request Request) (response Response) {
	var (
		apiurl string
		data   interface{}
	)
	apiurl = config.APIURL + config.CompanyID + "/e_invoices/" + request.EInvoicePDF.Data.ID + "/pdf"
	cli := http.Client{}
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Authorization", "Bearer "+api.Authentication.AccessToken)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&data)
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return response
	}
	json.Unmarshal(bytes, &response.EInvoicePDF)
	return response
}
