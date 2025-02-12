// my-portfolio.
//
// API to create my portfolio with hexagonal architecture
//
//     Schemes: https
//     BasePath: /
//     Version: 1.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: https://jhonquirama.com
//
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

package main

import (
	"context"

	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
	runner "github.com/jhonquirama/my-portfolio/runners"
)

func main() {
	ctx := context.Background()

	if err := runner.NewRunner().Run(ctx); err != nil {
		customLogger.Fatal(ctx, err)
	}
}
