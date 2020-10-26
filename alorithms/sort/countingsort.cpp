#include<iostream>

void countingsort(int* arr,int n){
	int C[10]={0}; // k = 10
	int i=0;

	for (i=0;i<n;i++) C[arr[i]]++;

	for (i=1;i<10;i++){
		C[i]=C[i]+C[i-1];
	}

	int Blen = i-1;
	int B[Blen]={0};

	for (i=n-1;i>=0;i--){
		B[C[arr[i]]]=arr[i];
		C[arr[i]]=C[arr[i]]-1;
	}

	i = 0;
	while(i<Blen){
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
