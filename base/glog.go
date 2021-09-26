package main

import (
  "flag"
  "fmt"
  "github.com/golang/glog"
)

// 避免没有引用fmt的编译错误
var _ = fmt.Println

func main() {

  logdir := "/tmp/log"
  myV := "2"  // info Verbose

  // init command line args, as: ./glog -log_dir="./" -v=3
  flag.Parse()

  flag.Lookup("logtostderr").Value.Set("false")
  flag.Lookup("log_dir").Value.Set(logdir)
  flag.Lookup("v").Value.Set(myV)

  glog.Info("hello, glog")
  glog.Warning("warning glog")
  glog.Error("error glog")

  glog.Infof("info %d", 1)
  glog.Warningf("warning %d", 2)
  glog.Errorf("error %d", 3)

  glog.V(3).Infoln("info with v 3")
  glog.V(2).Infoln("info with v 2")
  glog.V(1).Infoln("info with v 1")
  glog.V(0).Infoln("info with v 0")

  // 退出时调用，确保日志写入文件中
  glog.Flush()
}
