package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Example_write() {
	file, err := os.OpenFile("xx.out",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		defer file.Close()

		file.WriteString("some string data")
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func Example_write_json() {
	file, err := os.OpenFile("xx.json",
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.Encode("some json data")
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func Example_createTempFile() {
	file, err := os.CreateTemp(".", "temp-*.json")
	if err != nil {
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.Encode("some json data")
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func Example_commonLocations() {
	workDir, _ := os.Getwd()
	Printfln("work dir: %v", workDir)

	homeDir, _ := os.UserHomeDir()
	Printfln("home dir: %v", homeDir)

	cacheDir, _ := os.UserCacheDir()
	Printfln("cache dir: %v", cacheDir)

	configDir, _ := os.UserConfigDir()
	Printfln("config dir: %v", configDir)

	Printfln("temp dir: %v", os.TempDir())
	// Output:
}

func Example_filepath() {
	home, err := os.UserHomeDir()
	var path string
	if err == nil {
		path = filepath.Join(home, "MyApp", "MyTemp.json")
	}

	Printfln("Full path: %v", path)
	Printfln("Volume name: %v", filepath.VolumeName(path))
	Printfln("Dir: %v", filepath.Dir(path))
	Printfln("Base: %v", filepath.Base(path))
	Printfln("File extension: %v", filepath.Ext(path))

	// Output:
	//
}

func Example_ReadDir() {
	//path, err := os.Getwd()
	path, err := os.UserHomeDir()
	if err == nil {
		dirEntries, err := os.ReadDir(path)
		if err == nil {
			for _, d := range dirEntries {
				Printfln("Entry name: %v, IsDir: %v",
					d.Name(), d.IsDir())
			}
		}
	}

	if err != nil {
		Printfln("Error: %v", err.Error())
	}

	// Output:
	//
}

func Example_exists() {
	targetFiles := []string{"not_exist_file.txt", "config.json"}
	for _, name := range targetFiles {
		info, err := os.Stat(name)
		if os.IsNotExist(err) {
			Printfln("File not exist: %v", name)
		} else if err != nil {
			Printfln("Other error: %v", err.Error())
		} else {
			Printfln("File name: %v, Size: %v", info.Name(), info.Size())
		}
	}

	// Output:
	//File not exist: not_exist_file.txt
	//File name: config.json, Size: 262
}

func Example_locatingFilesWithPattern() {
	path, err := os.Getwd()
	if err == nil {
		matches, err := filepath.Glob(filepath.Join(path, "*.go"))
		if err == nil {
			for _, m := range matches {
				Printfln("Match: %v", m)
			}
		}
	}

	if err != nil {
		Printfln("Error: %v", err.Error())
	}

	// Output:

}

func Example_WalkDir() {
	path, err := os.Getwd()
	if err == nil {
		err = filepath.WalkDir(path, callback)
	}

	if err != nil {
		Printfln("Error: %v", err.Error())
	}

	// Output:

}
