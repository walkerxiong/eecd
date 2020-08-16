
# CCED
常用字符集编解码 (Common Charset Encoder and Decoder)

## 使用方式
```golang
    
buff, err := ioutil.ReadFile("pathtofile")
if err != nil {
   // handle err
}
res, err := GbkToUtf8(buff)
if err != nil{
    // handle err
}
fmt.Printf("%s",res) // output : utf8 encode
```

## 背景
    GB2312为继ACII2码之后，支持中文编码的一个国标性编码标准，不过收录汉字6763个，虽然基本
    满足了汉字的计算机需要，但是对于一些繁体字和汉字不能处理，后继GBK是对于GB2312的扩展规范，
    兼容于GB2312。

## 实现
    基于微软提供的CP936字码表完成对于GBK到Unicode编码的转换工作，由于Gbk编码是两个字节，Gbk
    本身对于ACII2(小于128)的编码兼容，所以只用比对高0x80的字节(Gb2312的高低位编码多加了0xao，
    而对于这个故事在知乎有对应故事https://www.zhihu.com/question/21918229/answer/29918779)。
    
    上述完成了Gbk转换unicode,Unicode作为一种符号集，而对于Unicode的实现方式有Utf8/Utf16/Utf32
    这几种编码，对于Utf8本身是个变长的编码实现方式。对于低于0x80的即为ACII2编码方式，而对于多字节
    编码方式，则是第一个字节最高位几个1代表有几个字节，后续字节为10开头。
    
    
    Unicode符号范围      |        UTF-8编码方式
    (十六进制)           |            （二进制）
    ----------------------+---------------------------------------------
    0000 0000-0000 007F | 0xxxxxxx
    0000 0080-0000 07FF | 110xxxxx 10xxxxxx
    0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
    0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxx
## 编码转换支持类型

- GBK -> UTF8
- GB2312 -> UTF8