package main

import (
	"html/template"
	"log"
	"os"
)

func buildFromGlob(file, glob string, data map[string]interface{}) error {
	out, err := os.Create(file)
	defer out.Close()
	if err != nil {
		return err
	}
	t, err := template.ParseGlob(glob)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}

func buildFromFile(file, tmpl string, data map[string]interface{}) error {
	out, err := os.Create(file)
	defer out.Close()
	if err != nil {
		return err
	}
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}

type some struct {
	Name string
}

func main() {
	var err error

	data := new(some)
	data.Name = "Rob"

	t, err := template.ParseFiles("some.yaml")
	if err != nil {
		log.Println(err)
	}
	t.Execute(os.Stdout, data)
	if err != nil {
		log.Println(err)
	}

	/*
		data := map[string]interface{}{}
		buf, err := os.ReadFile("data.yml")
		if err != nil {
			return
		}
		err = yaml.Unmarshal(buf, &data)
		if err != nil {
			return
		}
		entries, err := os.ReadDir("tmpl")
		if err != nil {
			return
		}
		for _, entry := range entries {
			fmt.Println(entry.Name())
		}
	*/

	// TODO detect tmpl and fail if not found
	// TODO iterate over tmpl directory and
	//    if directory build a file matching the name of the directory
	//    from the files in the directory
	//    or,
	//    if a file just build from that file
	//    make sure to detect hte template/html or template/text based on
	//    suffix
}
