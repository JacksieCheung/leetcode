#include<iostream>

void swap(int* a,int* b){
	int tmp = *a;
	*a = *b;
	*b = tmp;
}

void bubblesort(int* arr,int n){
	for (int i=0;i<n;i++){
		for(int j=n;j>=i+1;j--){
			if (arr[j]<arr[j-1])
				swap(&arr[j],&arr[j-1]);
		}
	}
}

int main(){
	int arr[]={1,4,3,5,9,6,2,5,4,7,8};
	int n=sizeof(arr)/sizeof(arr[0]);
	
	bubblesort(arr,n-1);

	for(int i=0;i<n;i++){
		std::cout<<arr[i]<<" ";
	}
	std::cout<<std::endl;
}
