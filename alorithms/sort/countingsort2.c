// C Program for counting sort 
#include <stdio.h> 
#include <string.h> 
#define RANGE 255 // 宏定义
  
void countSort(char arr[]) 
{
   // output 是输出数组，长度为传入的数组长度	
    char output[strlen(arr)]; 
  
    int count[RANGE + 1], i; 
    memset(count, 0, sizeof(count)); // 将 count 初始化为 0
  
    for (i = 0; arr[i]; ++i) 
        ++count[arr[i]-'0']; // 这里 arr[i] 可以转换成 int，但是会 warning，减去 '0' 用 ASCII码 转换即可。
  
    for (i = 1; i <= RANGE; ++i) // 遍历 count 数组的长度，让前面的数累加起来。
        count[i] += count[i - 1]; 
  
    for (i = 0; arr[i]; ++i) { // 将结果写入的过程
        output[count[arr[i]-'0'] - 1] = arr[i]; 
        --count[arr[i]-'0']; 
    } 
  
    for (i = 0; arr[i]; ++i) 
        arr[i] = output[i]; 
} 
  
int main() 
{ 
    char arr[] = "geeksforgeeks"; //"applepp"; 
  
    countSort(arr); 
  
    printf("Sorted character array is %s\n", arr); 
    return 0; 
} 
