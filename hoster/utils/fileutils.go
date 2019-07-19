package utils

type IRulesEnvFile interface {
	CommentReprBegin() string
	CommentReprEnd() string
	FilePath() string
	FileBytes() []byte
}

type rulesEnvFile struct {
	filePath string
}

func (*rulesEnvFile) CommentReprBegin() string {
	panic("implement me")
}

func (*rulesEnvFile) CommentReprEnd() string {
	panic("implement me")
}

func (*rulesEnvFile) FilePath() string {
	panic("implement me")
}

func (*rulesEnvFile) FileBytes() []byte {
	panic("implement me")
}

type IFileUtils interface {
	InjectRule(file IRulesEnvFile)
	ClearRule(file IRulesEnvFile)
	ClearAll()
}

type fileUtils struct {
	envFile IRulesEnvFile
}

func (*fileUtils) InjectRule(file IRulesEnvFile) {
	panic("implement me")
}

func (*fileUtils) ClearRule(file IRulesEnvFile) {
	panic("implement me")
}

func (*fileUtils) ClearAll() {
	panic("implement me")
}

func NewFileUtils(filePath string) IFileUtils {
	return &fileUtils{
		envFile: &rulesEnvFile{filePath: filePath},
	}
}
