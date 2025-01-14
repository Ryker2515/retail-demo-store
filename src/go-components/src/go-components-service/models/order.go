// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package models

// Order Struct
type Order struct {
	ID               string        `json:"id" yaml:"id"`
	Username         string        `json:"username" yaml:"username"`
	Items            OrderItems    `json:"items" yaml:"items"`
	Total            float32       `json:"total" yaml:"total"`
	BillingAddress   Address       `json:"billing_address" yaml:"billing_address"`
	ShippingAddress  Address       `json:"shipping_address" yaml:"shipping_address"`
	CollectionPhone  string        `json:"collection_phone" yaml:"collection_phone"`
	DeliveryType     string        `json:"delivery_type" yaml:"delivery_type"`
	DeliveryStatus   string        `json:"delivery_status" yaml:"delivery_status"`
	DeliveryComplete bool          `json:"delivery_complete" yaml:"delivery_complete"`
	Channel          string        `json:"channel" yaml:"channel"`
	ChannelDetail    ChannelDetail `json:"channel_detail" yaml:"channel_detail"`
}

// Orders Array
type Orders []Order

// OrderItem Struct
type OrderItem struct {
	ProductID   string  `json:"product_id" yaml:"product_id"`
	ProductName string  `json:"product_name" yaml:"product_name"`
	Quantity    int     `json:"quantity" yaml:"quantity"`
	Price       float32 `json:"price" yaml:"price"`
}

// OrderItems Array
type OrderItems []OrderItem

// ChannelDetail Struct
type ChannelDetail struct {
	ChannelId int    `json:"channel_id" yaml:"channel_id"`
	ChnnelGeo string `json:"channel_geo" yaml:"channel_geo"`
}
