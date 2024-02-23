package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ArrayRoman = [101]string{ //Массив римских чисел от 1 до 100, где индекс ячейки массива соответствует римскому числу
	"0",
	"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
	"XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
	"LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
	"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

func ClearSpace(Clean string) string { //Функция очистки строки от пробелов
	Clean = strings.ReplaceAll(Clean, " ", "")
	Clean = strings.ReplaceAll(Clean, " ", "")
	Clean = strings.ReplaceAll(Clean, "\n", "")
	Clean = strings.ReplaceAll(Clean, "\t", "")
	Clean = strings.ReplaceAll(Clean, "\r", "")
	Clean = strings.ReplaceAll(Clean, "\v", "")
	TextClean := Clean
	return TextClean
}

func FindRoman(Numbers []string) int { // Функция поиска римских чисел, возвращает количество найденных чисел
	var FindNumber int = 0 //Переменная количества найденных римских чисел в строке
	for i := 0; i < len(Numbers)-1; i++ {
		for a := 1; a < len(ArrayRoman)-90; a++ {
			if ArrayRoman[a] == Numbers[i] {
				FindNumber++
			}
		}
	}
	//fmt.Print(FindNumber, " :количество римских цифр\n")
	if FindNumber < 2 && FindNumber != 0 { //Условие, при котором выдает панику, если римские числа найдены, но их меньше двух
		panic("[7: Лишь один оператор является корректным римским числом.] \n Выражение не является корректной математической операцией.")
	}
	return FindNumber
}

func FindOper(AllFind string) { //Функция поиска операторов и проверки их расположения
	var ArrayOper = [4]string{"+", "-", "*", "/"} //Массив операторов, которые ищем
	//var DefaultStrings = AllFind                  //Копирование принимаемой строки
	var CountOper = 0 //Переменная найденных разных операторов в строке
	for i := 0; i < len(ArrayOper); i++ {
		for a := 0; a < len(ArrayOper); a++ { //Цикл поиска и разделение строки по операторам
			if strings.Contains(AllFind, ArrayOper[a]) {
				AllFind = strings.Replace(AllFind, ArrayOper[a], "", 1)
				CountOper++
				if CountOper > 1 { //Условие, при котором выдает панику, если было найдно более одного оператора
					panic("[3: Операторов в строке больше одного] \n Строка не является корректным выражением.")
				}
			}
		}
		if strings.HasPrefix(AllFind, ArrayOper[i]) { //Условие, при котором выдает панику, если оператор стоит в начале строки
			panic("[4: Оператор стоит в начале строки.] \n Строка не является математической операцией.")
		}
		if strings.HasSuffix(AllFind, ArrayOper[i]) { //Условие, при котором выдает панику, если оператор стоит в конце строки
			panic("[5: Оператор стоит в конце строки.] \n Строка не является математической операцией.")
		}
	}
	if CountOper == 0 { //Условие паники, если операторов не найдено
		panic("[6: Оператор не найден.] \n Выражение не является математической операцией.")
	}

}

func CharactersDefinition(Expression string) []string { //Функция определения оператора в строке и преобразования ее в срез. Возвращает срез с найденным оператором в конце
	var Expression1 []string
	var Operator string

	Expression1 = strings.Split(Expression, "+")
	Operator = "+"

	if len(Expression1) < 2 {
		Expression1 = strings.Split(Expression, "-")
		Operator = "-"
	}

	if len(Expression1) < 2 {
		Expression1 = strings.Split(Expression, "/")
		Operator = "/"
	}

	if len(Expression1) < 2 {
		Expression1 = strings.Split(Expression, "*")
		Operator = "*"
	}
	Expression1 = append(Expression1, Operator)
	return Expression1
}

func main() {

	for {
		Reader := bufio.NewReader(os.Stdin)
		fmt.Print("Введите выражение: ")
		InputText, _ := Reader.ReadString('\n')      //Считывание выражения из ввода в консоли
		var TextClean string = ClearSpace(InputText) //Очистка символов пробелов из строки
		if len(TextClean) < 3 {                      //Условие паники, если выражение слишком короткое
			panic("[1: Строка короче допустимой.] Строка не является математической операцией.")
		}

		if len(TextClean) > 9 { //Условие паники, если выражение слишком длинное
			panic("[2: Строка длиннее допустимой.] \n Строка не является корректным выражением.")
		}

		FindOper(TextClean)                           //Проверка операторов
		Expression := CharactersDefinition(TextClean) //Определение оператора и преобразование строки в срез
		CountRoman := FindRoman(Expression)           //Поиск римских чисел

		var TheOperand = []int{} //Создание среза для операндов выражения
		if CountRoman == 2 {     //Условие, при котором математическая операция проводится с римскими цифрами
			for i := 0; i < len(Expression)-1; i++ {
				for a := 1; a < len(ArrayRoman)-90; a++ { //Поиск и запись римских чисел в срез операндов выражения
					if ArrayRoman[a] == Expression[i] {
						TheOperand = append(TheOperand, a)
					}
				}
			}
			switch Expression[len(Expression)-1] { //Цикл поиска необходимой операции с операндами и ее выполнения
			case "+":
				Sum := TheOperand[0] + TheOperand[1]
				fmt.Print("Результат = ", ArrayRoman[Sum], "\n")
			case "-":
				Sum := TheOperand[0] - TheOperand[1]
				if Sum <= 0 {
					panic("[8: Результат меньше или равен нулю.] \n Выражение не является корректным.")
				}
				fmt.Print("Результат = ", ArrayRoman[Sum], "\n")
			case "*":
				Sum := TheOperand[0] * TheOperand[1]
				fmt.Print("Результат = ", ArrayRoman[Sum], "\n")
			case "/":
				Sum := TheOperand[0] / TheOperand[1]
				if Sum <= 0 {
					panic("[9: Результат меньше или равен нулю.] \n Выражение не является корректным.")
				}
				fmt.Print("Результат = ", ArrayRoman[Sum], "\n")
			}
		} else { //Если выражение не одержит двух римских чисел, то вычисления производятся с арабскими числами

			for i := 0; i < len(Expression)-1; i++ { //Цикл преобразования элементов строкового среза в срез целочисленный
				ArabicNumbers, err := strconv.Atoi(Expression[i])
				if err != nil { //Если элемент в строковом срезе нельзя преобразовать в целочисленный - вызов паники
					panic("[10: Одно или два операнда не являются целочисленными и оба не являются римскими числами не больше X.] \n Строка не является корректным выражением.")
				}
				if ArabicNumbers > 10 { //Проверка не является ли операнд числом больше 10
					panic("11: Операнд не может быть больше числа 10] \n Выражение не является корректным.")
				}
				if ArabicNumbers == 0 {
					panic("12: Операнд не может быть равен нулю.] \n Выражение не является корректным.")
				}
				TheOperand = append(TheOperand, ArabicNumbers)
			}
			switch Expression[len(Expression)-1] { //Цикл поиска необходимой операции с операндами и ее выполнения
			case "+":
				Sum := TheOperand[0] + TheOperand[1]
				fmt.Print("Результат = ", Sum, "\n")
			case "-":
				Sum := TheOperand[0] - TheOperand[1]
				fmt.Print("Результат = ", Sum, "\n")
			case "*":
				Sum := TheOperand[0] * TheOperand[1]
				fmt.Print("Результат = ", Sum, "\n")
			case "/":
				Sum := TheOperand[0] / TheOperand[1]
				fmt.Print("Результат = ", Sum, "\n")
			}
		}
	}
}
