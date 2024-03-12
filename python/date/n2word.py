
def int_to_word(num: int):
    unit1 = ["", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"]
    unit2 = ["", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"]
    
    ret = ''
    if num < 20:
        ret = unit1[num]
    elif num < 100:
        ret += str(unit2[num // 10]) + " " + str(unit1[num % 10])
        

    return ret
    
print(int_to_word(30));
