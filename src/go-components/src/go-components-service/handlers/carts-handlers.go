// // Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// // SPDX-License-Identifier: MIT-0

package handlers

import (
	"encoding/json"
	"fmt"
	"go-component-service/models"
	"go-component-service/repos"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gorilla/mux"
)

// Index Handler
func IndexNew(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Carts Web Service")
}

// CartIndex Handler
func CartIndex(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var values []models.Cart
	for _, value := range repos.Carts {
		values = append(values, value)
	}

	if err := json.NewEncoder(w).Encode(values); err != nil {
		panic(err)
	}
}

// CartShowByID Handler
func CartShowByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	cartID := vars["cartID"]

	if err := json.NewEncoder(w).Encode(repos.RepoFindCartByID(cartID)); err != nil {
		panic(err)
	}
}

// CartUpdate Func
func CartUpdate(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var cart models.Cart
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &cart); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	vars := mux.Vars(r)
	cartID := vars["cartID"]

	t := repos.RepoUpdateCart(cartID, cart)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

// CartCreate Func
func CartCreate(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var cart models.Cart
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &cart); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := repos.RepoCreateCart(cart)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

// Sign a payload for Amazon Pay - delegates to a Lambda function for doing this.
func SignAmazonPayPayload(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(awsSession)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var requestBody map[string]interface{}
	json.Unmarshal(body, &requestBody)

	result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("AmazonPaySigningLambda"), Payload: body})
	if err != nil {
		panic(err)
	}

	var responsePayload map[string]interface{}
	json.Unmarshal(result.Payload, &responsePayload)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responsePayload); err != nil {
		panic(err)
	}
}

// // // enableCors
// // func enableCors(w *http.ResponseWriter) {
// // 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// // 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, OPTIONS")
// // 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// // }
