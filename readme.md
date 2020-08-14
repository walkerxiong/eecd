#CCED
常用字符集编解码 (Common Charset Encoder and Decoder)

## 使用方式
```golang
    buff, err := ioutil.ReadFile("pathtofile")
	if err != nil {
        // handle err
	}
	res, err := GbkToUtf8(buff)
```

## 编码转换支持类型

- GBK -> UTF8
- GB2312 -> UTF8