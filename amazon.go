package amzsdk

import (
	"context"
	"fmt"

	"github.com/imroc/req/v3"
)

const (
	PROD_API_ENDPOINT = "https://sellingpartnerapi-eu.amazon.com"
)

// AmazonAuthTokenResponse represents the response from the Amazon OAuth token endpoint
type AmazonAuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// AmazonSDK provides methods to interact with Amazon APIs
// You can expand this struct with more config as needed
// (e.g., logger, http client, etc.)
type AmazonSDK struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
}

// NewAmazonSDK creates a new AmazonSDK instance
func NewAmazonSDK(clientID, clientSecret, refreshToken string) *AmazonSDK {
	return &AmazonSDK{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
	}
}

// GetAccessToken exchanges a refresh token for an access token
func (sdk *AmazonSDK) GetAccessToken(ctx context.Context) (*AmazonAuthTokenResponse, error) {
	var resp AmazonAuthTokenResponse
	r, err := req.C().R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type":    "refresh_token",
			"refresh_token": sdk.RefreshToken,
			"client_id":     sdk.ClientID,
			"client_secret": sdk.ClientSecret,
		}).
		SetSuccessResult(&resp).
		Post("https://api.amazon.com/auth/o2/token")
	if err != nil {
		return nil, err
	}
	if !r.IsSuccessState() {
		return nil, fmt.Errorf("failed to get access token: %s", r.String())
	}
	return &resp, nil
}

// BatchItemOfferRequest represents a single item offer request
type BatchItemOfferRequest struct {
	MarketplaceId string `json:"MarketplaceId"`
	ItemCondition string `json:"ItemCondition"`
	CustomerType  string `json:"CustomerType"`
	URI           string `json:"uri"`
	Method        string `json:"method"`
}

// BatchItemOffersPayload is the payload for the batch item offers API
type BatchItemOffersPayload struct {
	Requests []BatchItemOfferRequest `json:"requests"`
}

// BatchItemOffersResponse represents the response from the batch item offers API
type BatchItemOffersResponse struct {
	Responses []struct {
		Headers struct {
			XAmznRequestId string `json:"x-amzn-RequestId"`
			Date           string `json:"Date"`
		} `json:"headers"`
		Status struct {
			StatusCode   int64  `json:"statusCode"`
			ReasonPhrase string `json:"reasonPhrase"`
		} `json:"status"`
		Body struct {
			Payload struct {
				MarketplaceId string `json:"marketplaceId"`
				Identifier    struct {
					ASIN          string `json:"ASIN"`
					MarketplaceId string `json:"MarketplaceId"`
					ItemCondition string `json:"ItemCondition"`
				} `json:"Identifier"`
				ASIN    string `json:"ASIN"`
				Summary *struct {
					BuyBoxPrices []struct {
						Condition   string `json:"condition"`
						LandedPrice struct {
							CurrencyCode string  `json:"CurrencyCode"`
							Amount       float64 `json:"Amount"`
						} `json:"LandedPrice"`
						Shipping struct {
							CurrencyCode string  `json:"CurrencyCode"`
							Amount       float64 `json:"Amount"`
						} `json:"Shipping"`
						ListingPrice struct {
							CurrencyCode string  `json:"CurrencyCode"`
							Amount       float64 `json:"Amount"`
						} `json:"ListingPrice"`
					} `json:"BuyBoxPrices"`
					BuyBoxEligibleOffers []struct {
						Condition          string `json:"condition"`
						FulfillmentChannel string `json:"fulfillmentChannel"`
						OfferCount         int64  `json:"OfferCount"`
					} `json:"BuyBoxEligibleOffers"`
					LowestPrices []struct {
						Condition          string `json:"condition"`
						FulfillmentChannel string `json:"fulfillmentChannel"`
						LandedPrice        struct {
							CurrencyCode string  `json:"CurrencyCode"`
							Amount       float64 `json:"Amount"`
						} `json:"LandedPrice"`
						Shipping struct {
							CurrencyCode string  `json:"CurrencyCode"`
							Amount       float64 `json:"Amount"`
						} `json:"Shipping"`
						ListingPrice struct {
							CurrencyCode string  `json:"CurrencyCode"`
							Amount       float64 `json:"Amount"`
						} `json:"ListingPrice"`
					} `json:"LowestPrices"`
					NumberOfOffers []struct {
						Condition          string `json:"condition"`
						FulfillmentChannel string `json:"fulfillmentChannel"`
						OfferCount         int64  `json:"OfferCount"`
					} `json:"NumberOfOffers"`
					TotalOfferCount int64 `json:"TotalOfferCount"`
					SalesRankings   []struct {
						Rank              int64  `json:"Rank"`
						ProductCategoryId string `json:"ProductCategoryId"`
					} `json:"SalesRankings"`
				} `json:"Summary"`
				Offers []struct {
					ShippingTime struct {
						MinimumHours     int64  `json:"minimumHours"`
						MaximumHours     int64  `json:"maximumHours"`
						AvailabilityType string `json:"availabilityType"`
					} `json:"ShippingTime"`
					IsFulfilledByAmazon bool `json:"IsFulfilledByAmazon"`
					ListingPrice        struct {
						CurrencyCode string  `json:"CurrencyCode"`
						Amount       float64 `json:"Amount"`
					} `json:"ListingPrice"`
					IsBuyBoxWinner bool   `json:"IsBuyBoxWinner"`
					SellerId       string `json:"SellerId"`
					Shipping       struct {
						CurrencyCode string  `json:"CurrencyCode"`
						Amount       float64 `json:"Amount"`
					} `json:"Shipping"`
					ShipsFrom *struct {
						Country string `json:"Country"`
					} `json:"ShipsFrom,omitempty"`
					SubCondition         string `json:"SubCondition"`
					IsFeaturedMerchant   bool   `json:"IsFeaturedMerchant"`
					SellerFeedbackRating struct {
						FeedbackCount                int64   `json:"FeedbackCount"`
						SellerPositiveFeedbackRating float64 `json:"SellerPositiveFeedbackRating"`
					} `json:"SellerFeedbackRating"`
					PrimeInformation struct {
						IsPrime         bool `json:"IsPrime"`
						IsNationalPrime bool `json:"IsNationalPrime"`
					} `json:"PrimeInformation"`
				} `json:"Offers"`
				Status        string `json:"status"`
				ItemCondition string `json:"ItemCondition"`
			} `json:"payload"`
		} `json:"body"`
		Request struct {
			MarketplaceId string `json:"MarketplaceId"`
			CustomerType  string `json:"CustomerType"`
			ItemCondition string `json:"ItemCondition"`
			Asin          string `json:"Asin"`
		} `json:"request"`
	} `json:"responses"`
}

