package main

import (
	`io/ioutil`
	`os`
	`path/filepath`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

const sshConfig = `Host *
  StrictHostKeyChecking no
  UserKnownHostsFile /dev/null
  LogLevel ERROR
`

func ssh(conf *config, logger simaqian.Logger) (err error) {
	logger.Info(`创建目录成功`, field.String(`folder`, conf.Folder))
	home := filepath.Join(os.Getenv(`HOME`), `.ssh`)
	if err = makeSSHHome(home, logger); nil != err {
		return
	}
	if err = writeSSHKey(home, conf.SSHKey, logger); nil != err {
		return
	}
	err = writeSSHConfig(home, logger)

	return
}

func makeSSHHome(home string, logger simaqian.Logger) (err error) {
	homeField := field.String(`home`, home)
	if err = os.MkdirAll(home, os.ModePerm); nil != err {
		logger.Error(`创建SSH目录出错`, homeField, field.Error(err))
	} else {
		logger.Info(`创建SSH目录成功`, homeField)
	}

	return
}

func writeSSHKey(home string, key string, logger simaqian.Logger) (err error) {
	keyfile := filepath.Join(home, `id_rsa`)
	keyfileField := field.String(`keyfile`, keyfile)
	if err = ioutil.WriteFile(keyfile, []byte(key), 0400); nil != err {
		logger.Error(`写入密钥文件出错`, keyfileField, field.Error(err))
	} else {
		logger.Info(`写入密钥文件成功`, keyfileField)
	}

	return
}

func writeSSHConfig(home string, logger simaqian.Logger) (err error) {
	configFile := filepath.Join(home, `config`)
	configFileField := field.String(`config.file`, configFile)
	if err = ioutil.WriteFile(configFile, []byte(sshConfig), 0400); nil != err {
		logger.Error(`写入SSH配置文件出错`, configFileField, field.Error(err))
	} else {
		logger.Info(`写入SSH配置文件成功`, configFileField)
	}

	return
}