package main

import (
    "fmt"
    "os"
    "gopkg.in/ini.v1"
    "io/ioutil"
    //"regexp"
    //"reflect"
    //"strings"
    //"encoding/json"
    "text/template"
    "os/exec"
    "./utils"
)

type Config struct {
    Default struct {
        SiteName string `ini:"site_name"`
        SitePath string `ini:"site_path"`
        IsAddHosts bool `ini:"add_hosts"`
        IsCreateTestFolder bool `ini:"create_test_folder"`
        IsCreateNginx bool `ini:"is_create_nginx"`
        IsCreateHttpd bool `ini:"is_create_httpd"`

    } `ini:"DEFAULT"`
    Httpd struct {
        Port string `ini:"port"`
        Email string `ini:"email"`
        LogPath string `ini:"log_path"`
        Host string `ini:"host"`
    } `ini:"HTTPD"`
    Nginx struct {
        Port string `ini:"port"`
        ProxyPass string `ini:"proxy_pass"`
    } `ini:"NGINX"`
    System struct {
        HttpdConfigPath string `ini:"httpd_config_path"`
        HttpdPath string `ini:"httpd_path"`
        NginxConfigPath string `ini:"nginx_config_path"`
        NginxPath string `ini:"nginx_path"`
        FileHostsPath string `ini:"file_hosts_path"`
        HostsIP string `ini:"hosts_ip"`
        TemplatePagePath string `ini:"template_page_path"`
        HttpdTemplateName string `ini:"httpd_template_name"`
        NginxTemplateName string `ini:"nginx_template_name"`
        OwnerUserName string `ini:"owner_user_name"`
    } `ini:"SYSTEM"`
}

var ch chan string = make(chan string)

func main() {
    arguments := _getArguments()

    if len(arguments) < 1 {
        return;
    }

    options := get_options()

    if len(arguments[0]) > 1 {
        options.Default.SiteName = arguments[0]
    }
    doCreateHttpd(options)
    doCreateNginx(options)
    doCreateHosts(options)
    doCreateSite(options)

    //fmt.Println(<-ch)


//     var input string
//     fmt.Scanln(&input)
//     fmt.Println("Done")
    //fmt.Println(options.Httpd.Port);

}

func doCreateHosts(options *Config) {
    if !options.Default.IsAddHosts {
        return
    }

    f, err := os.OpenFile(options.System.FileHostsPath, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }

    defer f.Close()
    var text string = "\n#Created By NIL VirtualHost\n127.0.0.1 " + options.Default.SiteName
    if _, err = f.WriteString(text); err != nil {
        panic(err)
    }

     utils.Info("Add Record In "+options.System.FileHostsPath);

}

func doCreateSite(options *Config) {
    if !options.Default.IsCreateTestFolder {
        return
    }

    var folderPath string = options.Default.SitePath + options.Default.SiteName

    err := os.MkdirAll(folderPath, os.ModePerm)

    if err != nil {
        utils.Error("Folder No Created.");
        return
    }

    data, err := ioutil.ReadFile("data/index.html.data")
    if err != nil {
        utils.Error("Template File Not Found")
        return
    }

    file, err := os.Create(folderPath + "/index.html")

    if err != nil {
        utils.Error("Index File Not Created");
        return
    }

    //_ = file.Chown(501, 20)


    _, err = file.WriteString(string(data))
    if err != nil {
        utils.Error("Content Not Saved");
        return
    }

    utils.Info("Template Created. Path: " + folderPath)

    if len(options.System.OwnerUserName) <= 0 {
        return
    }

    err = exec.Command("chown", "-R", options.System.OwnerUserName, folderPath).Run()
    if err != nil {
        utils.Error("Command Whoami Failed")
        return
    }
    utils.Debug("Chown Folder Changed. File: " + folderPath)
}

func doCreateHttpd(options *Config) {
    if !options.Default.IsCreateHttpd {
        return
    }

    var fileName string;
    fileName = options.System.HttpdPath + options.Default.SiteName + ".conf";

    doCreateConfigFile(fileName, options, options.System.HttpdTemplateName)

//     err := exec.Command("apachectl", "-k", "restart").Run()
//     if err != nil {
//         utils.Error("Failed Apache Restarted")
//         return
//     }
//     utils.Debug("Apache was Restart")
}

func doCreateNginx(options *Config) {
     if !options.Default.IsCreateNginx {
        return
     }

     var fileName string;
     fileName = options.System.NginxPath + options.Default.SiteName + ".conf";

     doCreateConfigFile(fileName, options, options.System.NginxTemplateName)
}

func doCreateConfigFile(fileName string, options *Config, templateName string) {
    file, err := os.Create(fileName)

    if err != nil {
        utils.Error("Need Sudo Rules");
        return
    }
    defer file.Close()

    errorMsg := doParseTemplate(templateName, options, file)
    if len(errorMsg) > 1 {
        utils.Error(errorMsg)
        return
    }

    utils.Info("File: " +fileName + " Was Created!");
}

func doParseTemplate(fileTemplate string, options *Config, sourceFile *os.File) string {
    data, err := ioutil.ReadFile("data/"+fileTemplate)
    if err != nil {
        return "File reading error"
    }
    template, err := template.New("confData").Parse(string(data))

    err = template.Execute(sourceFile, options)

    if err != nil {
       return "Error in Replacing Template"
    }

    return ""
}

func get_options() *Config {
    cfg, err := ini.Load("config.ini")

    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
    }

    config := &Config{}

    err = cfg.Section("").MapTo(config)

    return config;
}

func _getArguments() []string {
    return os.Args[1:]
}