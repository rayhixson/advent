package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	ids := []string{}
	totalTwos, totalThrees := 0, 0
	scanner := bufio.NewScanner(strings.NewReader(boxids))
	for scanner.Scan() {
		val := scanner.Text()
		if val != "" {
			ids = append(ids, val)
		}
	}

	for _, val := range ids {
		twos, threes := countDupes(val)
		if twos {
			totalTwos += 1
		}
		if threes {
			totalThrees += 1
		}
	}

	fmt.Printf("Final 2s [%v], 3s [%v] = %v\n", totalTwos, totalThrees, (totalTwos * totalThrees))

	var matched *string
	for _, a := range ids {
		for _, b := range ids {
			if a == b {
				continue
			}
			if differ, value := differByOne(a, b); differ {
				// we'll find it twice because there are two values that match - make sure they're the same
				if matched != nil {
					if *matched != value {
						panic(fmt.Sprintf("Found too many matched: %s, %s", *matched, value))
					}
				} else {
					matched = &value
					fmt.Printf("This string is common [%s] to [%s] and [%s]\n", *matched, a, b)
				}
			}
		}
	}
}

func countDupes(s string) (hasTwo bool, hasThree bool) {
	counts := make(map[string]int)
	for _, r := range s {
		c := string(r)
		counts[c] += 1
	}

	for _, v := range counts {
		if v == 2 {
			hasTwo = true
		}
		if v == 3 {
			hasThree = true
		}
	}

	return hasTwo, hasThree
}

func differByOne(a, b string) (bool, string) {
	if len(a) != len(b) {
		panic(fmt.Sprintf("lengths don't match: %v != %v", a, b))
	}

	for i, _ := range a {
		minusa := fmt.Sprintf("%s%s", a[:i], a[(i+1):])
		minusb := fmt.Sprintf("%s%s", b[:i], b[(i+1):])
		if minusa == minusb {
			return true, minusa
		}
	}

	return false, ""
}

