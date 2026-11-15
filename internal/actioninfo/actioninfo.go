package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает обработку данных об активности.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info выводит информацию по каждой корректной записи в dataset.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка разбора данных %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка обработки данных %q: %v", data, err)
			continue
		}

		fmt.Print(info)
	}
}
