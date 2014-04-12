package godatamining

type ZeroR struct {
	Data *[][]string
}

func (z *ZeroR) getKeyLoc(key string) (int, bool) {
	keys := (*z.Data)[0]
	for i := 0; i < len(keys); i++ {
		if keys[i] == key {
			return i, true
		}
	}
	return -1, false
}

func (z *ZeroR) GetRules(successKey string) (string, bool) {
	if loc, ok := z.getKeyLoc(successKey); ok {

		results := make(map[string]int)

		for i := 1; i < len(*z.Data); i++ {
			results[(*z.Data)[i][loc]] += 1
		}

		var out string
		for a, b := range results {
			if len(out) == 0 {
				out = a
			} else {
				if results[out] < b {
					out = a
				}
			}
		}
		return out, true
	}
	return "", false
}

func (z *ZeroR) GetErrorRate(key string) (*Result, bool) {
	if rule, ok := z.GetRules(key); ok {
		loc, _ := z.getKeyLoc(key)
		n := len(*z.Data)
		mistakes := 0
		for i := 1; i < n; i++ {
			if (*z.Data)[i][loc] != rule {
				mistakes += 1
			}

		}
		output := float32(mistakes) / float32(n-1)
		out := &Result{key, output, &map[string]float32{rule: output}}
		return out, true
	}

	return new(Result), false
}

func (z *ZeroR) GetResult(successKey string) (*Result, bool) {
	return z.GetErrorRate(successKey)

}
