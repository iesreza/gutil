package path

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type file struct {
	path string
	mu   sync.Mutex
}

// File init file
func File(path string) *file {
	f := file{}
	f.path = path
	return &f
}

// Append append to file
func (f *file) Append(text string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	fp, err := os.OpenFile(f.path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer fp.Close()

	if _, err = fp.WriteString(text); err != nil {
		return err
	}

	return nil
}

// Write write to file
func (f *file) Write(text string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	ioutil.WriteFile(f.path, []byte(text), 0644)
	return nil
}

// WriteBytes write to file
func (f *file) WriteBytes(data []byte) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	ioutil.WriteFile(f.path, data, 0644)
	return nil
}

// Copy copy file to destination
func (f *file) Copy(dest string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	in, err := os.Open(f.path)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	out.Close()
	return nil
}

// Copy copy file from destination
func (f *file) CopyFrom(src string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(f.path)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	out.Close()
	return nil
}

// Create create file at location and write data
func (f *file) Create(data string) error {
	f.Dir().Create()
	return ioutil.WriteFile(f.path, []byte(data), 0644)
}

// GetDir return base directory
func (f *file) GetDir() string {
	return filepath.Dir(f.path)
}

// Dir return base directory as directory pointer
func (f *file) Dir() *dir {
	return Dir(f.GetDir())
}

// Remove remove file
func (f *file) Remove() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	return os.Remove(f.path)
}

// Remove move file
func (f *file) Move(dest string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	err := f.Copy(dest)
	if err != nil {
		return err
	}
	return f.Remove()
}

// Exist check if file exists
func (f *file) Exist() bool {
	if _, err := os.Stat(f.path); os.IsNotExist(err) {
		return false
	}

	return true
}

// Absolute return absolute path of file
func (f *file) Absolute() (string, error) {
	return filepath.Abs(f.path)
}

// Content return file content as string
func (f *file) Content() (string, error) {
	fp, err := ioutil.ReadFile(f.path) // just pass the file name
	if err != nil {
		return "", err
	}

	return string(fp), nil

}

func (f *file) JSON(output *interface{}) error {
	f.mu.Lock()
	b, err := ioutil.ReadFile(f.path)
	f.mu.Unlock()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, output)
}
