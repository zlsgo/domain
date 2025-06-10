package domain

import (
	"context"
	"testing"

)

func TestGetTxt(t *testing.T) {
	ctx := context.Background()
	domain := "qq.com"

	// test default dns
	client := NewClient()
	_, err := client.GetTxt(ctx, domain)
	if err != nil {
		t.Fatal(err)
	}

	// test custom dns
	client = NewClient("1.1.1.1")
	_, err = client.GetTxt(ctx, domain)
	if err != nil {
		t.Fatal(err)
	}
}
