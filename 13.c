// 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
// 给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
int romanToInt(char * s){
    int a[20]={0};
    int i=0,k,sum=0;
    while(*s!='\0'){
    switch(*s){
        case 'M':a[i]=1000;break;
        case 'D':a[i]=500;break;
        case 'C':a[i]=100;break;
        case 'L':a[i]=50;break;
        case 'X':a[i]=10;break;
        case 'V':a[i]=5;break;
        case 'I':a[i]=1;break;
    }
        i++;
        s++;
    }
    for(k=0;k<i;k++){
        if(a[k]<a[k+1])
            sum-=a[k];
        if(a[k]>=a[k+1])
            sum+=a[k];
    }
    return sum;
}

