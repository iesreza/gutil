package path

import (
	"os"
	"path/filepath"
	"strings"
)

type dir struct{
	path string
}

// Dir initialize directory
func Dir(path string) *dir {
	p := dir{}
	p.path = path
	p.path = strings.TrimRight(strings.TrimRight(p.path,"/"),"\\") + "/"
	return  &p
}

// Current get pointer to current directory
func Current() (*dir,error){
	path,err := os.Getwd()
	if err != nil{
		return nil,err
	}
	return Dir(path),nil
}
// Exist check if path is directory & exists
func (p *dir)Exist() bool  {
	return p.IsDir()
}

// IsDir check if path is directory & exists
func (p *dir)IsDir() bool {
	if stat, err := os.Stat(p.path); err == nil && stat.IsDir() {
		return true
	}
	return false
}

// Create create path
func (p *dir)Create() error {
	err := os.MkdirAll(p.path,0755)
	if err != nil {
		return err
	}
	return nil
}

// Find find files and directories in path also support * wildcard
func (p *dir)Find(criteria string) ([]string , error) {
	return filepath.Glob(p.path + criteria)
}

// Remove remove the path
func (p *dir)Remove() error  {
	return  os.RemoveAll(p.path)
}

// Remove remove contents inside the path
func (p *dir)Clean() error  {
	err :=  os.RemoveAll(p.path)
	if err != nil{
		return err
	}
	p.Create()
	return nil
}

// File get file pointer
func (p *dir)File(filename string) *file  {
	return File(p.path+filename)
}

// Parent go to parent directory
func (p *dir)Parent() *dir  {
	return Dir(filepath.Base(filepath.Dir(p.path)))
}