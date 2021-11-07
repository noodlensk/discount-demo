package app

import "github.com/noodlensk/discount-demo/internal/discount/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	GenerateDiscountCodes command.GenerateDiscountCodesHandler
	FetchDiscount         command.FetchDiscountCodeHandler
}
