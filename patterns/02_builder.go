package patterns

import (
	"fmt"
	"time"
)

const (
	AsusCollectorType = "Asus"
	HpCollectorType   = "HP"
)

type collector interface {
	SetCore()
	SetMemory()
	SetGraphicCard()
	SetMonitor()
	SetBrand()
	GetComputer() computer
}

type computer struct {
	Core        int
	Memory      int
	GraphicCard int
	Monitor     int
	Brand       string
}

func (c computer) Print() {
	fmt.Printf("%s Core: [%d], Memory: [%d], GraphicCard: [%d], Monitor: [%d]\n", c.Brand, c.Core, c.Memory, c.GraphicCard, c.Monitor)
}

func getCollector(collectorType string) collector {
	switch collectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HpCollectorType:
		return &HPCollector{}
	}
}

type AsusCollector struct {
	Core        int
	Memory      int
	GraphicCard int
	Monitor     int
	Brand       string
}

func (as *AsusCollector) SetCore() {
	as.Core = 4
}

func (as *AsusCollector) SetMemory() {
	as.Memory = 256
}

func (as *AsusCollector) SetGraphicCard() {
	as.GraphicCard = 1
}

func (as *AsusCollector) SetMonitor() {
	as.Monitor = 1
}

func (as *AsusCollector) SetBrand() {
	as.Brand = AsusCollectorType
}

func (as *AsusCollector) GetComputer() computer {
	return computer{
		Core:        as.Core,
		Memory:      as.Memory,
		GraphicCard: as.GraphicCard,
		Monitor:     as.Monitor,
		Brand:       as.Brand,
	}
}

type HPCollector struct {
	Core        int
	Memory      int
	GraphicCard int
	Monitor     int
	Brand       string
}

func (hp *HPCollector) SetCore() {
	hp.Core = 8
}

func (hp *HPCollector) SetMemory() {
	hp.Memory = 512
}

func (hp *HPCollector) SetGraphicCard() {
	hp.GraphicCard = 2
}

func (hp *HPCollector) SetMonitor() {
	hp.Monitor = 2
}

func (hp *HPCollector) SetBrand() {
	hp.Brand = HpCollectorType
}

func (hp *HPCollector) GetComputer() computer {
	return computer{
		Core:        hp.Core,
		Memory:      hp.Memory,
		GraphicCard: hp.GraphicCard,
		Monitor:     hp.Monitor,
		Brand:       hp.Brand,
	}
}

type Director struct {
	collector collector
}

func newDirector(cl collector) *Director {
	return &Director{collector: cl}
}

func (fc *Director) setCollector(collector collector) {
	fc.collector = collector
}

func (fc *Director) createcomputer() computer {
	fc.collector.SetCore()
	fc.collector.SetMemory()
	fc.collector.SetGraphicCard()
	fc.collector.SetMonitor()
	fc.collector.SetBrand()
	return fc.collector.GetComputer()
}

func BuilderPattern() {
	asusCollector := getCollector(AsusCollectorType)
	hpCollector := getCollector(HpCollectorType)

	director := newDirector(asusCollector)
	asusComputer := director.createcomputer()
	asusComputer.Print()

	time.Sleep(time.Second * 2)

	director.setCollector(hpCollector)
	hpComputer := director.createcomputer()
	hpComputer.Print()

	fmt.Println("Строитель реализован.")
}
