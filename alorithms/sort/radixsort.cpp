#include<iostream>

int getMax(int* arr,int len){
	int max = arr[0];
	for (int i=1;i<len;i++){
		if (max<arr[i]) max=arr[i];
	}
	return max;
}

void countingsort(int* arr,int n,int exp){
	int count[10]={0}; // k = 10
	int i=0;

	for (i=0;i<n;i++) count[(arr[i]/exp)%10]++;

	for (i=1;i<10;i++){
		count[i]=count[i]+count[i-1];
	}

	int output[n]={0};

	for (i=n-1;i>=0;i--){
		output[count[(arr[i]/exp)%10]-1]=arr[i];   // count数组保存的就是我们要的次序，用arr[i]访问得到
		count[(arr[i]/exp)%10]=count[(arr[i]/exp)%10]-1; // 特定的字符的次序。然后该次序减1,如果有重复的元
	}                          // 素，给他们留出位置。所以count保存的同一个字符最后
	                           // 出现的位置
	for (i=0;i<n;i++) arr[i] = output[i]; // output直接存了完整的数。然后再赋值来排序
}

void radixsort(int* arr,int len){
	int m=getMax(arr,len);
	for (int exp = 1;m/exp>0;exp*=10) countingsort(arr,len,exp);
}

int main(){
	int arr[]={9696,176,254,7412,4568,4523,84555};
	int len = sizeof(arr)/sizeof(arr[0]);
	radixsort(arr,len);
	for (int i=0;i<len;i++) std::cout<<arr[i]<<" ";
	std::cout<<std::endl;
	return 0;
}
