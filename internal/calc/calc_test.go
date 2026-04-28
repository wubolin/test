package calc

import "testing"

func TestSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		nums    []int
		want    int
		wantErr bool
	}{
		{name: "normal", nums: []int{1, 2, 3}, want: 6},
		{name: "with negative", nums: []int{5, -2, 1}, want: 4},
		{name: "empty", nums: []int{}, wantErr: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Sum(tt.nums...)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Sum() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("Sum() got = %d, want %d", got, tt.want)
			}
		})
	}
}
