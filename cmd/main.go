package main

import (
	"context"
	"fmt"
	"net/url"

	swagger "github.com/cjunior1/solace-semp-api-golang"
)

func main() {

	cfg := swagger.NewConfiguration()
	client := swagger.NewAPIClient(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// credenciais
	credencial := &swagger.BasicAuth{}
	credencial.UserName = "rd-ecommerce-admin"
	credencial.Password = "lbqob38eocodjs27pik5ikkb5k"

	ctx = context.WithValue(ctx, swagger.ContextBasicAuth, *credencial)

	// estado atual
	queueName := "xx-demo-producer-happy-path-dev"
	serviceName := "rd-ecommerce"
	subscriptionTopic := "xx/demo/producer/happy/path/dev5"

	resp, _, err := client.MsgVpnApi.GetMsgVpnQueueSubscription(ctx, serviceName, queueName, url.PathEscape(subscriptionTopic), nil)
	if err != nil {
		//fmt.Println(resp.Meta.ResponseCode)
		//		if resp.Meta.Error_.Status == "NOT FOUND" {
		fmt.Println("Topic Not found : ", subscriptionTopic)
		novoTopic := swagger.MsgVpnQueueSubscription{MsgVpnName: serviceName, QueueName: queueName, SubscriptionTopic: subscriptionTopic}
		resp, _, err := client.MsgVpnApi.CreateMsgVpnQueueSubscription(ctx, serviceName, queueName, novoTopic, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("New Topic created : ", resp.Data.SubscriptionTopic)

	} else {

		fmt.Println("Topic found : ", subscriptionTopic)
		fmt.Println("Status=", resp.Meta.ResponseCode)

	}
	//	}

	subs, _, err := client.MsgVpnApi.GetMsgVpnQueueSubscriptions(ctx, serviceName, queueName, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("#############")
	for _, s := range subs.Data {
		fmt.Println(s.SubscriptionTopic)
	}

}
