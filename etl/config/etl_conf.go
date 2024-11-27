// Created at: 2024/11/27
// Project : transfer-data

package config

import (
    "encoding/json"
    "errors"
    "gopkg.in/yaml.v3"
    "os"
    "strings"
)

// SysConf 系统基础配置
var SysConf = struct {
    KafkaReaderList []ReaderConfig `json:"kafka_reader_list" yaml:"kafkaReaderList"`
}{}

func LoadSysConf(filePath string) error {
    fileType := filePath[strings.LastIndex(filePath, ".")+1:]
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }

    defer func(file *os.File) {
        err := file.Close()
        if err != nil {
            panic("load sys conf err, err: " + err.Error())
        }
    }(file)
    switch fileType {
    case "yaml":
        decoder := yaml.NewDecoder(file)
        err = decoder.Decode(&SysConf)
        break
    case "json":
        decoder := json.NewDecoder(file)
        err = decoder.Decode(&SysConf)
        break
    default:
        err = errors.New("sys conf file format err, only support json or yaml")
    }
    if err != nil {
        return err
    }
    return nil
}
