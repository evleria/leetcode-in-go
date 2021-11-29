package _721_accounts_merge

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	// key: email, value: index of accounts
	emailMap := make(map[string]int)
	// indexConnections, initial state: indexConnections[i]=i
	// if indexConnections[i]=j, i and j are connected
	indexConnections := make([]int, len(accounts))
	for i := range indexConnections {
		indexConnections[i] = i
	}

	for accountIndex, emails := range accounts {
		for i := 1; i < len(emails); i++ {
			if existIndex, exist := emailMap[emails[i]]; exist {
				for {
					// connect existIndex and accountIndex by set indexArray[existIndex] as accountIndex
					// if existIndex has connect to another index, we need connect them together
					if indexConnections[existIndex] == existIndex {
						indexConnections[existIndex] = accountIndex
						break
					}
					existIndex, indexConnections[existIndex] = indexConnections[existIndex], accountIndex
				}
			} else {
				emailMap[emails[i]] = accountIndex
			}
		}
	}

	// after this, we will connect all emails belong to same person
	for i := len(indexConnections) - 1; i >= 0; i-- {
		indexConnections[i] = indexConnections[indexConnections[i]]
	}

	// key: accounts index, value: [name, emails...]
	magic := make(map[int][]string)
	for email, index := range emailMap {
		if _, exist := magic[indexConnections[index]]; !exist {
			magic[indexConnections[index]] = []string{accounts[index][0]}
		}
		magic[indexConnections[index]] = append(magic[indexConnections[index]], email)
	}

	// trans map to slice, sort emails
	retSlice := make([][]string, 0, len(magic))
	for _, v := range magic {
		sort.Strings(v[1:])
		retSlice = append(retSlice, v)
	}
	return retSlice
}
