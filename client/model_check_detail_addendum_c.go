/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing, and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// CheckDetailAddendumC struct for CheckDetailAddendumC
type CheckDetailAddendumC struct {
	// CheckDetailAddendumC ID
	Id string `json:"id,omitempty"`
	// RecordNumber is a number representing the order in which each CheckDetailAddendumC was created. CheckDetailAddendumC shall be in sequential order starting with 1.
	RecordNumber int32 `json:"recordNumber"`
	// EndorsingBankRoutingNumber is a valid routing and transit number indicating the bank that endorsed the check.
	EndorsingBankRoutingNumber string `json:"endorsingBankRoutingNumber"`
	// BOFDEndorsementBusinessDate is the date of endorsement.
	BofdEndorsementBusinessDate time.Time `json:"bofdEndorsementBusinessDate"`
	// EndorsingItemSequenceNumber is a number that identifies the item at the endorsing bank.
	EndorsingBankSequenceNumber string `json:"endorsingBankSequenceNumber,omitempty"`
	// TruncationIndicator identifies if the institution truncated the original check item.
	TruncationIndicator string `json:"truncationIndicator"`
	// EndorsingBankConversionIndicator is a code that indicates the conversion within the processing institution between original paper check, image, and IRD. The indicator is specific to the action of the institution identified in the EndorsingBankRoutingNumber.  * `0` - Did not convert physical document * `1` - Original paper converted to IRD * `2` - Original paper converted to image * `3` - IRD converted to another IRD * `4` - IRD converted to image of IRD * `5` - Image converted to an IRD * `6` - Image converted to another image (e.g., transcoded) * `7` - Did not convert image (e.g., same as source) * `8` - Undetermined
	EndorsingBankConversionIndicator string `json:"endorsingBankConversionIndicator,omitempty"`
	// EndorsingBankCorrectionIndicator identifies whether and how the MICR line of this item was repaired by the creator of this CheckDetailAddendumC Record for fields other than Payor Bank Routing Number and Amount.  * `0` - No Repair * `1` - Repaired (form of repair unknown) * `2` - Repaired without Operator intervention * `3` - Repaired with Operator intervention * `4` - Undetermined if repair has been done or not
	EndorsingBankCorrectionIndicator int32 `json:"endorsingBankCorrectionIndicator,omitempty"`
	// ReturnReason is a code that indicates the reason for non-payment.
	ReturnReason string `json:"returnReason,omitempty"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField,omitempty"`
	// * `0` - Depository Bank (BOFD) - this value is used when the CheckDetailAddendumC Record reflects the Return * `Processing Bank in lieu of BOFD. * `1` - Other Collecting Bank * `2` - Other Returning Bank * `3` - Payor Bank
	EndorsingBankIdentifier int32 `json:"endorsingBankIdentifier,omitempty"`
}
