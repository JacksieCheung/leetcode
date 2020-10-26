// 字符串压缩
// 要考虑溢出9问题
void partition(char* str,int num,int* index){
    if (num/10!=0){
        partition(str,num/10,index);
    }
    str[(*index)]=num%10+'0';
    (*index)++;
}

char* compressString(char* S){
    if (strlen(S)==0) return S;

	char* res = (char*)malloc(sizeof(char)*500000);
	int slow = 0;
	int fast = 1;
	int count = 1;
	int index = 0;

	while(S[fast]!='\0'){
		if (S[fast]==S[slow]) {
			count++;
		} else {
			res[index]=S[slow];
			index++;
            if (count<9) {
                res[index]=count+'0';
                index++;
            } else {
                partition(res,count,&index);
            }
			count=1;
			slow=fast;
		}
		fast++;
	}

	res[index]=S[slow];
	index++;
    if (index<9){
	    res[index]=count+'0';
        index++;
    } else {
        partition(res,count,&index);
    }
	res[index]='\0';
	
	if (strlen(res)<strlen(S)) return res;
	return S;
}
