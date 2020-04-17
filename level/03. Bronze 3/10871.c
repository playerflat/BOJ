#include <stdio.h>
int main(void){
    int x, y, i, z;
    scanf("%d %d",&x,&y);
    
    for(i=0;i<x;i++){
        scanf("%d",&z);
        if(y>z)
            printf("%d ",z);
    }
    
    return 0;
}