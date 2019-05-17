# Gutil
Golang utilities

```
func main()  {

	fmt.Println( str.S(1.00000045).Trim("5").Quote().ReplaceAll("0","9") )

	logExample()
	fileExample()


}

func logExample()  {

	//Logger
	log.Critical("This is Critical!")
	log.Debug("This is Debug!")
	log.DebugF("Here are some numbers: %d %d %f", 10, -3, 3.14)
	// Give the Warning
	log.Warning("This is Warning!")
	log.WarningF("This is Warning!")
	// Show the error
	log.Error("This is Error!")
	log.ErrorF("This is Error!")
	// Notice
	log.Notice("This is Notice!")
	log.NoticeF("%s %s", "This", "is Notice!")
	// Show the info
	log.Info("This is Info!")
	log.InfoF("This is %s!", "Info")

	log.SetLogLevel( logger.ErrorLevel )

	log.SetFormat("[%{module}] [%{level}] %{message}")
	log.Warning("This is Warning!") // output: "[test] [WARNING] This is Warning!"
	// Also you can set your format as default format for all new loggers
	logger.SetDefaultFormat("%{message}")


}

func fileExample()  {

	//create directory pointer
	dir := path.Dir("c:/workingdir/test1/test2")

	//check if directory exists
	if !dir.Exist() {
		//create directory if not exist
		dir.Create()
	}

	//create file pointer inside directory
	fp := dir.File("test.txt")

	//create file with given content
	fp.Create("test content here")

	//append to file
	fp.Append("\nNew line")

	//get file content
	content,_ := fp.Content()
	fmt.Println(content)

	//return absolute path of file
	abs,_ := fp.Absolute()
	fmt.Println(abs)

	//create pointer to file
	fp = path.File("c:/workingdir/test1/test2/test2.txt")

	//create file with given content
	fp.Create("")

	//write text to file
	fp.Write("Hi")

	//remove file
	fp.Remove()

	//find files inside directory
	files,_ := dir.Find("*,txt")
	fmt.Println(files)

	//go to parent dir
	parent := dir.Parent()
	parent.File("parent.file").Create("test data")

	//get current working directory
	current,_ := path.Current()
	files,_ = current.Find("*")
	fmt.Println(files)

	//remove directory files
	dir.Clean()

	//remove directory
	dir.Remove()




}

```