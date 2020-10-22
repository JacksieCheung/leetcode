#include<iostream>

// 交换函数
void swap(int* m,int* n){
	int tmp = *m;
	*m = *n;
	*n = tmp;
}

// 默认以最后一个元素作为中枢
// 比中枢大的放右边，比中枢小的放左边。一直递归求解即可。
int partition(int* arr,int low,int high){
	int pivot = arr[high];
	int i = low - 1;

	// 采用了双指针比较，一遍循环结束
	for (int j=low;j<=high-1;j++){
		if (arr[j]<pivot) {
			i++;
			swap(&arr[j],&arr[i]);
		}
	}
	swap(&arr[i+1],&arr[high]);
	return (i+1);
}

void quicksort(int* arr,int low,int high){
	if (low<high){
		int pi = partition(arr,low,high);

		quicksort(arr,low,pi-1);
		quicksort(arr,pi+1,high);
	}
}

int main(){
	int a[]={5,4,1,4,4,8,6,8,4,2,6,5};
	int n = sizeof(a)/sizeof(a[0]);
	quicksort(a,0,n-1);
	
	for (int i=0;i<n;i++){
		std::cout<<a[i]<<" ";
	}
	std::cout<<std::endl;
	return 0;
}
