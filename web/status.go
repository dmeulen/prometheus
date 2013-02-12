// Copyright 2013 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"github.com/prometheus/prometheus/appstate"
	"html/template"
	"net/http"
)

type PrometheusStatus struct {
	Status  string
	Config  string
	Rules   string
	Targets string
}

type StatusHandler struct {
	appState *appstate.ApplicationState
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := &PrometheusStatus{
		Status:  "TODO: add status information here",
		Config:  h.appState.Config.ToString(0),
		Rules:   "TODO: list rules here",
		Targets: "TODO: list targets here",
	}
	t, _ := template.ParseFiles("web/templates/status.html")
	t.Execute(w, status)
}
