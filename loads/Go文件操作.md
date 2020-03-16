# Go文件操作

## 打开和关闭文件

`os.Open()`函数能够打开一个文件，返回一个`*File`和一个`err`。对得到的文件实例调用`close()`方法能够关闭文件。

## 读取文件

### file.Read()

**基本使用**

```go
func (f *File) Read(b []byte) (n int, err error)
```

它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回`0`和`io.EOF`。 举个例子：

```go
func main() {
	file,err :=os.Open("./xx.txt")
	defer file.Close()
	if err != nil{
		fmt.Println(err)
		return
	}
	b := make([]byte,128)
	n,err :=file.Read(b)
	if err != nil{
		fmt.Println("read file failed，err:",err)
	}
	if err == io.EOF{
		fmt.Println("文件读完了")
	}
	fmt.Printf("读取了%d个字节\n",n)
	fmt.Println(string(b))


}
```

### 循环读取

使用for循环读取文件中的所有数据。

```go
func main() {
	file,err :=os.Open("./xx.txt")

	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	b := make([]byte,128)

	var content []byte
	for {
		n,err :=file.Read(b)
		if err == io.EOF{
			fmt.Println("文件已经读完了")
			break
		}
		if err != nil{
			fmt.Println("打开文件失败,err:",err)
			return
		}

		content = append(content, b[:n]...)
	}
	fmt.Println(string(content))



}
```

## bufio读取数据

```go
func main() {
	file,err :=os.Open("./xx.txt")
	if err !=nil{
		fmt.Println("文件打开错误,err:",err)
	}
	defer file.Close()
	reader :=bufio.NewReader(file)
	for {
		str,err :=reader.ReadString('\n')
		if err == io.EOF{
			fmt.Println("文件已经读完了")
			break
		}
		if err !=nil{
			fmt.Println("读取失败")
			return
		}
		fmt.Printf(str)

	}
}
```

## io/ioutil

```go
// ioutil 读取整个文件
func readFile(f *os.File){
	bytes,err :=ioutil.ReadAll(f)
	if err != nil{
		fmt.Println("读取失败,err:",err)
	}
	fmt.Println(string(bytes))
}

func readFile1(f string){
	b,err :=ioutil.ReadFile(f)
	if err != nil{
		fmt.Println("读取失败,err:",err)
		return
	}
	fmt.Println(string(b))
}
```





## 文件写入操作

`os.OpenFile()`函数能够以指定模式打开文件，从而实现文件写入相关功能。

```go
func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
```

其中：

`name`：要打开的文件名 `flag`：打开文件的模式。 模式有以下几种：

|     模式      |   含义   |
| :-----------: | :------: |
| `os.O_WRONLY` |   只写   |
| `os.O_CREATE` | 创建文件 |
| `os.O_RDONLY` |   只读   |
|  `os.O_RDWR`  |   读写   |
| `os.O_TRUNC`  |   清空   |
| `os.O_APPEND` |   追加   |

`perm`：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。



```go
func writeFile(f *os.File,str string){
	b:= []byte(str)
	n,err:=f.Write(b)
	if err !=nil{
		fmt.Println("写入失败")
	}else{
		fmt.Printf("写入%d字节\n",n)
	}
}

func wirteFile2(f *os.File,str string){
	n,err:=f.WriteString(str)
	if err !=nil{
		fmt.Println("写入失败")
	}else{
		fmt.Printf("写入%d字节\n",n)
	}
}
```

## bufio.NewWriter

```go
func main() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}
```

## ioutil.WriteFile

```go
func main() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
```