package leetcode


type kv struct {
	Key   string
	Value int
}

func TopKFrequent692(words []string, k int) []string {
	freqMap := make(map[string]int)
	for _, str := range words {
		freqMap[str] += 1
	}

	kvList1 := make([]kv, 0)
	insertedMap := make(map[string]int)
	for _, str := range words {
		if _,ok:=insertedMap[str];!ok{
			insertedMap[str]=1
			kvList1 = append(kvList1, kv{str, freqMap[str]})
		}
	}
	//	sort.Slice(kvList,func(i,j int)bool{return kvList[i].Key<kvList[j].Key})
	//TODO 未做完
	result := make([]string, 0)
	for i := 0; i < k; i++ {
		result = append(result, kvList1[i].Key)
	}
	return result
}
