#include<iostream>

void heapify(int* arr,int n,int i){
	int largest = i;
	int l = 2*i + 1;
	int r = 2*i + 2;

	if (l<n&&arr[l]>arr[largest])
		largest = l;

	if (r<n&&arr[r]>arr[largest])
		largest = r;

	if (largest != i){
		// swap(arr[i],arr[largest]);
		int tmp = arr[i];
		arr[i] = arr[largest];
		arr[largest] = tmp;
		
		heapify(arr,n,largest);
	}
}

int main(){
	int arr[]={5,1,4,6,8,3 };
	int n = sizeof(arr)/sizeof(arr[0]);
	for (int i = n/2-1;i>=0;i--)
		heapify(arr,n,i);
	
	for (int i=0;i<n;i++)
		std::cout<<arr[i]<<" ";
	std::cout<<std::endl;
	return 0;
}
