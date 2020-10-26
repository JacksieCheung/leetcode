#include<stdio.h>

int main(){
	long int nums[10]={1,100,10000,1000000};
	printf("%d\n",*nums);
	printf("%d",*++nums);

	return 0;
}
