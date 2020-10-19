// 比较含退格（删除）的字符串
// 可以用栈，或者直接在原数组上修改再对比
// 比较快的是逆序双指针

void func(char **S){
    int index = 0;
    while((*S)[index]!='\0'){
        if ((*S)[index]=='#'){
            if (index==0){
                int i = index;
                while((*S)[i]!='\0'){
                    (*S)[i]=(*S)[i+1];
                    i++;
                }
                continue;
            }
            int i = index;
            while((*S)[i]!='\0'){
                (*S)[i-1]=(*S)[i+1];
                i++;
            }
            index--;
            continue;
        }
        index++;
    }
}

bool backspaceCompare(char * S, char * T){
    func(&S);
    func(&T);
    printf("%s %s",S,T);
    if (!strcmp(S,T)) return 1;
    return 0;
}
