## 解析各大视频站下载视频

### 1，API接口形式提供解析服务
```console  
1，在项目根目录下执行 go run main.go -p=api -port=8002 (prot默认参数为8002)
2，请求格式：IP:8002/video/info?url=http://v.youku.com/v_show/id_XMjgyODc0NTU2MA==.html

返回数据格式：

{
    "createTime": 1516670558949,
    "desc": "normal 表示标清，hd1 表示高清，hd2 表示超清，hd3 表示720p hd4 表示1080p",
    "downloadInfo": {
        "hd1": {
            "m3u8_url": "http://pl.cp31.ott.cibntv.net/playlist/m3u8?vid=XMzM0MDY3MDk5Mg%3D%3D&type=mp4&ups_client_netip=781a0dda&ups_ts=1516670558&utid=XnztEiZ8v2MCAXgaDdoE5iGF&ccode=0508&psid=fc8204226369f98c569b41e6f30e5597&duration=63&expire=18000&ups_key=1a8d930e6cd01cf71d7419e79d946515",
            "urls": [
                "http://k.cp31.ott.cibntv.net/player/getFlvPath/sid/0516670558928126ccb2d/st/mp4/fileid/03000801005A655B51309801233E8D7A152D9E-7D04-A5B9-65B3-91CDDEFE1568?k=d383f06d12a4b949282cffd9&hd=1&myp=0&ts=63&ctype=12&ev=1&token=0544&oip=2014973402&ep=cieVHE%2BKVssF7SrdgD8bYX%2BxJnVbXP4J9h%2BFidJjALshOe66nEzUxp6wR%2FdCF%2FsacVcOZOyAqtXl%0AHklhYfc3qWwQrkvaMPrm%2B4Lg5aRSt%2BR0E29Dc8jRvFSeRjT1&ccode=0508&duration=63&expire=18000&psid=fc8204226369f98c569b41e6f30e5597&ups_client_netip=781a0dda&ups_ts=1516670558&ups_userid=&utid=XnztEiZ8v2MCAXgaDdoE5iGF&vid=XMzM0MDY3MDk5Mg%3D%3D&vkey=Aaa19ebf50308e204cad3d5e18f01c3c8"
            ]
        },
        "hd2": {},
        "mp4hd2v2": {},
        "mp4sd": {},
        "normal": {}
    },
    "duration": 63,
    "title": "萌娃边哭边演出 网友: 差个敬业奖",
    "url": "http://v.youku.com/v_show/id_XMzM0MDY3MDk5Mg==.html"
}   

备注：有些站点解析出的下载地址，会增加一些header的认证，如果需要，回把header数据一并返回，下载时只需加上返回的header即可   
```    

### 2，本地下载模式

```console  
$go run main.go http://v.youku.com/v_show/id_XMjgyODc0NTU2MA==.html
        
site:               youku
title:              朴志浩视角 韩服 VS 国服 4V4 刘勇赫朴敏秀朴志浩KoguryoTeam SSS张博麟涛XX 20170614 A
type:               hd2
urls:               19
size:               814.13 MiB (853675175 bytes)
Downloading 朴志浩视角 韩服 VS 国服 4V4 刘勇赫朴敏秀朴志浩KoguryoTeam SSS张博麟涛XX 20170614 A ...
814.13 MiB/814.13 MiB [=======================================================================] 100%  24s
Saving Me at the 朴志浩视角 韩服 VS 国服 4V4 刘勇赫朴敏秀朴志浩KoguryoTeam SSS张博麟涛XX 20170614 A ...Done.
Merge Video at the 朴志浩视角 韩服 VS 国服 4V4 刘勇赫朴敏秀朴志浩KoguryoTeam SSS张博麟涛XX 20170614 A ...Done.
```

#### 备注：
        
    本地下载需要视频合并，需要安装ffmpeg工具 