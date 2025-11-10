package amzsdk

// CatalogItemsRequestParams holds the query parameters for GetCatalogItems
type CatalogItemsRequestParams struct {
	MarketplaceIds  string // comma-separated
	IdentifiersType string // e.g. "GTIN"
	Identifiers     string // comma-separated
	IncludedData    string // comma-separated, e.g. "identifiers,summaries,salesRanks,images,attributes"
}

// FeesEstimateRequestPayload represents a single request item for the GetFeesEstimate API.
// Used as the element type in the request slice for GetFeesEstimate.
type FeesEstimateRequestPayload struct {
	FeesEstimateRequest FeesEstimateRequest `json:"FeesEstimateRequest"` // The main request details
	IdType              string              `json:"IdType"`              // e.g. "ASIN"
	IdValue             string              `json:"IdValue"`             // e.g. "B000IWYOHU"
}

// FeesEstimateRequest represents the inner request object for a fees estimate.
type FeesEstimateRequest struct {
	MarketplaceId       string              `json:"MarketplaceId"`       // Marketplace identifier
	Identifier          string              `json:"Identifier"`          // Unique identifier for the request
	IsAmazonFulfilled   bool                `json:"IsAmazonFulfilled"`   // Whether the item is fulfilled by Amazon
	PriceToEstimateFees PriceToEstimateFees `json:"PriceToEstimateFees"` // Price details for fee estimation
}

// PriceToEstimateFees holds the price information for fee estimation.
type PriceToEstimateFees struct {
	ListingPrice MoneyType `json:"ListingPrice"` // Listing price details
}

// MoneyType represents a currency and amount.
type MoneyType struct {
	CurrencyCode string  `json:"CurrencyCode"` // e.g. "EUR"
	Amount       float64 `json:"Amount"`       // e.g. 10.0
}

// FeesEstimateResponsePayload represents a single response item from the GetFeesEstimate API.
type FeesEstimateResponsePayload struct {
	Status                 string                 `json:"Status"`                 // e.g. "Success"
	FeesEstimateIdentifier FeesEstimateIdentifier `json:"FeesEstimateIdentifier"` // Identifies the estimate
	FeesEstimate           *FeesEstimate          `json:"FeesEstimate,omitempty"` // The actual estimate (may be nil)
}

// FeesEstimateIdentifier identifies a fees estimate in the response.
type FeesEstimateIdentifier struct {
	MarketplaceId         string              `json:"MarketplaceId"`
	IdType                string              `json:"IdType"`
	SellerId              string              `json:"SellerId"`
	SellerInputIdentifier string              `json:"SellerInputIdentifier"`
	IsAmazonFulfilled     bool                `json:"IsAmazonFulfilled"`
	IdValue               string              `json:"IdValue"`
	PriceToEstimateFees   PriceToEstimateFees `json:"PriceToEstimateFees"`
}

// FeesEstimate contains the estimated fees and details.
type FeesEstimate struct {
	TimeOfFeesEstimation string      `json:"TimeOfFeesEstimation"` // ISO8601 timestamp
	TotalFeesEstimate    MoneyType   `json:"TotalFeesEstimate"`    // Total estimated fees
	FeeDetailList        []FeeDetail `json:"FeeDetailList"`        // List of fee details
}

// FeeDetail provides details for a specific fee type.
type FeeDetail struct {
	FeeType               string      `json:"FeeType"` // e.g. "ReferralFee"
	FeeAmount             MoneyType   `json:"FeeAmount"`
	FinalFee              MoneyType   `json:"FinalFee"`
	FeePromotion          MoneyType   `json:"FeePromotion"`
	IncludedFeeDetailList []FeeDetail `json:"IncludedFeeDetailList,omitempty"` // Nested fee details (optional)
}

// CatalogItemsResponse represents the response from the GetCatalogItems API
type CatalogItemsResponse struct {
	NumberOfResults int           `json:"numberOfResults"`
	Items           []CatalogItem `json:"items"`
}

type CatalogItem struct {
	ASIN        string                 `json:"asin"`
	Attributes  map[string]interface{} `json:"attributes"`
	Identifiers []struct {
		MarketplaceId string `json:"marketplaceId"`
		Identifiers   []struct {
			IdentifierType string `json:"identifierType"`
			Identifier     string `json:"identifier"`
		} `json:"identifiers"`
	} `json:"identifiers"`
	Images []struct {
		MarketplaceId string `json:"marketplaceId"`
		Images        []struct {
			Variant string `json:"variant"`
			Link    string `json:"link"`
			Height  int    `json:"height"`
			Width   int    `json:"width"`
		} `json:"images"`
	} `json:"images"`
	SalesRanks []struct {
		MarketplaceId       string `json:"marketplaceId"`
		ClassificationRanks []struct {
			ClassificationId string `json:"classificationId"`
			Title            string `json:"title"`
			Link             string `json:"link"`
			Rank             int    `json:"rank"`
		} `json:"classificationRanks"`
		DisplayGroupRanks []struct {
			WebsiteDisplayGroup string `json:"websiteDisplayGroup"`
			Title               string `json:"title"`
			Link                string `json:"link"`
			Rank                int    `json:"rank"`
		} `json:"displayGroupRanks"`
	} `json:"salesRanks"`
	Summaries []struct {
		MarketplaceId        string `json:"marketplaceId"`
		AdultProduct         bool   `json:"adultProduct"`
		Autographed          bool   `json:"autographed"`
		Brand                string `json:"brand"`
		BrowseClassification struct {
			DisplayName      string `json:"displayName"`
			ClassificationId string `json:"classificationId"`
		} `json:"browseClassification"`
		Color                   string `json:"color"`
		ItemClassification      string `json:"itemClassification"`
		ItemName                string `json:"itemName"`
		Manufacturer            string `json:"manufacturer"`
		Memorabilia             bool   `json:"memorabilia"`
		ModelNumber             string `json:"modelNumber"`
		PackageQuantity         int    `json:"packageQuantity"`
		PartNumber              string `json:"partNumber"`
		ReleaseDate             string `json:"releaseDate"`
		Size                    string `json:"size"`
		Style                   string `json:"style"`
		TradeInEligible         bool   `json:"tradeInEligible"`
		WebsiteDisplayGroup     string `json:"websiteDisplayGroup"`
		WebsiteDisplayGroupName string `json:"websiteDisplayGroupName"`
	} `json:"summaries"`
}
