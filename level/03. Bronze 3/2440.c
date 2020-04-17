#include <stdio.h>
int main(void){
    int N,i,j;
    int k=0;
    scanf("%d",&N);
    
    for(i=0;i<N;i++){
        for(j=0;j<N-k;j++){
            printf("*");    
        }
        printf("\n");
        k+=1;
    }
}