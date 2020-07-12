/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

var chartFolder = ""

// Get the directory of charts
func GetChartsDirectory() string {
	if chartFolder != "" {
		return chartFolder
	}
	homeChartsFolder := os.Getenv("HOME") + "/charts"
	if !PathExists(homeChartsFolder) {
		chartFolder = "/charts"
	} else {
		chartFolder = homeChartsFolder
	}
	return chartFolder
}