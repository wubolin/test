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

func TestMulti(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		x, y int
		want int
	}{
		{name: "positive", x: 3, y: 4, want: 12},
		{name: "zero", x: 0, y: 9, want: 0},
		{name: "negative", x: -2, y: 5, want: -10},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Multi(tt.x, tt.y); got != tt.want {
				t.Fatalf("Multi(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
