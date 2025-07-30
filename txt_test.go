package domain_test

import (
	"context"
	"testing"

	"github.com/zlsgo/domain"
)

func TestGetTxt(t *testing.T) {
	ctx := context.Background()
	domainToTest := "qq.com"

	// test default dns
	client := domain.NewClient()
	_, err := client.GetTxt(ctx, domainToTest)
	if err != nil {
		t.Fatal(err)
	}

	// test custom dns
	client = domain.NewClient("1.1.1.1")
	_, err = client.GetTxt(ctx, domainToTest)
	if err != nil {
		t.Fatal(err)
	}
}
