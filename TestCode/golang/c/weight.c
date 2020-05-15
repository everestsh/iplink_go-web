#include <stdio.h>
//现在有1克到127克之间127个整数重量，
/*
    使用循环解决砝码问题的练习
*/
int main(){
	int weight = 0, last = 0;
	printf("需要的砝码:\n");
	for(weight = 1; weight <= 100; weight++){
		if (weight >= 2 * last){
			printf("%d ", weight);
			last = weight;
		}
	}
	printf("\n");
}
