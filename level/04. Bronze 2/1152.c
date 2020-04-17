#include <stdio.h>
#include <string.h>

main() {
	char word[1000001];
	int a, cnt=1;

	gets(word);
	a = strlen(word);

	if (word[0] == ' ' || word[0] == '\0')
		cnt--;
	for (int i = 0; i < a; i++) {
		if (word[i] == ' ')
			cnt++;
		if (word[i] == ' ' && word[i + 1]== ' ')
			cnt--;
		if ((word[i] == ' ') && (word[i + 1] == '\0'))
			cnt--;
		}
	printf("%d\n", cnt);
}
