package domain

import (
	"context"
	"testing"
)

func TestGetTxt(t *testing.T) {
	ctx := context.Background()
	domainToTest := "qq.com"

	client := &Client{
		resolver: testResolver{
			lookupTXT: func(context.Context, string) ([]string, error) {
				return []string{"v=spf1 include:spf.mail.qq.com ~all"}, nil
			},
		},
	}
	_, err := client.GetTxt(ctx, domainToTest)
	if err != nil {
		t.Fatal(err)
	}

	client = &Client{
		resolver: testResolver{
			lookupTXT: func(context.Context, string) ([]string, error) {
				return []string{"v=spf1 include:spf.mail.qq.com ~all"}, nil
			},
		},
	}
	_, err = client.GetTxt(ctx, domainToTest)
	if err != nil {
		t.Fatal(err)
	}
}
