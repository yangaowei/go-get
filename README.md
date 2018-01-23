## 解析各大视频站下载视频

###1，API接口形式提供解析服务

       1，在项目根目录下执行 go run main.py -p=api -port=8002 (prot默认参数为8002)
       2，请求格式：IP:8002/video/info?url=http://v.youku.com/v_show/id_XMjgyODc0NTU2MA==.html
       返回解析数据,数据格式为json
       
###2，本地下载模式

```console  
$go run main.py http://v.youku.com/v_show/id_XMjgyODc0NTU2MA==.html
        
site:               youku
title:              朴志浩视角 韩服 VS 国服 4V4 刘勇赫朴敏秀朴志浩KoguryoTeam SSS张博麟涛XX 20170614 A
type:               hd2
urls:               19
size:               814.13 MiB (853675175 bytes)
Downloading 朴志浩视角 韩服 VS 国服 4V4 刘勇赫朴敏秀朴志浩KoguryoTeam SSS张博麟涛XX 20170614 A ...
814.13 MiB/814.13 MiB [====================================================================================================] 100%  24s
Saving Me at the 朴志浩视角 韩服 VS 国服 4V4 刘勇赫朴敏秀朴志浩KoguryoTeam SSS张博麟涛XX 20170614 A ...Done.
Merge Video at the 朴志浩视角 韩服 VS 国服 4V4 刘勇赫朴敏秀朴志浩KoguryoTeam SSS张博麟涛XX 20170614 A ...Done.
```

####备注：
        
    本地下载需要视频合并，需要安装ffmpeg工具 