def int_to_word(num: int):
    unit1 = ["", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"]
    unit2 = ["", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"]
    unit3 = [
        "Trillion", "Billion", "Million", "Thousand", "hundred"
    ]
    unit3_val = [1000000000000, 1000000000, 1000000, 1000, 100]
    
    ret = ''
    
    if num < 20:
        ret = unit1[num]
    elif num < 100:
        ret += str(unit2[num // 10]) + " " + str(unit1[num % 10])
    else:
        i = 0 
        unit_len = len(unit3)
        while i < unit_len:
            d = num // unit3_val[i]
            r = num % unit3_val[i]
            if d > 0:
                ret += str(int_to_word(d)) + " " + str(unit3[i]) + " " + str(int_to_word(r))
                break
            i += 1

    return ret
    
print(int_to_word(2000123));
