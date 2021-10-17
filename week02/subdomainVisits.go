package week02

import (
	"strconv"
	"strings"
)

/**
811
https://leetcode-cn.com/problems/subdomain-visit-count/submissions/
 */

/**
主要思想利用hash记录每次访问过的值累加
 */

func subdomainVisits(cpdomains []string) []string {
	countMap := map[string]int{}

	for _, cpdomain := range cpdomains {

		// 切割次数 和 域名
		countDomain := strings.Split(cpdomain, " ")
		count,_ := strconv.Atoi(countDomain[0])
		countMap[countDomain[1]] += count

		domainArr := strings.Split(countDomain[1], ".")
		countMap[domainArr[len(domainArr)-1]] += count

		if len(domainArr) > 2 {
			countMap[domainArr[len(domainArr)-2] +"."+domainArr[len(domainArr)-1]] += count
		}
	}

	result := []string{}

	for key,val := range countMap {
		result = append(result, strconv.Itoa(val) + " " + key)
	}

	return result
}