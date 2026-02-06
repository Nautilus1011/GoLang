package mylib

import "fmt"

type Animal interface {
	Run(int) string
	Cry() string
}

type Cat struct {
	Name string
}

type Human struct {
	Name string
	Job  string
}

func (c *Cat) Cry() string {
	return fmt.Sprintf("%sちゃんはお腹が減っているので泣きます。", c.Name)
}

func (h *Human) Cry() string {
	return fmt.Sprintf("%sさんは%sの仕事が辛いので泣きます。", h.Name, h.Job)
}

func (c *Cat) Run(speed int) string {
	return fmt.Sprintf("%sちゃんは獲物を追う時に時速%dkmで走ります。", c.Name, speed)
}

func (h *Human) Run(speed int) string {
	return fmt.Sprintf("%sさんはマラソンで時速%dkmで走ります。", h.Name, speed)
}

func AnimalCry() {
	cat := &Cat{Name: "クー"}
	human := &Human{Name: "田中", Job: "ソフトウェアエンジニア"}

	fmt.Println(cat.Cry())
	fmt.Println(human.Cry())
	fmt.Println(cat.Run(30))
	fmt.Println(human.Run(10))
}
