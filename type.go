package amzsdk

// CatalogItemsRequestParams holds the query parameters for GetCatalogItems
type CatalogItemsRequestParams struct {
	MarketplaceIds  string // comma-separated
	IdentifiersType string // e.g. "GTIN"
	Identifiers     string // comma-separated
	IncludedData    string // comma-separated, e.g. "identifiers,summaries,salesRanks,images,attributes"
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

