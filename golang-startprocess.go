package main
import(
"os"
)

func main() {
// if len(os.Args) < 2 {
// println("usage: supervisor program arg1 arg2 ...")
// return
// }
//
 procAttr := new(os.ProcAttr)
var args []string
  args = "www.baidu.com -c1"
// if len(os.Args) == 2 {
// args = nil
// } else {
// args = os.Args[2:]
// }

//process, err := os.StartProcess(os.Args[1], args, procAttr)
process, err := os.StartProcess('/bin/ping', []string{'www.baidu.com', '-c3'}, procAttr)
if err != nil {
println("start process failed:" + err.String(),procAttr)
return
}

//waitMsg, err := process.Wait(os.WSTOPPED)
//if err != nil {
//println("Wait Error:" + err.String())
//}
//println("waitMsg:" + waitMsg.String())
//}
}
