package localization

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Lang struct {
	Language string `yaml:"languageMain"`

	Lang string `yaml:"lang"`

	EntryLabel string `yaml:"entryName"`
	EntryInput string `yaml:"textinput"`

	ConnectUser string `yaml:"connectUser"`
	Chats       string `yaml:"chatList"`
	Settings    string `yaml:"settignsList"`

	BackMain string `yaml:"backMain"`

	HeaderConnect string `yaml:"headerConnect"`
	ExitButton    string `yaml:"exitButton"`
	MoveLists     string `yaml:"moveMainButton"`
	MoveToLists   string `yaml:"moveAnotherList"`
	ScrollMessage string `yaml:"sctollMessageList"`

	Connectig        string `yaml:"connecting"`
	ConnectError     string `yaml:"connectError"`
	ConnectSucessful string `yaml:"connectSucessful"`
	ConnectTimeout   string `yaml:"connectTimeout"`
}

func LangRead(path string) ([]Lang, error) {

	files, err := checkfiles(path)
	if err != nil {
		return nil, err
	}
	var langs []Lang
	for _, i := range files {
		data, err := os.ReadFile(i)
		l := defaultLang()
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, &l)
		if err != nil {
			return nil, err
		}
		langs = append(langs, l)
	}
	if len(langs) == 0 {
		langs = append(langs, defaultLang())
	}
	return langs, nil
}

func checkfiles(path string) ([]string, error) {
	var files []string
	n, err := os.ReadDir(path)
	if err != nil {
		return files, err
	}
	for _, i := range n {
		if !i.IsDir() {
			files = append(files, fmt.Sprintf("%s/%s", path, i.Name()))
		}
	}
	return files, nil
}

func defaultLang() Lang {
	return Lang{
		Language: "Русский",
		Lang:     "Язык",

		EntryLabel: "Необходимо ввестии имя под которым вас будут видеть другие пользователи",
		EntryInput: "Ваше имя",

		ConnectUser: "Подключиться",
		Chats:       "Чаты",
		Settings:    "Настройки",

		BackMain: "Назад",

		HeaderConnect: "Введите адрес пользователя",
		ExitButton:    "Ctrl+C — выход",
		MoveLists:     "↑/↓ — листать список",
		MoveToLists:   "Tab — для перемещеня",
		ScrollMessage: "Ctrl+W / Ctrl+S — листать",

		Connectig:        "Подключение успешно",
		ConnectError:     "Подлкючение...",
		ConnectSucessful: "Ошибка подлкючения",
		ConnectTimeout:   "Подлкючение к пользовтелю заняло слишком много времени",
	}
}
