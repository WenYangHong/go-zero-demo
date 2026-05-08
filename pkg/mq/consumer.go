package mq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	model2 "go-zero-mall/app/order/service/model"
)

func StartOrderConsumer(model model2.HomestayOrderModel) {

	msgs, err := Channel.Consume(
		"order.release.queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	go func() {

		for msg := range msgs {

			var data OrderDelayMsg

			_ = json.Unmarshal(msg.Body, &data)

			order, err := model.FindOneBySn(context.Background(), data.OrderSn)
			if err != nil {
				continue
			}
			log.Printf(" 订单编号: %s\n", order.Sn)

			// 只要不是待支付 - 返回
			if order.TradeState != model2.HomestayOrderTradeStateWaitPay {
				continue
			}

			// 自动取消
			order.TradeState = model2.HomestayOrderTradeStateCancel
			order.UpdateTime = time.Now()
			_, err = model.Update(context.Background(), nil, order)

			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf("订单自动取消: %s\n", order.Sn)
		}
	}()
}
