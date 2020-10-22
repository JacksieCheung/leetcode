/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
// 划分字母区间，贪心算法+双指针
// 遍历一遍取字母最后出现的索引
// 遍历字符串，取end为当前片段的结尾索引，遇到的字符取其最大最后索引，当end等于当前索引时，说明到头了。
int* partitionLabels(char * S, int* returnSize){
    int last[26];
    int sLen = strlen(S);
    for (int i = 0;i<sLen;i++){
        last[S[i]-'a'] = i;
    }

    int end = 0,start = 0;
    *returnSize = 0;
    int *res = (int*)malloc(sizeof(int)*sLen);
    for (int i = 0;i<sLen;i++){
        end = max(last[S[i]-'a'],end);
        if (end == i){
            res[*returnSize] = end-start+1;
            (*returnSize)++;
            start = end+1;
        }
    }
    return res;
}

int max(int i,int j){
    if (i>j) return i;
    return j;
}
