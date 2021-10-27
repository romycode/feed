package internal

import (
	"errors"
)

type Report struct {
	InvalidCount    int
	ValidCount      int
	DuplicatedCount int
}

type Deduplicator struct {
	invalidCount    int
	validCount      int
	duplicatedCount int
	values          []SKU
}

func NewDeduplicator() *Deduplicator {
	return &Deduplicator{values: []SKU{}}
}

func (d *Deduplicator) Process(data string) error {
	sku, err := NewSKUFromString(data)
	if err != nil {
		if errors.Is(err, ErrInvalidSKU) {
			d.invalidCount++
		}
		return err
	}

	if d.isSaved(sku) {
		d.duplicatedCount++
		return nil
	}

	d.validCount++
	d.values = append(d.values, sku)

	return nil
}

func (d Deduplicator) Report() Report {
	return Report{
		InvalidCount:    d.invalidCount,
		ValidCount:      d.validCount,
		DuplicatedCount: d.duplicatedCount,
	}
}

func (d *Deduplicator) isSaved(sku SKU) bool {
	for _, saved := range d.values {
		if sku.Value() == saved.Value() {
			return true
		}
	}

	return false
}
