package redis_wrapper

/**
 *@author LanguageY++2013
 *2019/2/20 5:31 PM
 **/
func LPush(key string, value []byte) error {
	return wrapper.LPush(key, value)
}

func RPop(key string) ([]byte, error) {
	return wrapper.RPop(key)
}