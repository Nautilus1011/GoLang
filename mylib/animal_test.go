package mylib

import (
	"strings"
	"testing"
)

func TestCat_Cry(t *testing.T) {
	cat := &Cat{Name: "クー"}
	result := cat.Cry()

	if !strings.Contains(result, "クー") {
		t.Errorf("Cat.Cry() の結果に猫の名前が含まれていません: %s", result)
	}

	if !strings.Contains(result, "泣きます") {
		t.Errorf("Cat.Cry() の結果に「泣きます」が含まれていません: %s", result)
	}
}

func TestCat_Run(t *testing.T) {
	cat := &Cat{Name: "クー"}
	speed := 30
	result := cat.Run(speed)

	if !strings.Contains(result, "クー") {
		t.Errorf("Cat.Run() の結果に猫の名前が含まれていません: %s", result)
	}

	if !strings.Contains(result, "30km") {
		t.Errorf("Cat.Run() の結果に速度が含まれていません: %s", result)
	}
}

func TestHuman_Cry(t *testing.T) {
	human := &Human{Name: "田中", Job: "ソフトウェアエンジニア"}
	result := human.Cry()

	if !strings.Contains(result, "田中") {
		t.Errorf("Human.Cry() の結果に名前が含まれていません: %s", result)
	}

	if !strings.Contains(result, "ソフトウェアエンジニア") {
		t.Errorf("Human.Cry() の結果に職業が含まれていません: %s", result)
	}
}

func TestHuman_Run(t *testing.T) {
	human := &Human{Name: "田中", Job: "ソフトウェアエンジニア"}
	speed := 10
	result := human.Run(speed)

	if !strings.Contains(result, "田中") {
		t.Errorf("Human.Run() の結果に名前が含まれていません: %s", result)
	}

	if !strings.Contains(result, "10km") {
		t.Errorf("Human.Run() の結果に速度が含まれていません: %s", result)
	}
}

func TestAnimalInterface(t *testing.T) {
	tests := []struct {
		name   string
		animal Animal
	}{
		{"CatはAnimalインターフェースを実装", &Cat{Name: "テスト猫"}},
		{"HumanはAnimalインターフェースを実装", &Human{Name: "テスト太郎", Job: "テスター"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cryResult := tt.animal.Cry()
			runResult := tt.animal.Run(20)

			if cryResult == "" {
				t.Error("Cry() が空文字列を返しました")
			}

			if runResult == "" {
				t.Error("Run() が空文字列を返しました")
			}
		})
	}
}
