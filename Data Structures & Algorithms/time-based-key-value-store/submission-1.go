type val struct{
    mood string
    tim int
}
type TimeMap struct {
  timMap map[string][]val
}

func Constructor() TimeMap {
    tm := make(map[string][]val)

    return TimeMap{tm}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    this.timMap[key] = append(this.timMap[key], val{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
    res := ""
    if val, exists := this.timMap[key]; exists{

        left, right := 0, len(val)-1
        
        
        for left <= right{
            mid := left + (right - left)/2

            if val[mid].tim<= timestamp{
                res = val[mid].mood
                left = mid + 1
            }else{
                right = mid-1
            }
        }
    }

    return res
}

func max(a, b int) int{
    if a< b{
        return b
    }

    return a
}
