package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
)

type SettingId struct {
	SubscriptionId string
	Name           string
}

func NewSettingID(subscriptionId, name string) SettingId {
	return SettingId{
		SubscriptionId: subscriptionId,
		Name:           name,
	}
}

func (id SettingId) String() string {
	segments := []string{
		fmt.Sprintf("Name %q", id.Name),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Setting", segmentsStr)
}

func (id SettingId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/settings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Name)
}

// SettingID parses a Setting ID into an SettingId struct
func SettingID(input string) (*SettingId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := SettingId{
		SubscriptionId: id.SubscriptionID,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.Name, err = id.PopSegment("settings"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
