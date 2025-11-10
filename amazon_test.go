package amzsdk

import (
	"context"
	"testing"
)

func TestGetListingsRestrictionsSignature(t *testing.T) {
	// This test ensures the GetListingsRestrictions method exists and has the correct signature
	sdk := NewAmazonSDK("test-client", "test-secret", "test-token")
	ctx := context.Background()
	
	params := ListingsRestrictionsRequestParams{
		ASIN:           "B07TTY5YS8",
		ConditionType:  "new_new",
		SellerID:       "AJI6WKJB10KAL",
		MarketplaceIds: "A13V1IB3VIYZZH",
	}
	
	// We don't make an actual API call since we don't have valid credentials
	// This test just ensures the method signature is correct and compiles
	_, err := sdk.GetListingsRestrictions(ctx, "dummy-token", params)
	
	// We expect an error since we're not making a real API call
	if err == nil {
		t.Log("No error returned (unexpected in test environment)")
	}
}

func TestListingsRestrictionsTypes(t *testing.T) {
	// Test that the response types are correctly structured
	resp := ListingsRestrictionsResponse{
		Restrictions: []Restriction{
			{
				MarketplaceID: "A13V1IB3VIYZZH",
				ConditionType: "new_new",
				Reasons: []Reason{
					{
						ReasonCode: "APPROVAL_REQUIRED",
						Message:    "Test message",
						Links: []Link{
							{
								Resource: "https://example.com",
								Verb:     "GET",
								Title:    "Test Link",
								Type:     "text/html",
							},
						},
					},
				},
			},
		},
	}
	
	if len(resp.Restrictions) != 1 {
		t.Errorf("Expected 1 restriction, got %d", len(resp.Restrictions))
	}
	
	if resp.Restrictions[0].MarketplaceID != "A13V1IB3VIYZZH" {
		t.Errorf("Expected marketplace ID 'A13V1IB3VIYZZH', got '%s'", resp.Restrictions[0].MarketplaceID)
	}
	
	if resp.Restrictions[0].Reasons[0].ReasonCode != "APPROVAL_REQUIRED" {
		t.Errorf("Expected reason code 'APPROVAL_REQUIRED', got '%s'", resp.Restrictions[0].Reasons[0].ReasonCode)
	}
}
