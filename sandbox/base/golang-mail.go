package main

import (
  "bytes"
  "errors"
  "crypto/tls"
  "fmt"
  "log"
  "net"
  "net/smtp"
)

type loginAuth struct {
  username, password string
  host               string
}

func LoginAuth(username, password, host string) smtp.Auth {
  return &loginAuth{username, password, host}
}

func (a loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
  if server.Name != a.host {
    return "", nil, errors.New("wrong host name")
  }
  return "LOGIN", nil, nil
}

func (a loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
  if more {
    println(string(fromServer))
    if bytes.EqualFold([]byte("username:"), fromServer) {
      return []byte(a.username), nil
    } else if bytes.EqualFold([]byte("password:"), fromServer) {
      return []byte(a.password), nil
    }
  }
  return nil, nil
}

func main() {

  var (
    host     = "bjmail.kingsoft.com"
    port     = 587
    from     = "zhanglei16@kingsoft.com"
    password = "Kingsoftzhang0"
    to       = []string{"osfun@qq.com"}
    msg      = []byte("This is my message")
    auth     = LoginAuth(from, password, host)
  )

  conf := &tls.Config{
    InsecureSkipVerify: true,
    ServerName:         host,
  }

  serverAddr := fmt.Sprintf("%s:%d", host, port)

  //  Dial returns a new Client connected to an SMTP server at addr. The addr must include a port number.
  conn, err := net.Dial("tcp", serverAddr)
  if err != nil {
    log.Printf("Error Dialing %s\n", err)
    return
  }

  //  NewClient returns a new Client using an existing connection and host as a server name to be used when authenticating.
  client, err := smtp.NewClient(conn, host)
  if err != nil {
    log.Printf("Error SMTP connection: %s\n", err)
    return
  }

  // StartTLS sends the STARTTLS command and encrypts all further communication. Only servers that advertise the STARTTLS extension support this function.
  if err = client.StartTLS(conf); err != nil {
    log.Printf("Error performing StartTLS: %s\n", err)
    return
  }

  // Extension reports whether an extension is support by the server. The extension name is case-insensitive. If the extension is supported, Extension also returns a string that contains any parameters the server specifies for the extension.
  if ok, _ := client.Extension("AUTH"); ok {
    if err := client.Auth(auth); err != nil {
      log.Printf("Error during AUTH %s\n", err)
      return
    }
  }

  if err := client.Mail(from); err != nil {
    log.Printf("Error: %s\n", err)
    return
  }

  for _, addr := range to {
    if err := client.Rcpt(addr); err != nil {
      log.Printf("Error: %s\n", err)
      return
    }
  }

  w, err := client.Data()
  if err != nil {
    log.Printf("Error: %s\n", err)
    return
  }

  fmt.Println(msg)
  //_, err1 = w.Write(msg)
  _, err = w.Write([]byte("aaaaaaaaaaaaaaaaaaaaa"))
  if err != nil {
    fmt.Println("=======================")
    log.Printf("Error: %s\n", err)
    return

  }

 // close the connection
  err = w.Close()
  if err != nil {
    log.Printf("Error: %s\n", err)
    return

  }

  // mail client exit
  client.Quit()
}
