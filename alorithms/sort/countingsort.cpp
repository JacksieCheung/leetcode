#include<iostream>

void countingsort(int* arr,int n){
	int C[10]={0}; // k = 10
	int i=0;

	for (i=0;i<n;i++) C[arr[i]]++;

	for (i=1;i<10;i++){
		C[i]=C[i]+C[i-1];
	}

	int B[n]={0};

	for (i=n-1;i>=0;i--){
		B[C[arr[i]]-1]=arr[i];   // C数组保存的就是我们要的次序，用arr[i]访问得到
		C[arr[i]]=C[arr[i]]-1; // 特定的字符的次序。然后该次序减1,如果有重复的元
	}                          // 素，给他们留出位置。所以C保存的同一个字符最后
	                           // 出现的位置

	i = 0;
	while(i<n){
		std::cout<<B[i]<<" ";
		i++;
	}
	std::cout<<std::endl;
}

int main(){
	int arr[]={5,2,1,4,8,6,5,2};
	int n = sizeof(arr)/sizeof(arr[0]);
	countingsort(arr,n);
	return 0;
}
