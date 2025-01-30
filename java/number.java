public class number {
	public static void main(String []args) {
		String []ones = {
						"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
						"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen",
						"eighteen", "nineteen"
		};

		String []tens = { "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"};

		int []unit_int = {1000000};
		int []unit_int_d = {1000};

		String []unit_str = {"Thousand"};

		int input = 20000;

		String rwords = "";

		int index = 0;

		while (true) {
			if (input <= 0) {
				break;
			}
			else if (input < 20) {
				rwords += ones[input - 1];
				break;
			} else if (input < 100) {
				rwords += tens[(input / 10) - 2] + " " + ones[(input % 10) - 1];
				break;
			} else if (input < 1000) {
				rwords += ones[(input / 100) - 1] + " hundred ";
				input %= 100;
			} else if (input < unit_int[index]) {
				int a = input / unit_int_d[index];
				input = input % unit_int_d[index];
				break;
			}
			index++;
		}

		System.out.println(rwords);

	}
}
