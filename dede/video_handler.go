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
	"sync"

	"github.com/gorilla/mux"
)

type videoHanlder struct {
	sync.RWMutex
	recorder *videoRecorder
}

func (v *videoHanlder) startRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	v.RLock()
	ok := v.recorder == nil
	v.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	vp, err := pathFromVars(vars, "video.mp4")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recorder := newVideoRecorder(vp, 1900, 1080, 10)
	if err := recorder.start(); err != nil {
		Log.Errorf("error while starting video record: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	v.Lock()
	v.recorder = recorder
	v.Unlock()

	Log.Infof("start video recording %s", idFromVars(vars, "video"))
	w.WriteHeader(http.StatusOK)
}

func (v *videoHanlder) stopRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	v.RLock()
	recorder := v.recorder
	v.RUnlock()
	if recorder == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	recorder.stop()

	Log.Infof("stop video recording %s", idFromVars(vars, "video"))
	w.WriteHeader(http.StatusOK)
}

func registerVideoHandler(router *mux.Router) *videoHanlder {
	t := &videoHanlder{}

	router.HandleFunc(baseURL+"/video/start-record", t.startRecord)
	router.HandleFunc(baseURL+"/video/stop-record", t.stopRecord)

	return t
}