const boxids = `
uqcipadzntnheslgvjjozmkfyr
uqcipadzwtnhexlzvxjobmkfkr
cqcipadpwtnheslgyxjobmkfyr
ubnipadzwtnheslgvxjobmkfyw
uqcisadzwtnheslgvxjfbmkfor
uqcisaezwtnheslgvxkobmkfyr
uqcguadzwtnheslgvxjobmkfir
uqcipadzmtnhesldvxdobmkfyr
uqcipadzwtzheslgdxjtbmkfyr
uquipadzwtcheslgvxjobmkfbr
uqctpadzwtnhesjbvxjobmkfyr
ueciparzwtnheslgvxjobmkfyx
uqcipadzwtnhessgvxjkbmkfkr
uqcipamzwtnheslgvxioamkfyr
uciizadzwtnheslgvxjobmkfyr
uqcieadzwtnhesfgvxeobmkfyr
fqcipadzwtnreslgvkjobmkfyr
uqcipadzrtnherlgvxjobmklyr
uqcypadzwtnheslgvxjobmkxfr
uqcipadzwtnhemlgvxjobmvfur
uwciuadzwwnheslgvxjobmkfyr
uqcipadzwtnhcscgvxjobmkuyr
upripadzwtnheslovxjobmkfyr
uqcipadzltnheslgvxjobmkftc
uqcipadzwtnheslgvgjobmifsr
uqoipadzwtnheslgvxjosmkfkr
uqcipadzwtbhesrqvxjobmkfyr
uqcipadzwtnheslpvxjobmhfyx
uhcinadzwtnheslgvxjybmkfyr
uqcipadzwtnhhslgvxjabmkbyr
uecipadzwtnheslgvxjobqyfyr
uqcipadfwtnheslwvxjobgkfyr
uqcipadzvtnheshgvxzobmkfyr
fqcipadzwtcheslgvxjobmkfyt
uecipadzwtnheslgpxjbbmkfyr
uqclpadzwtnheslgvnjobukfyr
qqciprdzetnheslgvxjobmkfyr
uqcipahpwtnheslgvxjtbmkfyr
uqcidadzwtnhesljvxyobmkfyr
uqciradswtnqeslgvxjobmkfyr
uqcipadzwtrhmslgvxjobmkfyf
urcipadzjtnheslgvxfobmkfyr
uqcipadzwznheslgvxjobmkfcv
uqcipadowtnheslgyxjobmkfym
uqcigadzwtnheslgvxjoomkmyr
uqjipafzwtnheslgvejobmkfyr
uqcioadzwtnhhslgvxzobmkfyr
uqcgpadkwtnheslgvxjobhkfyr
ufciiadewtnheslgvxjobmkfyr
uqoipadzwtnheslgvxjllmkfyr
uqcipadzutnheslgwxxobmkfyr
uqcipadzwtlheslgaxjobmkfwr
uqcbpadzutnheslgvxjbbmkfyr
uucipadzwvnhesngvxjobmkfyr
uqcifadzwtnceslgvxjoumkfyr
ujcipadzwteheslgvxjobmkfyj
uqcipadzwtnheslqvxjobmkuyp
uqcipadzwtnheslgvxjoxmkxyw
uqcipaduwtnheslgvujmbmkfyr
uicipadnwtnheslgvxjobmbfyr
uqcipadzwteheslgvxjobbmfyr
uqcipadzwgnneslgvxjobmklyr
uqcipadzxtnhwslgvbjobmkfyr
uqcipaxwwtnheslxvxjobmkfyr
uocipadzwtnheslgvxjobqdfyr
uqciaauzwtnheslgtxjobmkfyr
uncipagzwtnkeslgvxjobmkfyr
uqcipadzwtnhehlgvxjohdkfyr
uqcipadzwtnheslgvxjobmspyz
uccipadzwtnhvsltvxjobmkfyr
uacipagzwtnheslgvxjoqmkfyr
tqcipaduwtnheslgvxjobmmfyr
uqcipadzwtnheslgvxqebmifyr
uecipadthtnheslgvxjobmkfyr
uocipadzwtnhdslgvkjobmkfyr
uqcipadtwtnheslgvxhobmufyr
uqkipadzwtnleslgtxjobmkfyr
uqcipadzjunheslgvxjobmnfyr
ubcipadzwtvheslgvxjobmkfyf
uqcipadzwpfheslgvxjsbmkfyr
uocipadzwtndeslgvxjobmmfyr
uqcipadzwtnheslgtxjobhkfyq
uqcipadzwtrheslgvxjobmyfya
uqcipadzwtvheslgvxjolgkfyr
uqcipidzwtaheslgvxjobmkfxr
uyzixadzwtnheslgvxjobmkfyr
uqyihadzwtnhedlgvxjobmkfyr
uqcipadzwtnhesltvejobqkfyr
uqciptdzwtnheslgyxlobmkfyr
uqcipzdzwtnhzslgvxjosmkfyr
uqcipadzwtnbeslgexjobmkfvr
uqcipadzwtnheslcwxjobmkkyr
uqcapadzwcnheslgvxjolmkfyr
uqcjpadzwtnhejlgvxjxbmkfyr
uqcipadwwtxweslgvxjobmkfyr
uqmipadzwtnhezlgvxjobmkyyr
uqcipubzwtnpeslgvxjobmkfyr
uecvpadzwtnheslgvxjocmkfyr
uqcipadzwfnheslgvxjibmkdyr
uqcipadzwtnheslgvxvfbykfyr
uqcipadzwtnheslgvgjoimkfyt
dqcqpaqzwtnheslgvxjobmkfyr
uqcipbdzwtnheslgvxjobmkghr
jqcipadzwtnheslgvxjgbmkzyr
uqcipadzwtnheslgvxqkqmkfyr
uqcipadzptnheslgvxjxbokfyr
uucijadzwtwheslgvxjobmkfyr
uccfpadzwtnheslgvxjobpkfyr
uqcipadzwtnheslgvxjobakeyq
uqcipadzwtnheolgvxqobjkfyr
imiipadzwtnheslgvxjobmkfyr
uqcehadzwtnheslgvxjobmkuyr
uqcipadzztnheslgvxjorokfyr
rqcixadzwtnheelgvxjobmkfyr
uqcipadzwtzheslgvxjodmkfyi
uqcipaezwtnwuslgvxjobmkfyr
uqcipadzwtnheslggxjobjkfyq
uqcipadzwkghesagvxjobmkfyr
uqcypqdzwtnheslgvxjobakfyr
iqcipadzwtnhezltvxjobmkfyr
uxcimadzwtnheslgvxjobmxfyr
uqcipaizwtvhwslgvxjobmkfyr
uqcipafzwtnheslgvxjpbmkfym
uqcipadzwinheslgvxlobmpfyr
uqcupadzwtnheslkvxmobmkfyr
uqcapadzwtnhesrgvxjobmkfsr
urcipafzwtnheslgvxjobmkfur
uqcipaczwtnheslgvbjobmknyr
uqcizadzztgheslgvxjobmkfyr
uqcipfdzwtnhesxgvxjobmkfyw
uqcipbdzwtnhyslgvxjobmcfyr
uqcipadzwanhezlgvxjobmkfwr
uvcipadzwtnheslgvxjbkmkfyr
uqcipajzwtnseslgvxjobmkfyq
uqcipvdzwtnheslgvmlobmkfyr
uqcipadzdgnheslgmxjobmkfyr
uqcipddzwtnhestgvpjobmkfyr
umcipadzwtdheslgvxjzbmkfyr
uqciuwdzwtnheslgvxjobmkflr
uqcipadzwtnheslgsxabbmkfyr
uceipadzwtnheslgvxjobgkfyr
mqcipadzwtnhesrgvxjobmjfyr
aqcipadvwtnheslgvxjobmkryr
uqsipadzwtnofslgvxjobmkfyr
uqcixadzwtfheslgvxjzbmkfyr
uqcipadnwfnheslgvxjohmkfyr
uqcivadzwtnheslfvxjobmkfyz
uqciprdzwtnheslgvxjobmkjir
uqcipadhbtnheslgvxjoxmkfyr
fqcipadzwtnhesfgvxjobmkfye
uqoipqdzwtnheqlgvxjobmkfyr
uqcipadzwtnhesltvxmobmkzyr
uqcipadzwtnhebqgvsjobmkfyr
uqcipadzwtnheslglxjobmfbyr
gqcipadzwtgheslgvxjobwkfyr
uqcipadzwtnheslgfxjzbmlfyr
ujcnpadzwtnheslrvxjobmkfyr
ujcivadzwtnheglgvxjobmkfyr
uqcitadzwgnheslgvxjofmkfyr
uqcipahzatnhmslgvxjobmkfyr
uqzipaizwtnheslgvujobmkfyr
uqcipadzltnheylgvnjobmkfyr
uqcidadzwtnhwsljvxyobmkfyr
uqcipadzwtihetlgvxjobhkfyr
oqcipabzwtnheslgvfjobmkfyr
uqcipadzwtnveslgvxjobzkfzr
uqcipadzwtjheslgqxjobmlfyr
uqcnpadzztnheslgvxjobmkoyr
uqciuadzwonheslgvxjobmkfyz
tqcipadzwtnheslgvxaobmqfyr
uqcipadtwtnhqslgvxjobmkeyr
uqcipadzwbnheslgvajobmsfyr
ubcopadzwtnhgslgvxjobmkfyr
uqcipydzwtwheslgvxjobakfyr
cqbijadzwtnheslgvxjobmkfyr
uscipadowtnheslgvxjobmkfcr
uqcipadzwtgheslnvxjobskfyr
uqcipzdzwtnzeslgkxjobmkfyr
uqcipawzwtnhrslgbxjobmkfyr
uqcipadzatchyslgvxjobmkfyr
uqcipadzotnpeslgvxjobmjfyr
uqcipagzwtnheslgvxjobmvfyt
uqcipadzwhnheslgvxyobmkfyo
uqcipadzwtnheslgmqjobmkfyc
uqcupadzwgnheslgvcjobmkfyr
uqcipabzwbnheslgvxjobmkwyr
uqciiadzwtnheslgvxjobmkfmz
uqkipauzwtnheslgvxjjbmkfyr
uqcipidzetnheslgvxjobmkfyi
uqcipadzwtnheslgqxjokmkfmr
uqcipadzqtnhesllvxjobmkfyk
uqccpadzwtnheslgmxsobmkfyr
uqcipadzwteheslgvljfbmkfyr
uqcipadxwinheslgaxjobmkfyr
uqcipadzwtnheslhvxyobmkfjr
aqcipadzwnnheslgvxjqbmkfyr
uvcipadzwtnheszgvxjobmkfyg
uqcipahzmtnheslgvxjobmkfir
ukcipadzbtnheslgvxjobmkfyb
uqcipadzwtnhemlgvqjobmkfpr
uqcipadzwtnheslgvmeobmkfpr
uqciphdrwtnheslgvxjobmkfyw
uqcipadzwtnheslevxqobzkfyr
uqcipadzwknzeslgvxnobmkfyr
wqcipadzwjnheslgvxjobbkfyr
uqcipadzwtdheslgvmjobmkjyr
uqvipadzwtnhextgvxjobmkfyr
uqhipadzwtnheslwvxjzbmkfyr
uqcipadzwtnherlgsxjobmksyr
uqcipadzwtnhesqgvxjotmvfyr
udcipadzwtnhekwgvxjobmkfyr
uqcjprdzwtnheslgvxjobmkfpr
uqcipadzatnheclgvqjobmkfyr
uqcbpadzctnheslqvxjobmkfyr
uqcipadzqtnhesluvxjobrkfyr
uqcipadzwtnhcslgvxjoomwfyr
uqcppadzwxnheslgwxjobmkfyr
uqcipadcwtnheslrvxjdbmkfyr
ukcipadzwtnhhslgvxjobmkgyr
uqckpadzwtnheslgvxjokmkiyr
uqcspadzwtjheslgvxjobmkfjr
uqcipadpwtnhsslgvxjobmkfyu
uqcepadzwtnheilgvbjobmkfyr
jqcipadiwtnheslgvxjobmkjyr
uqcipadzrtnseslgqxjobmkfyr
sqmipadzwtnhewlgvxjobmkfyr
uqcieadzhtnheslgvgjobmkfyr
uqcipadzwkwhewlgvxjobmkfyr
uqcipadzwtzheslgvxjpbqkfyr
uzcipadzjtnheslgvxjobmlfyr
uqcipadzwtnheslnvxjobmkfee
uqciyanzwtnheslgvxjoimkfyr
uqcipadqwtnheswghxjobmkfyr
uycipadzwtnheslovxjobmofyr
uqcipadzwtnheslgvxcozmxfyr
uqmipadzwtnxezlgvxjobmkfyr
uqcipadzftnheslgvxjotmkffr
aqcipaizwtnhesagvxjobmkfyr
uqcipcdzwtnheslgoajobmkfyr
uqcypadgwtnhesbgvxjobmkfyr
uqcipcdzwtnheslgvxjebmkfyb
uhcvpadzwtnheslgvxjobzkfyr
uqcipadzwtnpesagvxmobmkfyr
uqcipadzwtnidslgvxjobmkfor
uqcipadkwtnhesigvxjzbmkfyr
uqcypadlwtnheslsvxjobmkfyr
qqcipadzwtnheswgvxjobmkoyr
uqcipadzwtnheslgvxjhbmmcyr
uqcipadzwtnhesogvxjormkfmr
uqcipadzwtnhetcgvxgobmkfyr
`
