#include <stdio.h>

int main(void) {
	char grade[3];
	double score = 0;
	scanf("%s", grade);

		switch (grade[0]) {
		case 'A':
			score += 4;
			break;
		case 'B':
			score += 3;
			break;
		case 'C':
			score += 2;
			break;
		case 'D':
			score += 1;
		default:
			break;
		}

		switch (grade[1]) {
		case '+':
			score += 0.3;
			break;
		case '-':
			score -= 0.3;
			break;
		}
		printf("%.1lf", score);

}