package mylib

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{"正の数の加算", 3, 4, 7},
		{"ゼロを含む加算", 5, 0, 5},
		{"負の数を含む加算", -3, 5, 2},
		{"負の数同士の加算", -3, -4, -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.x, tt.y)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.x, tt.y, result, tt.expected)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		name      string
		x         string
		y         string
		expectedX string
		expectedY string
	}{
		{"通常の文字列", "hello", "world", "world", "hello"},
		{"空文字列を含む", "", "test", "test", ""},
		{"同じ文字列", "same", "same", "same", "same"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultX, resultY := Swap(tt.x, tt.y)
			if resultX != tt.expectedX || resultY != tt.expectedY {
				t.Errorf("Swap(%q, %q) = (%q, %q); want (%q, %q)",
					tt.x, tt.y, resultX, resultY, tt.expectedX, tt.expectedY)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name      string
		sum       int
		expectedX int
		expectedY int
	}{
		{"17を分割", 17, 7, 10},
		{"9を分割", 9, 4, 5},
		{"0を分割", 0, 0, 0},
		{"100を分割", 100, 44, 56},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y := Split(tt.sum)
			if x != tt.expectedX || y != tt.expectedY {
				t.Errorf("Split(%d) = (%d, %d); want (%d, %d)",
					tt.sum, x, y, tt.expectedX, tt.expectedY)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(3, 4)
	}
}

func BenchmarkSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Swap("hello", "world")
	}
}
