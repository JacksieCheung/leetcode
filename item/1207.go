// 独一无二的出现次数
// 用两个哈希表就出来了

func uniqueOccurrences(arr []int) bool {
    numsMap:=make(map[int]int);
    for i:=0;i<len(arr);i++{
        numsMap[arr[i]]++;
    }
    // fmt.Println(numsMap);

    tmpMap:=make(map[int]bool)

    for _,v:=range numsMap{
        if tmpMap[v]==true {
            return false;
        }
        tmpMap[v]=true;
    }
    return true;
}
