
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