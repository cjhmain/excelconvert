package machine

import (
	"fmt"
	"log"

	"excelconvert/config"
	"excelconvert/util"
	"sync"

	"github.com/xuri/excelize/v2"
)

type Machiner struct {
	index int
}

func New(index int) *Machiner {
	return &Machiner{
		index: index,
	}
}

func (m *Machiner) Run(full_path string, wg *sync.WaitGroup, ch chan *Machiner) {
	defer wg.Done()
	m.read(full_path)
	ch <- m
}

func (m *Machiner) SetIndex(index int) {
	m.index = index
}

func (m *Machiner) GetIndex() int {
	return m.index
}

func (m *Machiner) read(full_path string) {
	fmt.Println("producer read ", full_path)

	file, err := excelize.OpenFile(full_path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for i := 0; i < len(file.GetSheetMap()); i++ {
		sheet_name := file.GetSheetName(i + 1)
		if sheet_name != "" && !util.IsChineseChar(sheet_name) {
			fmt.Println("sheet:", sheet_name)

			rows, err := file.GetRows(sheet_name)
			if err != nil {
				fmt.Println(err)
				return
			}
			for i, row := range rows {
				if i != config.SheetHeadNote {
					for j, col_row := range row {
						if j != config.SheetHeadColumn {
							fmt.Print("columns_", j, ":", col_row, " ")
						}
					}
					fmt.Print("\n")
				}
			}
			fmt.Print("\n")
		}
	}
}
