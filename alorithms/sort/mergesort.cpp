#include<iostream>

void merge(int* arr,int low,int middle,int high){
	int i,j,k;
	int n1 = middle-low+1;
	int n2 = high-middle; // 这里把原有部分分成两个部分，其中middle算前半部分。
	int a[n1];
	int b[n2];
	// 对辅助数组赋值
	for (i=0;i<n1;i++) a[i]=arr[low+i];
	for (j=0;j<n2;j++) b[j]=arr[middle+1+j];

	i=0,j=0,k=low;    // 记得要对三个索引初始化
	// merge 对两个已经有序的数组采用归并排序，各走一遍
	while(i<n1&&j<n2){
		if (a[i]<=b[j]) {
			arr[k]=a[i];
			i++;
		} else {
			arr[k]=b[j];
			j++;
		}
		k++;
	}
	while(i<n1) {
		arr[k]=a[i];
		k++;
		i++;
	}
	while(j<n2){
		arr[k]=b[j];
		k++;
		j++;
	}
}

// 调用 mergesort 对数组进行拆分，拆到最后底层为一个，开始回溯
// 然后一个合并成两个，两个的片段合并成四个元素片段，以此回溯
void mergesort(int* arr,int low,int high){
	if (low<high){
		int middle=(low+high)/2;
		mergesort(arr,low,middle);
		mergesort(arr,middle+1,high);
		merge(arr,low,middle,high);
	}
}

int main(){
	int arr[]={9,8,4,5,2,1,4,5,8,7,6};
	int n = sizeof(arr)/sizeof(arr[0]);
	mergesort(arr,0,n-1);
	for (int i=0;i<n;i++){
		std::cout<<arr[i]<<" ";
	}
	std::cout<<std::endl;
	return 0;
}
