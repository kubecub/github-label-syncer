// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
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

package commands

import "github.com/spf13/cobra"

type BuildInfo struct {
	GoVersion string `json:"goVersion"`
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	Date      string `json:"date"`
}

type Executor struct {
	rootCmd    *cobra.Command
	runCmd     *cobra.Command
	lintersCmd *cobra.Command

	exitCode  int
	buildInfo BuildInfo
}
