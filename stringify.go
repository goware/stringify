package stringify

import "fmt"

func Stringify(src, dest interface{}) error {
	return fmt.Errorf("TODO")
}

func StringifyMap(src, dest map[string]interface{}) error {
	for k, v := range src {
		switch t := v.(type) {
		case fmt.Stringer:
			dest[k] = t.String()
		case string:
			dest[k] = t
		case nil:
			dest[k] = ""
		case map[string]interface{}:
			m := map[string]interface{}{}
			if err := StringifyMap(t, m); err != nil {
				return err
			}
			dest[k] = m
		case []string:
			//..
		case []interface{}:
			// ..
		case []byte:
			dest[k] = string(t)
		default:
			// return fmt.Errorf("StringifyMap: unknown type '%T'", t)
			dest[k] = fmt.Sprintf("%v", t)
		}
	}
	return nil
}

func StringifyArray(src, dest []interface{}) error {
	return fmt.Errorf("TODO")
}
