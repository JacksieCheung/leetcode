#include<algorithm>
#include<iostream>
#include<vector>

using namespace std;

void bucketSort(float arr[],int n){
	vector<float> b[n]; // 是用了cpp的容器，其实这里就是一个动态的二维数组

	for (int i=0;i<n;i++){
		int bi = n*arr[i];
		b[bi].push_back(arr[i]);// 插入元素到末尾
	}

	for (int i=0;i<n;i++) sort(b[i].begin(),b[i].end()); // 排序

	int index = 0;
	for (int i=0;i<n;i++){
		for (int j=0;j<b[i].size();j++){
			arr[index++] = b[i][j]; // 连接
		}
	}
}

int main(){
	float arr[] = {0.897,0.565,0.656,0.1234,0.665,0.3434};
	int n = sizeof(arr)/sizeof(arr[0]);
	bucketSort(arr,n);

	cout<<"Sorted arry is \n";
	for (int i=0;i<n;i++) cout<<arr[i]<<" ";
	return 0;
}
