#/bin/bash
APPNAME=go-get
PORT=8002

start_server(){
    echo "Build Server"
    go build
    echo "Start Server"
    nohup ./$APPNAME &
}

stop_server(){
    echo "Stop Server"
    for pid in `ps x | grep $APPNAME| grep -v grep | awk '{print $1}'`;do
        echo "SHUTDOWN process $pid"
        kill -3 $pid
    done
}


restart_server(){
    stop_server
    start_server
}


case $1 in 
    start)
        start_server
        ;;
    stop)
        stop_server
        ;;
    restart)
        restart_server
        ;;
    *)
        echo "Usage: $SCRIPT_NAME (start|stop|restart)"
        ;;
esac