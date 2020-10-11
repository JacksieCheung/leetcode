// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串""
char * longestCommonPrefix(char ** strs, int strsSize){
    if(strsSize==0)
        return "";
    int len=strlen(*strs);
    int k=0,i=0,n=0;
    int l=len;
    char*p=(char*)malloc(128);
    memset(p,0,len);
    p=*(strs+n);
    for(i=0;i<strsSize;i++){
        for(k=0;k<strlen(*(strs+i));k++){
            if(p[k]!=strs[i][k])
                   break;
        }
        if(k<len)
            len=k;
    }
     for(i=len;i<l;i++){
         p[i]='\0';
     }   
    return p;
}
