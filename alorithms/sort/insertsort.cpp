#include<iostream>

void insertsort(int* arr,int len){
	for (int j=1;j<len;j++) {
		int key = arr[j];
		int i = j-1;
		while(i>=0&&arr[i]>key) {
			arr[i+1]=arr[i];
			i--;
		}
		arr[i+1] = key;
	}
}

int main(){
	int arr[] = {5,4,1,2,5,4,8,9,7,10,4};
	int len = sizeof(arr)/sizeof(arr[0]);
	insertsort(arr,len);
	for (int i=0;i<len;i++){
		std::cout<<arr[i]<<" ";
	}
	std::cout<<std::endl;
	return 0;
}
