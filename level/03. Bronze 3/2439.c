#include <stdio.h>
int main(void){
    int N,i,j,k;
    scanf("%d",&N);
    
    for(i=1;i<=N;i++){
        for(k=0;k<N-i;k++){
            printf(" ");
        }
        for(j=0;j<i;j++){
            printf("*");    
        }
        printf("\n");
    }
}