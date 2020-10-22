// 比较含退格的字符串
// 用双指针

bool backspaceCompare(char * S, char * T){
    int i = strlen(S)-1;
    int j = strlen(T)-1;

    int skipi = 0;
    int skipj = 0;

    while(i>=0||j>=0){
        while (i>=0){
            if (S[i]=='#'){
                skipi++;
                i--;
            } else if (skipi>0) {
                skipi--;
                i--;
            } else {
                break;
            }
        }

        while (j>=0){
            if (T[j]=='#'){
                skipj++;
                j--;
            } else if (skipj>0){
                skipj--;
                j--;
            } else {
                break;
            }
        }

        if (i>=0&&j>=0){
            if (S[i]!=T[j]) return false;
            i--;
            j--;
        } else {
            if (i>=0||j>=0){
                return false;
            }
        }
    }
    return true;
}
