package elastic

var FailedKeys []*FailedKey


func AddFailedKey(username, key, src_ip, dest_ip, actionType, category string) {
	FailedKey := &FailedKey{
		Username: username,
		Key: key,
		SourceIP: src_ip,
		DestinationIP: dest_ip,
		Type : actionType,
		Category : category,
		FailedCount: 1,
	}
	FailedKeys = append(FailedKeys, FailedKey)
}

func IncreaseFailedCount(fk *FailedKey) {
	fk.FailedCount += 1
}

func FindIP(src_ip string) *FailedKey {
	for i, keys := range FailedKeys {
		if keys.SourceIP == src_ip {
			return FailedKeys[i]
		}
	}
	return nil
}