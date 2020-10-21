// 长按键入
// 用双指针
bool isLongPressedName(char * name, char * typed){
    int i=0,j=0;
    if (strlen(name)>strlen(typed)) return false;
    if (name[0]!=typed[0]) return false;
    while(typed[j]!='\0'){
        if (name[i]==typed[j]) {
            i++;
            j++;
            continue;
        }
        if (name[i-1]==typed[j]){
            j++;
            continue;
        }
        return false;
    }
    return strlen(name)==i;
}
