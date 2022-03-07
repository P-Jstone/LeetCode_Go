package Search

import "fmt"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := make([][]string, 0)
	wordMap := make(map[string]bool, 0)
	// 判断单词表中是否存在单词
	for _, w := range wordList {
		wordMap[w] = true
	}
	if !wordMap[endWord] {
		return res
	}
	delete(wordMap, beginWord)
	delete(wordMap, endWord)
	q1, q2 := []string{beginWord}, []string{endWord}
	next := make(map[string][]string, 0)
	// reversed用于切换遍历的队列，found用于确定是否找到目标元素
	reversed, found := false, false
	for len(q1) > 0 {
		q := []string{}
		q2has := make(map[string]bool, 0)
		for _, v := range q2 {
			q2has[v] = true
		}
		for _, w := range q1 {
			str := []byte(w)
			// 判断q1中元素与哪些元素相邻
			for i := 0; i < len(str); i++ {
				ch := str[i]
				for j := 0; j < 26; j++ {
					str[i] = byte(j + 'a')
					s := string(str)
					// 如果当前元素在q2中出现，说明找到了最短的路径
					if q2has[s] {
						if reversed {
							next[s] = append(next[s], w)
						} else {
							next[w] = append(next[w], s)
						}
						found = true
					}
					// 查找在没有遍历过的单词中是否有相近单词
					if wordMap[s] {
						if reversed {
							next[s] = append(next[s], w)
						} else {
							next[w] = append(next[w], s)
						}
						q = append(q, s)
					}
				}
				str[i] = ch
			}
		}
		if found {
			break
		}
		// 将遍历过的单词从映射中删除
		for _, w := range q {
			delete(wordMap, w)
		}
		// 双向BFS选择较短的队列进行遍历
		if len(q) <= len(q2) {
			q1 = q
		} else {
			reversed = !reversed
			q1 = q2
			q2 = q
		}
	}

	var backtrack func(src, dst string, track []string)
	backtrack = func(src, dst string, track []string) {
		if src == dst {
			res = append(res, append([]string(nil), track...))
			return
		}
		for _, s := range next[src] {
			track = append(track, s)
			backtrack(s, dst, track)
			track = track[:len(track)-1]
		}
	}

	if found {
		track := []string{beginWord}
		backtrack(beginWord, endWord, track)
	}
	return res
}

func main() {
	beginWord, endWord := "magic", "pearl"
	wordList := []string{
		"flail", "halon", "lexus", "joint", "pears", "slabs", "lorie", "lapse", "wroth", "yalow", "swear", "cavil", "piety", "yogis", "dhaka", "laxer", "tatum", "provo", "truss", "tends", "deana", "dried", "hutch", "basho", "flyby", "miler", "fries", "floes", "lingo", "wider", "scary", "marks", "perry", "igloo", "melts", "lanny", "satan", "foamy", "perks", "denim", "plugs", "cloak", "cyril", "women", "issue", "rocky", "marry", "trash", "merry", "topic", "hicks", "dicky", "prado", "casio", "lapel", "diane", "serer", "paige", "parry", "elope", "balds", "dated", "copra", "earth", "marty", "slake", "balms", "daryl", "loves", "civet", "sweat", "daley", "touch", "maria", "dacca", "muggy", "chore", "felix", "ogled", "acids", "terse", "cults", "darla", "snubs", "boats", "recta", "cohan", "purse", "joist", "grosz", "sheri", "steam", "manic", "luisa", "gluts", "spits", "boxer", "abner", "cooke", "scowl", "kenya", "hasps", "roger", "edwin", "black", "terns", "folks", "demur", "dingo", "party", "brian", "numbs", "forgo", "gunny", "waled", "bucks", "titan", "ruffs", "pizza", "ravel", "poole", "suits", "stoic", "segre", "white", "lemur", "belts", "scums", "parks", "gusts", "ozark", "umped", "heard", "lorna", "emile", "orbit", "onset", "cruet", "amiss", "fumed", "gelds", "italy", "rakes", "loxed", "kilts", "mania", "tombs", "gaped", "merge", "molar", "smith", "tangs", "misty", "wefts", "yawns", "smile", "scuff", "width", "paris", "coded", "sodom", "shits", "benny", "pudgy", "mayer", "peary", "curve", "tulsa", "ramos", "thick", "dogie", "gourd", "strop", "ahmad", "clove", "tract", "calyx", "maris", "wants", "lipid", "pearl", "maybe", "banjo", "south", "blend", "diana", "lanai", "waged", "shari", "magic", "duchy", "decca", "wried", "maine", "nutty", "turns", "satyr", "holds", "finks", "twits", "peaks", "teems", "peace", "melon", "czars", "robby", "tabby", "shove", "minty", "marta", "dregs", "lacks", "casts", "aruba", "stall", "nurse", "jewry", "knuth",
	}
	res := findLadders(beginWord, endWord, wordList)
	for _, r := range res {
		fmt.Println(r)
	}
}
