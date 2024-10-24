package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {

	fmt.Println(`_Калькулятор индекса массы тела_`)
	for {
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			/*fmt.Println("Некорректно введены данные")
			continue*/
			panic("Некорректно введены данные")
		}
		outputResult(IMT)
		isUserChoise := askUserChoise()
		if !isUserChoise {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит веса")
	case imt < 18.5:
		fmt.Println("Дефицит массы тела")
	case imt < 25:
		fmt.Println("Норма")
	case imt < 30:
		fmt.Println("Избыточная масса")
	default:
		fmt.Println("Степень ожирения")
	}

}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMETERS_ENTERED")
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах:")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес:")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
func askUserChoise() bool {
	var choise string
	fmt.Println("Вы хотите продолжить?, (Y/n): ")
	fmt.Scan(&choise)
	if choise == "Y" || choise == "y" {
		return true
	}
	return false
}
