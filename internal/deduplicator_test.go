package internal

import (
	"reflect"
	"testing"
)

func TestDeduplicator_Process(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    Report
		wantErr bool
	}{
		{
			name: "save five valid sku",
			args: []string{
				"AAaa-0011\n",
				"AAab-0011\n",
				"AAac-0011\n",
				"AAad-0011\n",
				"AAaf-0011\n",
			},
			wantErr: false,
			want: Report{
				invalidCount:    0,
				validCount:      5,
				duplicatedCount: 0,
			},
		},
		{
			name: "don't save duplicates",
			args: []string{
				"AAaa-0011\n",
				"AAaa-0011\n",
				"AAaa-0011\n",
				"AAaa-0011\n",
				"AAaf-0011\n",
				"AAaf-0011\n",
				"AAaf-0011\n",
			},
			wantErr: false,
			want: Report{
				invalidCount:    0,
				validCount:      2,
				duplicatedCount: 5,
			},
		},
		{
			name: "process invalid sku without '\\n'",
			args: []string{
				"AAaf-0011",
			},
			wantErr: true,
			want: Report{
				invalidCount:    1,
				validCount:      0,
				duplicatedCount: 0,
			},
		},
		{
			name: "process invalid sku with digits",
			args: []string{
				"AA00-0011",
			},
			wantErr: true,
			want: Report{
				invalidCount:    1,
				validCount:      0,
				duplicatedCount: 0,
			},
		},
		{
			name: "process invalid sku with letters",
			args: []string{
				"AAAA-AAAA",
			},
			wantErr: true,
			want: Report{
				invalidCount:    1,
				validCount:      0,
				duplicatedCount: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDeduplicator()
			for _, arg := range tt.args {
				if err := d.Process(arg); (err != nil) != tt.wantErr {
					t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if got := d.Report(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process don't work as expected, got %v, want %v", got, tt.want)
			}
		})
	}
}
