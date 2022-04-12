package resources

/**
 * @author  巨昊
 * @date  2021/9/14 10:48
 * @version 1.15.3
 */

type Resources struct {
	Name  string
	Verbs []string
}
type GroupResources struct {
	Group     string
	Version   string
	Resources []*Resources
}
