package azure

func GetSubscriptionResourceList(subscriptionId string) string {
	return "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resources?api-version=2021-04-01"
}
