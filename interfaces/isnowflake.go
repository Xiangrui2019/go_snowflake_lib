package interfaces

type ISnowflake interface {
	// 生成下一个Id
	NextId() (int64, error)
	// 解析生成的Id包含的内容
	ParseId(id int64) ()
}
