#include <stdio.h>

main() {
	int i, n, sum = 0;
	char num[100];
	scanf("%d", &n);
	scanf("%s", num);
	
	for (i = 0; i < n; i++) {
		num[i] -= 48;
		if (num[i] > 0)
			sum += num[i];
	}
	
	printf("%d", sum);
}