/**
 * Copyright 2025-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
	"github.com/coinbase-samples/advanced-trade-sdk-go/credentials"
	"github.com/coinbase-samples/advanced-trade-sdk-go/portfolios"
)

func main() {
	credentials := &credentials.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("ADVANCED_CREDENTIALS")), credentials); err != nil {
		log.Fatalf("unable to deserialize advanced trade credentials JSON: %v", err)
	}

	httpClient, err := client.DefaultHttpClient()
	if err != nil {
		log.Fatalf("unable to load default http client: %v", err)
	}

	restClient := client.NewRestClient(credentials, httpClient)
	portfoliosService := portfolios.NewPortfoliosService(restClient)

	resp, err := portfoliosService.ListPortfolios(context.Background(), &portfolios.ListPortfoliosRequest{})
	if err != nil {
		log.Fatalf("error listing portfolios: %v", err)
	}

	jsonResponse, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("error marshaling response to JSON: %v", err))
	}
	fmt.Println(string(jsonResponse))
}
