// Copyright 2022 The Witness Contributors
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

package options

import "github.com/spf13/cobra"

type RunOptions struct {
	KeyOptions       KeyOptions
	ArchivistOptions ArchivistOptions
	WorkingDir       string
	Attestations     []string
	OutFilePath      string
	StepName         string
	Tracing          bool
	TimestampServers []string
}

func (ro *RunOptions) AddFlags(cmd *cobra.Command) {
	ro.KeyOptions.AddFlags(cmd)
	ro.ArchivistOptions.AddFlags(cmd)
	cmd.Flags().StringVarP(&ro.WorkingDir, "workingdir", "d", "", "Directory from which commands will run")
	cmd.Flags().StringSliceVarP(&ro.Attestations, "attestations", "a", []string{"environment", "git"}, "Attestations to record")
	cmd.Flags().StringVarP(&ro.OutFilePath, "outfile", "o", "", "File to which to write signed data.  Defaults to stdout")
	cmd.Flags().StringVarP(&ro.StepName, "step", "s", "", "Name of the step being run")
	cmd.Flags().BoolVar(&ro.Tracing, "trace", false, "Enable tracing for the command")
	cmd.Flags().StringSliceVar(&ro.TimestampServers, "timestamp-servers", []string{}, "Timestamp Authority Servers to use when signing envelope")
}

type ArchivistOptions struct {
	Enable bool
	Url    string
}

func (o *ArchivistOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&o.Enable, "enable-archivist", false, "Use Archivist to store or retrieve attestations")
	cmd.Flags().StringVar(&o.Url, "archivist-server", "https://archivist.testifysec.io", "URL of the Archivist server to store or retrieve attestations")
}
