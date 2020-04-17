#include <stdio.h>

main() {
	int i=0,j,x,y;
	scanf("%d", &j);
	while (i < j) {
		scanf("%d %d", &x, &y);
		printf("%d\n", x + y);
		i++;
	}
}