/*
NabooApi V1

Testing CashoutServiceService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package naboo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/naboopay/naboo-go"
)

func Test_naboo_CashoutServiceService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CashoutServiceService CashOutOrangeMoneyCashoutOrangeMoneyPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CashoutService.CashOutOrangeMoneyCashoutOrangeMoneyPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CashoutServiceService CashOutWaveCashoutWavePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CashoutService.CashOutWaveCashoutWavePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}