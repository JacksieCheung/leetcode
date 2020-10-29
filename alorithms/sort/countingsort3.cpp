#include<iostream>
#include<cstring>
#define RANGE 255 // 宏定义一个范围做输入元素的范围

void countingSort(int* arr,int len){
	int output[len];
	
	int count[RANGE]; // 用来保存计数的数组
	memset(count,0,sizeof(count)); // 此数组初始化为0

	for (int i=0;i<len;i++) count[arr[i]]++; // 统计每个数字出现的次数

	for (int i=1;i<RANGE;i++) count[i]+=count[i-1];

	for (int i=0;i<len;i++){
		output[count[arr[i]]-1]=arr[i];
		count[arr[i]]--;
	}

	for (int i=0;i<len;i++){
		arr[i]=output[i];
	}
}

int main(){
	int arr[]={1,2,5,8,5,4,7,6,3,4};
	int len = sizeof(arr)/sizeof(arr[0]);
	countingSort(arr,len);
	
	for (int i=0;i<len;i++){ 
		std::cout<<arr[i];
	}
	std::cout<<std::endl;
	return 0;
}
