package main 


type roman_data struct {
    num int
    str string
}

var romans []roman_data = []roman_data{
    {1000, "M"},
    {900, "CM"},
    {500,"D"},
    {400,"CD"},
    {100 , "C"},
    {90 , "XC"},
    {50 , "L"},
    {40, "XL"},
    {10 , "X"},
    {9 , "IX"},
    {5 , "V"},
    {4 , "IV"},
    {1, "I"},
}

func intToRoman(num int) string {
    roman_str := ""
    for _, val := range romans {
        if num == 0 {
            break
		}
        multiple := num / val.num
		for range multiple {
			roman_str += val.str
		}
        num  -= (multiple * val.num)
    }
	return roman_str
}