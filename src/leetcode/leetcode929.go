package leetcode

func numUniqueEmails(emails []string) int {
	//m.y+name@email.com 将转发到 my@email.com
	emailsSet := make(map[string]struct{})
	for _, emStr := range emails {
		local := make([]byte, 0, len(emStr))
		domain := make([]byte, 0)
		var domainIndex int
		for i, c := range emStr {
			if c == '@' {
				domain = append(domain, emStr[i:]...)
				domainIndex = i
			}
		}
		for _, c := range emStr[:domainIndex] {
			if c == '+' {
				break
			} else if c == '.' {
				continue
			} else {
				local = append(local, byte(c))
			}
		}
		local = append(local, '@')
		local = append(local, domain...)
		emailsSet[string(local)] = struct{}{}
	}
	return len(emailsSet)
}