// GetBatchItemOffers calls the batch item offers endpoint
func (sdk *AmazonSDK) GetBatchItemOffers(ctx context.Context, accessToken string, payload BatchItemOffersPayload) (*BatchItemOffersResponse, error) {
	var resp BatchItemOffersResponse
	r, err := req.C().EnableDumpAllWithoutRequest().R().
		SetContext(ctx).
		SetHeader("x-amz-access-token", accessToken).
		SetHeader("user-agent", "elevate-seller").
		SetHeader("Content-Type", "application/json").
		SetBody(&payload).
		SetSuccessResult(&resp).
		Post(PROD_API_ENDPOINT + "/batches/products/pricing/v0/itemOffers")
	if err != nil {
		return nil, err
	}
	if !r.IsSuccessState() {
		return nil, fmt.Errorf("failed to get batch item offers: %s", r.String())
	}
	return &resp, nil
}

// GetCatalogItems calls the Amazon SP API Catalog Items endpoint
func (sdk *AmazonSDK) GetCatalogItems(ctx context.Context, accessToken string, params CatalogItemsRequestParams) (*CatalogItemsResponse, error) {
	var resp CatalogItemsResponse
	r, err := req.C().R().
		SetContext(ctx).
		SetHeader("x-amz-access-token", accessToken).
		SetHeader("user-agent", "elevate-seller").
		SetQueryParams(map[string]string{
			"marketplaceIds":  params.MarketplaceIds,
			"identifiersType": params.IdentifiersType,
			"identifiers":     params.Identifiers,
			"includedData":    params.IncludedData,
		}).
		SetSuccessResult(&resp).
		Get(PROD_API_ENDPOINT + "/catalog/2022-04-01/items")
	if err != nil {
		return nil, err
	}
	if !r.IsSuccessState() {
		return nil, fmt.Errorf("failed to get catalog items: %s", r.String())
	}
	return &resp, nil
}

// GetFeesEstimate calls the Amazon SP API Fees Estimate endpoint.
//
// It takes a slice of FeesEstimateRequestPayload and returns a slice of FeesEstimateResponsePayload.
//
// Example usage:
//
//	reqs := []FeesEstimateRequestPayload{ ... }
//	resp, err := sdk.GetFeesEstimate(ctx, accessToken, reqs)
//
// See Amazon documentation for details:
// https://developer-docs.amazon.com/sp-api/docs/products-fees-api-v0-reference#post-feesestimate
//
// The accessToken must be a valid SP-API access token. The payload must match the API schema.
// Returns a slice of FeesEstimateResponsePayload or an error.
func (sdk *AmazonSDK) GetFeesEstimate(ctx context.Context, accessToken string, payload []FeesEstimateRequestPayload) ([]FeesEstimateResponsePayload, error) {
	var resp []FeesEstimateResponsePayload
	r, err := req.C().R().
		SetContext(ctx).
		SetHeader("x-amz-access-token", accessToken).
		SetHeader("user-agent", "elevate-seller").
		SetHeader("Content-Type", "application/json").
		SetBody(&payload).
		SetSuccessResult(&resp).
		Post(PROD_API_ENDPOINT + "/products/fees/v0/feesEstimate")
	if err != nil {
		return nil, err
	}
	if !r.IsSuccessState() {
		return nil, fmt.Errorf("failed to get fees estimate: %s", r.String())
	}
	return resp, nil
}
