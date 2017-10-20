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


_DedeKeyboardGrab = function() {
  var self = this;

  self.keypressed = 0;

  self.keydownHandler = function (e) {
    self.keypressed = e.keyCode;
    e.preventDefault();
    e.cancelBubble = true;
  };

  self.cancelEvent = function (e) {
    e.preventDefault();
    e.cancelBubble = true;
  };

  self.grab();
  self.initIcon();
};

_DedeKeyboardGrab.prototype.grab = function () {
  if (document.addEventListener) {
      document.addEventListener('keydown', this.keydownHandler, true);
      document.addEventListener('keyup', this.cancelEvent, true);
  }
};

_DedeKeyboardGrab.prototype.ungrab = function () {
  if (document.addEventListener) {
      document.removeEventListener('keydown', this.keydownHandler, true);
      document.removeEventListener('keyup', this.cancelEvent, true);
  }
};

_DedeKeyboardGrab.prototype.sleep = function (ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
};

_DedeKeyboardGrab.prototype.waitForKeyPress = async function (callback) {
  this.keypressed = null;
  for (var i = 0; i < 10; i++) {
    await this.sleep(1000);
    if (this.keypressed) {
      callback(this.keypressed);
      return;
    }
  }
}

_DedeKeyboardGrab.prototype.initIcon = function () {
  var sheet;
  for (var i = 0; i < window.document.styleSheets.length; i++) {
    sheet = window.document.styleSheets[i];
    if (sheet.cssRules) break;
  }

  var img = "url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAB3RJTUUH4AYeCAgm8dE6YwAAAMZJREFUSMft1S1OQ0EYheGnNySMJwXDBiZIRBWVrKEbqCuqG6hDkSCQLIEFYGADpJa7gFaVUP0ZEkybTG5q6I8p98g3mfPOfDPJ0Obo04E6pQEecfHH9XMMc8RrCeuUTvGUI4ZrwQLdLTf5mSOuivJLvKCXIzrVind3mMJZUd7HFL01q/Y16zqlEd5wXvJqj/c5wUkTVod+Ra3guAT3+DmYIEc84BZfmwTLHbq/C8k7rvHRFNxhsUX5DOPGSWa4wXP7l/yT/AJMwSmjcTvlbQAAAABJRU5ErkJggg==')";

  sheet.insertRule('#dede-keyboard-grab {position:absolute; z-index:9999999; width:24px; height:24px; background-image: ' + img + ';}', sheet.cssRules ? sheet.cssRules.length : 0);

  this.pointer = document.createElement('div');
  this.pointer.setAttribute('id', 'dede-keyboard-grab');
  this.pointer.setAttribute('style', 'left: 10; top: 10; visibility: hidden;');
  document.body.append(this.pointer);
};

_DedeKeyboardGrab.prototype.showIcon = function (state) {
  if (state) {
    this.pointer.style.visibility = "visible";
  } else {
    this.pointer.style.visibility = "hidden";
  }
};

DedeKeyboardGrab = new _DedeKeyboardGrab();
