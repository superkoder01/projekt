package mongodb

const (
	INVOICE_COLLECTION_NAME = "invoices"
	INVOICE_NUMBER_KEY      = "payload.invoiceDetails.number"
	INVOICE_ISSUE_DATE      = "payload.invoiceDetails.issueDt"
	INVOICE_CUSTOMER_ID     = "payload.invoiceDetails.registrationNumber"
	HEADER_CONTENT_TYPE     = "header.content.type"
	INVOICE                 = "invoice"
)
