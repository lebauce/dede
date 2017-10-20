/*
 * Copyright (C) 2017 Red Hat, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package dede

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/skydive-project/dede/statics"
)

type keyboardGrabHandler struct{}

func (t *keyboardGrabHandler) keyboardGrabInstall(w http.ResponseWriter, r *http.Request) {
	asset := statics.MustAsset("statics/js/keyboard-grab.js")

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	w.Write(asset)
}

func registerKeyboardGrabHandler(router *mux.Router) *keyboardGrabHandler {
	f := &keyboardGrabHandler{}

	router.HandleFunc("/keyboard-grab/install", f.keyboardGrabInstall)

	return f
}
