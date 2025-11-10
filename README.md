
# Example usage of GetCatalogItems:
```go
	ctx := context.Background()
	sdk := NewAmazonSDK("<client_id>", "<client_secret>", "<refresh_token>")
	tokenResp, err := sdk.GetAccessToken(ctx)
	if err != nil {
		panic(err)
	}
	params := CatalogItemsRequestParams{
		MarketplaceIds:  "A13V1IB3VIYZZH",
		IdentifiersType: "GTIN",
		Identifiers:     "00000080466437",
		IncludedData:    "identifiers,summaries,salesRanks,images,attributes",
	}
	catalogResp, err := sdk.GetCatalogItems(ctx, tokenResp.AccessToken, params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found %d items\n", catalogResp.NumberOfResults)
	for _, item := range catalogResp.Items {
		fmt.Println("ASIN:", item.ASIN)
		// ... process item ...
	}
```

# Example usage of GetListingsRestrictions:
```go
	ctx := context.Background()
	sdk := NewAmazonSDK("<client_id>", "<client_secret>", "<refresh_token>")
	tokenResp, err := sdk.GetAccessToken(ctx)
	if err != nil {
		panic(err)
	}
	params := ListingsRestrictionsRequestParams{
		ASIN:           "B07TTY5YS8",
		ConditionType:  "new_new",
		SellerID:       "AJI6WKJB10KAL",
		MarketplaceIds: "A13V1IB3VIYZZH",
	}
	restrictionsResp, err := sdk.GetListingsRestrictions(ctx, tokenResp.AccessToken, params)
	if err != nil {
		panic(err)
	}
	for _, restriction := range restrictionsResp.Restrictions {
		fmt.Printf("Marketplace: %s, Condition: %s\n", restriction.MarketplaceID, restriction.ConditionType)
		for _, reason := range restriction.Reasons {
			fmt.Printf("  Reason: %s - %s\n", reason.ReasonCode, reason.Message)
			for _, link := range reason.Links {
				fmt.Printf("    Link: %s (%s)\n", link.Title, link.Resource)
			}
		}
	}
```