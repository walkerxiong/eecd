
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
    
    上述完成了Gbk转换unicode,Unicode作为一种符号集，其实现方式主要有Utf8/Utf16/Utf32
    这几种编码，Utf8是个变长的编码实现方式。对于低于0x80的即为ACII2编码方式，而对于多字节
    ，则第一个字节最高位几个1代表有几个字节，后续字节为10开头，如下表所示：
    
    
    Unicode符号范围      |        UTF-8编码方式
    (十六进制)           |            （二进制）
    ----------------------+---------------------------------------------
    0000 0000-0000 007F | 0xxxxxxx
    0000 0080-0000 07FF | 110xxxxx 10xxxxxx
    0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
    0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxx
    
    unicode中的utf16的编码，与utf8的变长不同，utf16是按照16位也就是2个字节来编码，当unicode
    编码小于等于0xffff的时候，直接复用unicode的编码。当unicode编码大于0xffff的时候，这个时候
    utf16编码的规则为以下几个步骤：
        1. unicode编码减去0x10000
        2. 取高位10bits与0xD800相加
        3. 取低位10bits与0xDC00相加
    经过上面的编码即可形成utf16位编码，但是utf16有大小尾编码序，后续更新再做说明，暂时不支持
    大小尾序。 

## 编码转换支持类型

- GBK -> UTF8
- GB2312 -> UTF8
- GBK -> UTF16
- GBK -> UTF16