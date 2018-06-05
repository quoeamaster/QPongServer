/*
 *  Copyright Project - CFactor, Author - quoeamaster, (C) 2018
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package util

import (
	"os"
	"fmt"
	"CFactor/TOML"
	"reflect"
)

const QpongConfPathEnv = "QPONG_CONF_PATH"
const QpongConfPathHardCoded = "/usr/local/qpong/server_conf.toml"
const QpongDefaultConfFilepath = "server_conf.toml"

/**
 *  struct to encapsulate the QPong server config
 */
type Config struct {
	AllowedAccessList []string `toml:"allowed-access-list"`
	ESHost string `toml:"es.host"`
	ServerPort int `toml:"server.port"`
	ServerDataPath string `toml:"server.data.path"`
}

/**
 *  getting the config file from 3 ways:
 *  1. get location from ENVIRONMENT variable => QPONG_CONF_PATH
 *  2. get from a hardcoded path /usr/share/qpong/server_conf.toml
 *  3. get from relative path (which is probably not a good idea)
 */
func GetConfigFile() (*os.File, error) {
	sPath := os.Getenv(QpongConfPathEnv)

	// 1. ENV
	if !IsStringEmpty(sPath) {
		_, err := isFilepathValid(sPath)
		if err != nil {
			return nil, err
		}   // end -- if (filepath valid??)
		// load the file
		return os.OpenFile(sPath, os.O_RDONLY, 0777)
	}

	// 2. hardcoded path
	_, err := isFilepathValid(QpongConfPathHardCoded)
	if err == nil {
		return os.OpenFile(QpongConfPathHardCoded, os.O_RDONLY, 0777)

	} else {
		// 3. relative filepath
		_, err = isFilepathValid(QpongDefaultConfFilepath)
		if err == nil {
			return os.OpenFile(QpongDefaultConfFilepath, os.O_RDONLY, 0777)
		} else {
			return nil, fmt.Errorf("could not open [%v], %v", QpongDefaultConfFilepath, err)
		}
	}
	return nil, fmt.Errorf("not config file available for loading")
}

/**
 *  helper to check if the provided path is valid (non-directory)
 */
func isFilepathValid(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, fmt.Errorf("%v", err)
	}
	if fileInfo.IsDir() {
		return false, fmt.Errorf("given path [%v] is not a valid File, instead it is a Directory", path)
	}
	return true, nil
}

func LoadConfigFromFilepath(path string) (cfgPtr *Config, err error) {
	cfgReader := TOML.TOMLConfigImpl{}
	cfg := Config{}

	cfgReader.Name = path
	cfgReader.StructType = reflect.TypeOf(cfg)

	_, err = cfgReader.Load(&cfg)
	if err != nil {
		err = fmt.Errorf("something wrong when loading the config [%v] => %v", path, err)
	}
	cfgPtr = &cfg
	return cfgPtr, err
}

/**
 *  a must provided lifecyclehook method
 */
func (o *Config) SetStructsReferences(refs *map[string]interface{}) error {
	return nil
}


