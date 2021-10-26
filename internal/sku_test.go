package internal

import (
	"reflect"
	"testing"
)

func TestNewSKU(t *testing.T) {
	if got := NewSKU(); len(got.Value()) != 9 {
		t.Errorf("NewSKU() = %v, is invalid SKU format", got)
	}
}

func TestNewSKUFromString(t *testing.T) {
	type args struct {
		sku string
	}
	tests := []struct {
		name    string
		args    args
		want    SKU
		wantErr bool
	}{
		{
			name: "parse incorrect SKU",
			args: args{
				sku: "AB66-1234",
			},
			want:    SKU{},
			wantErr: true,
		},
		{
			name: "parse correct SKU",
			args: args{
				sku: "ABCD-1234",
			},
			want: SKU{
				value: "ABCD-1234",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSKUFromString(tt.args.sku)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSKUFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSKUFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
