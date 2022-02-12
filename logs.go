// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import (
	"encoding/json"
)

type LogConfig struct {
	Filename string `json:filename`
	Level    int    `json:level`
	Maxlines int    `json:maxlines`
	Maxsize  int    `json:maxsize`
	Daily    bool   `json:daily`
	Maxdays  int    `json:maxdays`
}

// BeeLogger references the used application logger.
var AppLogger *BeeLogger

func InitLog(filename string, level int,adapterFile string,maxdays int) {
	var config LogConfig = LogConfig{
		Filename: filename,
		Level:    level,
		Maxlines: 20000,
		Maxsize:  0,
		Daily:    true,
		Maxdays:  maxdays,
	}

	con, _ := json.Marshal(config)
	AppLogger = GetBeeLogger()
	AppLogger.SetLogger(adapterFile, string(con))
}

