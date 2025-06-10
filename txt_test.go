package domain

import (
	"context"
	"testing"

	"github.com/sohaha/zlsgo"
)

func TestGetTxt(t *testing.T) {
	tt := zlsgo.NewTest(t)
	ctx := context.Background()

	txts, err := GetTxt(ctx, "google.com")
	tt.Log(txts, err)

	txts, err = GetTxt(ctx, "google.com", "8.8.8.8")
	tt.Log(txts, err)
}
