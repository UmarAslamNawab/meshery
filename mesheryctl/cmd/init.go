// Copyright 2019 The Meshery Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Meshery",
	Long:  `Check and installs docker and docker-compose if not exists`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := exec.Command("chmod", "+x", dockerComposeBinary).Run(); err != nil {
			log.Info("Prerequisite Docker Compose not available. Attempting Docker Compose installation...")
			ostype, osarch := prereq()
			osdetails := strings.TrimRight(string(ostype), "\r\n") + "-" + strings.TrimRight(string(osarch), "\r\n")
			dockerComposeBinaryURL := dockerComposeBinaryURL + "-" + osdetails
			if err := downloadFile(dockerComposeBinary, dockerComposeBinaryURL); err != nil {
				log.Fatal(err)
			}
			if err := exec.Command("chmod", "+x", dockerComposeBinary).Run(); err != nil {
				log.Fatal(err)
			}
		}
		log.Info("Prerequisite Docker Compose is installed.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
